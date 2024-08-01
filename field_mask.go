package field_mask

import (
	"cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"reflect"
	"slices"
	"strings"
	"sync"
)

func Apply(mask *fieldmaskpb.FieldMask, pb proto.Message, out any) error {
	if pb == nil {
		return nil
	}

	var valuesMap map[string]any

	if mask == nil {
		fields := pb.ProtoReflect().Descriptor().Fields()
		valuesMap = make(map[string]any, fields.Len())

		for i := range fields.Len() {
			switch {
			case fields.Get(i).IsList():
				valuesMap[fields.Get(i).TextName()] = protoListToSlice(pb.ProtoReflect().NewField(fields.Get(i)).List())
			case fields.Get(i).IsMap():
				valuesMap[fields.Get(i).TextName()] = protoMapToGoMap(fields.Get(i), pb.ProtoReflect().NewField(fields.Get(i)).Map())
			default:
				valuesMap[fields.Get(i).TextName()] = nil
			}
		}

		mapValues(pb, valuesMap)
	} else {
		valuesMap = make(map[string]any, len(mask.Paths))
		if !mask.IsValid(pb) {
			return ErrInvalidFieldMask
		}

		mapFilteredValues(pb, mask, valuesMap)
	}

	return setFields(out, valuesMap)
}

// mapFilteredValues reads the supplied mask and Nested and sets the respective fields
// All inputs are expected to be non-nil
func mapFilteredValues(pb proto.Message, mask *fieldmaskpb.FieldMask, fields map[string]any) {
	for _, path := range mask.Paths {
		fields[path] = nil
	}

	pb.ProtoReflect().Range(findFilteredFields(&strings.Builder{}, fields))
}

// findFilteredFields returns a callback-function that recursively searches for all keys in the fields parameter and sets their corresponding value
// The function uses fieldPath to build the field path as it recurses deeper
// The callback is to be used with google.golang.org/protobuf/reflect/protoreflect.Nested.Range()
func findFilteredFields(fieldPath *strings.Builder, fields map[string]any) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		var fp string
		if fieldPath.Len() > 0 {
			fp = fieldPath.String() + "."
		}
		fp += fd.TextName()

		if _, ok := fields[fp]; ok {
			switch {
			case fd.IsList():
				s := protoListToSlice(v.List())
				fields[fp] = s

			case fd.IsMap():
				m := protoMapToGoMap(fd, v.Map())
				fields[fp] = m

			default:
				fields[fp] = v.Interface()
			}
			return true
		}

		if fd.Kind() == protoreflect.MessageKind {
			inner := &strings.Builder{}
			inner.WriteString(fp)

			v.Message().Range(findFilteredFields(inner, fields))
		}

		return true
	}
}

// mapValues creates a map with paths to the non-empty values in the protobuf and their respective values
func mapValues(pb proto.Message, fields map[string]any) {
	pb.ProtoReflect().Range(findFields(&strings.Builder{}, fields))
}

// findFields returns a function to pass to protoreflect.Message.Range() that searches for non-empty fields
func findFields(fieldPath *strings.Builder, fields map[string]any) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		var fp string
		if fieldPath.Len() > 0 {
			fp = fieldPath.String() + "."
		}
		fp += fd.TextName()

		switch {
		case fd.IsList():
			sl := protoListToSlice(v.List())
			fields[fp] = sl

		case fd.IsMap():
			m := protoMapToGoMap(fd, v.Map())
			fields[fp] = m

		case fd.Kind() == protoreflect.MessageKind:
			inner := &strings.Builder{}
			inner.WriteString(fp)

			v.Message().Range(findFields(inner, fields))

		default:
			fields[fp] = v.Interface()
		}
		return true
	}
}

// protoListToSlice takes a protobuf slice and returns a golang slice
// The first arg is only used for type inference
func protoListToSlice(pb protoreflect.List) any {
	s := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(pb.NewElement().Interface())), 0, pb.Len())
	for i := range pb.Len() {
		s = reflect.Append(s, reflect.ValueOf(pb.Get(i).Interface()))
	}
	return s.Interface()
}

// protoMapToGoMap takes a protobuf map and generates a new go map
func protoMapToGoMap(t protoreflect.FieldDescriptor, pb protoreflect.Map) any {
	m := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(t.MapKey().Default().Interface()), reflect.TypeOf(t.MapValue().Default().Interface())))
	pb.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		m.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(v.Interface()))
		return true
	})
	return m.Interface()
}

func setFields(out any, filteredValues map[string]any) error {
	outValue := reflect.ValueOf(out)
	switch ov := outValue.Kind(); ov {
	case reflect.Pointer:
		switch ovElem := outValue.Elem().Kind(); ovElem {
		default:
			return ErrReceiverNotPointerToStruct
		case reflect.Struct:
			break
		}
	default:
		return ErrReceiverNotPointerToStruct
	}

	fields := loadCachedStructFields(reflect.TypeOf(out).Elem())

	for fieldPath, v := range filteredValues {
		// first match whole fieldPath to allow hoisting
		ft := fields.byExactName[fieldPath]

		if ft == nil {

			// if no exact match were found; search for the top-most field in the hierarchy
			// if found; load the relevant fields and see if we get an exact match with the remaining fieldPath
			// if not; cut the remaining fieldPath again and do it again

			var (
				current, remaining, ok = strings.Cut(fieldPath, ".")
				index                  []int
			)

			for ok {
				ft = fields.byExactName[current]
				if ft != nil {
					index = append(index, ft.index...)
					fields = loadCachedStructFields(ft.typ)
				} else {
					// No match
					break
				}

				// Check for most specific field first to allow hoisting
				ft = fields.byExactName[remaining]
				if ft != nil {
					// found it
					ft = &field{
						name:   ft.name,
						typ:    ft.typ,
						tagged: ft.tagged,
						index:  append(index, ft.index...),
					}
					break
				}
				// no match found, cut remaining fieldPath and iterate again, breaking if there are no more levels
				current, remaining, ok = strings.Cut(remaining, ".")
			}
		}

		if ft != nil {
			// iterate down the hierarchy and allocate any nil pointers encountered
			subV := outValue
			for _, i := range ft.index {
				if subV.Kind() == reflect.Pointer {
					if subV.IsNil() {
						// If a struct embeds a pointer to an unexported type,
						// it is not possible to set a newly allocated value
						// since the field is unexported.
						//
						// See https://golang.org/issue/21357
						if !subV.CanSet() {
							return ErrUnsettableReceiver
						}

						subV.Set(reflect.New(subV.Type().Elem()))
					}
					subV = subV.Elem()
				}
				subV = subV.Field(i)
			}

			fv := outValue.Elem().FieldByIndex(ft.index)

			var vv reflect.Value
			if v != nil {
				vv = reflect.ValueOf(v)
			} else {
				vv = reflect.Zero(ft.typ)
			}

			if vv.Kind() == reflect.Pointer {
				vv = vv.Elem()
			}

			// set the value

			if fv.Kind() == reflect.Pointer {
				if fv.Type().Elem().Kind() != vv.Kind() {
					return ErrInvalidReceiverField
				}

				if fv.IsNil() {
					fv.Set(reflect.New(fv.Type().Elem()))
				}

				fv.Elem().Set(vv)
			} else {
				if fv.Kind() != vv.Kind() {
					return ErrInvalidReceiverField
				}
				fv.Set(vv)
			}

		}
	}

	return nil
}

type field struct {
	name   string
	typ    reflect.Type
	tagged bool
	index  []int
}

type structFields struct {
	list        []field
	byExactName map[string]*field
}

var internalCache sync.Map // map[reflect.Type]structFields

func loadCachedStructFields(t reflect.Type) structFields {
	if v, ok := internalCache.Load(t); ok {
		return v.(structFields)
	}

	f, _ := internalCache.LoadOrStore(t, typeFields(t))
	return f.(structFields)
}

func typeFields(t reflect.Type) structFields {
	var (
		fields    structFields
		fieldList []field

		current = []field{}
		next    = []field{{typ: t}}

		count, nextCount map[reflect.Type]int

		visited = map[reflect.Type]bool{}
	)

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.typ] {
				continue
			}
			visited[f.typ] = true

			for i := 0; i < f.typ.NumField(); i++ {
				sf := f.typ.Field(i)
				if sf.Anonymous {
					t := sf.Type
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
					if !sf.IsExported() && t.Kind() != reflect.Struct {
						// Ignore embedded fields of unexported non-struct types
						continue
					}

					// Do not ignore embedded fields of unexported struct types
					// since they might have exported fields.
				} else if !sf.IsExported() {
					// Ignore unexported non-embedded fields.
					continue
				}

				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Kind() == reflect.Pointer {
					ft = sf.Type.Elem()
				}

				var (
					fieldName string
					tagged    bool
				)

				if !sf.Anonymous && f.name != "" {
					fieldName = f.name + "."
				}
				if tt, ok := sf.Tag.Lookup(tagName); ok && tt != "" {
					if tt == "-" {
						continue
					}
					tagged = true
					fieldName += tt
				} else {
					if !sf.Anonymous {
						fieldName += sf.Name
					}
				}

				// can fieldName become `""` here?
				if !sf.Anonymous || ft.Kind() != reflect.Struct {
					res := field{
						tagged: tagged,
						name:   fieldName,
						typ:    ft,
						index:  index,
					}

					fieldList = append(fieldList, res)

					if ft.Kind() == reflect.Struct {
						next = append(next, res)
					}

					if count[f.typ] > 1 {
						fieldList = append(fieldList, fieldList[len(fieldList)-1])
					}
					continue
				}

				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, field{name: fieldName, index: index, typ: ft})
				}
			}
		}
	}

	slices.SortFunc(fieldList, func(a, b field) int {

		// sort field by name, breaking ties with depth, then
		// breaking ties with "name came from struct tag", then
		// breaking ties with index sequence.

		if c := strings.Compare(a.name, b.name); c != 0 {
			return c
		}
		if c := cmp.Compare(len(a.index), len(b.index)); c != 0 {
			return c
		}
		if a.tagged != b.tagged {
			if a.tagged {
				return -1
			}
			return 1
		}
		return slices.Compare(a.index, b.index)
	})

	out := fieldList[:0]
	for advance, i := 0, 0; i < len(fieldList); i += advance {
		// One iteration per name
		fi := fieldList[i]
		name := fi.name
		for advance = 1; i+advance < len(fieldList); advance++ {
			fj := fieldList[i+advance]
			if fj.name != name {
				break
			}
		}
		if advance == 1 { // Only one field with this name was found
			out = append(out, fi)
			continue
		}

		if dominant, ok := dominantField(fieldList[i : i+advance]); ok {
			out = append(out, dominant)
		}
	}

	fieldList = out
	slices.SortFunc(fieldList, func(a, b field) int {
		return slices.Compare(a.index, b.index)
	})

	fields.list = fieldList
	fields.byExactName = make(map[string]*field, len(fieldList))

	for j, ff := range fields.list {
		fields.byExactName[ff.name] = &fields.list[j]
	}

	return fields
}

// dominantField looks through the fields, all of which are known to
// have the same name, to find the single field that dominates the
// others using Go's embedding rules, modified by the presence of
// our struct tags. If there are multiple top-level fields, the boolean
// will be false: This condition is an error in Go, and we skip all
// the fields.
func dominantField(fields []field) (field, bool) {
	// The fields are sorted in increasing index-length order, then by presence of tag.
	// That means that the first field is the dominant one. We need only check
	// for error cases: two fields at top level, either both tagged or neither tagged.
	if len(fields) > 1 && len(fields[0].index) == len(fields[1].index) && fields[0].tagged == fields[1].tagged {
		return field{}, false
	}
	return fields[0], true
}
