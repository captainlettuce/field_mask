package field_mask

import (
	"errors"
	"github.com/captainlettuce/field_mask/testing/pb"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"reflect"
	"testing"
)

type base struct {
	unexportedField string // ignored
	IgnoredField    string `field_mask:"-"`

	A             string `field_mask:"a"`
	B             int32  `field_mask:"b"`
	C             string `field_mask:"c"`
	UntaggedField string
}

type override struct {
	NewUntagged   string `field_mask:"UntaggedField"`
	UntaggedField string
}

type sliceMap struct {
	Slice []int32          `field_mask:"array"`
	Map   map[int32]string `field_mask:"map"`
}

type pointerReceivers struct {
	A             *string `field_mask:"a"`
	B             *int32  `field_mask:"b"`
	C             *string `field_mask:"c"`
	UntaggedField *string
}

type nested struct {
	A       string `field_mask:"a"`
	NestedA string `field_mask:"base.a"`
	Base    base   `field_mask:"base"`
}

type nestedRecursive struct {
	A       string           `field_mask:"a"`
	B       int32            `field_mask:"b"`
	NestedA string           `field_mask:"nested.a"`
	Nested  *nestedRecursive `field_mask:"nested"`
}

type nestedRecursiveVariantA struct {
	A *nestedRecursiveVariantA `field_mask:"a"`
	B *nestedRecursiveVariantB `field_mask:"b"`
	C string                   `field_mask:"c"`
}

type nestedRecursiveVariantB struct {
	A *nestedRecursiveVariantA `field_mask:"a"`
	B *nestedRecursiveVariantB `field_mask:"b"`
	C string                   `field_mask:"c"`
}

type embeddedNoTag struct {
	base        // embedded fields will be mapped as if they were top level fields
	A    string `field_mask:"a"` // A will take precedence over embedded `base.A` because of lesser field depth
}

type embeddedTagged struct {
	base `field_mask:"base"` // Embedded fields will be mapped as `base.X`

	A string `field_mask:"a"` // name no longer collides with `base.name`
}

// embeds 2 equal structs to make sure the code path for multiple structs is run
// child2 & child3 will cancel each-other out since they are identical at the same depth
type childEmbedder struct {
	child2
	child3
}

type child2 struct {
	base
}

type child3 struct {
	base
}

type invalidReceiver struct {
	A int32    `field_mask:"a"`
	B *float64 `field_mask:"b"`
}

type tieChildA struct {
	A string `field_mask:"a"`
}

type tieChildB struct {
	B int32  `field_mask:"b"`
	A string `field_mask:"a"`
}

type tieBreakIndexSequence struct {
	ChildB tieChildB `field_mask:"base"`
	ChildA tieChildA `field_mask:"base"`
}

func TestApply(t *testing.T) {
	tests := []struct {
		name    string
		pb      proto.Message
		mask    *fieldmaskpb.FieldMask
		want    any
		out     any
		wantErr error
		cmpOpts []cmp.Option
	}{
		// Happy path

		{
			name: "base",
			out:  &base{},
			pb: &pb.Base{
				A:             "test",
				B:             1,
				UntaggedField: "tset",
				C:             ref("optional"),
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "b", "UntaggedField"}},
			want: &base{
				A:             "test",
				B:             1,
				UntaggedField: "tset",
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(base{})},
		},
		{
			// empty field mask should be treated as if all fields were set
			// effectively wiping unset values
			// https://protobuf.dev/reference/protobuf/google.protobuf/#field-masks-updates
			name: "empty field mask",
			out:  &pointerReceivers{},
			pb: &pb.Base{
				A:             "test",
				B:             1,
				C:             nil,
				UntaggedField: "",
			},
			mask: nil,
			want: &pointerReceivers{
				A:             ref("test"),
				B:             ref[int32](1),
				C:             ref(""),
				UntaggedField: ref(""),
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(base{})},
		},
		{
			name: "tag override",
			out:  &override{},
			pb: &pb.Base{
				UntaggedField: "test",
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"UntaggedField"}},
			want: &override{NewUntagged: "test"},
		},
		{
			name: "slice map",
			out:  &sliceMap{},
			pb: &pb.ArrayMap{
				Array: []int32{1, 2, 3},
				Map:   map[int32]string{1: "test"},
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"array", "map"}},
			want: &sliceMap{Slice: []int32{1, 2, 3}, Map: map[int32]string{1: "test"}},
		},
		{
			name: "slice map nil field mask",
			out:  &sliceMap{},
			pb: &pb.ArrayMap{
				Array: []int32{1, 2, 3},
				Map:   map[int32]string{1: "test"},
			},
			want: &sliceMap{
				Slice: []int32{1, 2, 3},
				Map:   map[int32]string{1: "test"},
			},
		},
		{
			name: "slice map nil field mask unset",
			out:  &sliceMap{},
			pb:   &pb.ArrayMap{},
			want: &sliceMap{
				Slice: []int32{},
				Map:   map[int32]string{},
			},
		},
		{
			name: "pointer receivers",
			out:  &pointerReceivers{},
			pb: &pb.Base{
				A:             "test",
				B:             3,
				UntaggedField: "",
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "b", "UntaggedField"}},
			want: &pointerReceivers{
				A:             ref("test"),
				B:             ref(int32(3)),
				UntaggedField: ref(""),
			},
		},
		{
			name: "optional message fields",
			out:  &pointerReceivers{},
			pb: &pb.OptionalFields{
				A: ref(""),
				B: ref[int32](1),
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "b"}},
			want: &pointerReceivers{
				A: ref(""),
				B: ref[int32](1),
			},
		},
		{
			name: "nested message empty field mask",
			out:  &nested{},
			pb: &pb.Nested{
				A: "level 0",
				Base: &pb.Base{
					A:             "level 1",
					B:             1,
					UntaggedField: "",
				},
			},
			mask: nil,
			want: &nested{
				A:       "level 0",
				NestedA: "level 1",
				Base:    base{B: 1},
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(base{})},
		},
		{
			name: "nested message",
			out:  &nested{},
			pb: &pb.Nested{
				A: "test",
				Base: &pb.Base{
					A:             "inner",
					B:             1,
					UntaggedField: "untagged",
				},
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "base.a", "base.b"}},
			want: &nested{
				A:       "test",
				NestedA: "inner",
				Base:    base{B: 1},
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(base{})},
		},
		{
			name: "recursive message",
			out:  &nestedRecursive{},
			pb: &pb.NestedRecursive{
				A: "level 0",
				Nested: &pb.NestedRecursive{
					A: "level 1",
					Nested: &pb.NestedRecursive{
						A: "level 2",
						Nested: &pb.NestedRecursive{
							A: "level 3",
							B: 3,
							Nested: &pb.NestedRecursive{
								A: "level 4",
							},
						},
					},
				},
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "nested.nested.nested.a", "nested.nested.nested.b"}},
			want: &nestedRecursive{
				A: "level 0",
				Nested: &nestedRecursive{
					Nested: &nestedRecursive{
						NestedA: "level 3",
						Nested: &nestedRecursive{
							B: 3,
						},
					},
				},
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(nestedRecursive{})},
		},
		{
			name: "recursive recursive message",
			out:  &nestedRecursiveVariantA{},
			pb: &pb.NestedRecursiveVariantA{
				A: &pb.NestedRecursiveVariantA{
					B: &pb.NestedRecursiveVariantB{
						B: &pb.NestedRecursiveVariantB{
							B: &pb.NestedRecursiveVariantB{
								C: "test",
							},
							C: "not this",
						},
					},
				},
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a.b.b.b.c"}},
			want: &nestedRecursiveVariantA{
				A: &nestedRecursiveVariantA{
					B: &nestedRecursiveVariantB{
						B: &nestedRecursiveVariantB{
							B: &nestedRecursiveVariantB{
								C: "test",
							},
						},
					},
				},
			},
		},
		{
			name: "embedded struct no tags",
			pb: &pb.Base{
				A: "a",
				B: 1,
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "b"}},
			out:  &embeddedNoTag{},
			want: &embeddedNoTag{
				base: base{B: 1},
				A:    "a", // the higher-level field gets priority, according to go embed rules
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(embeddedNoTag{}, base{})},
		},
		{
			name: "embedded struct with tags",
			pb: &pb.Nested{
				A: "a",
				Base: &pb.Base{
					A: "inner", // note that `base.a` is not selected in the mask, so it gets ignored
					B: 1,
				},
			},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"a", "base.b"}},
			out:  &embeddedTagged{},
			want: &embeddedTagged{
				base: base{B: 1},
				A:    "a",
			},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(embeddedTagged{}, base{})},
		},
		{
			name: "tie break index sequence",
			out:  &tieBreakIndexSequence{},
			pb:   &pb.Nested{Base: &pb.Base{A: "test", B: 1}},
			mask: &fieldmaskpb.FieldMask{Paths: []string{"base.a", "base.b"}},
			want: &tieBreakIndexSequence{
				ChildB: tieChildB{B: 1},
			},
		},
		// Sad paths
		{
			name: "invalid receiver field type",
			out:  &invalidReceiver{},
			pb: &pb.Base{
				A: "test",
			},
			mask:    &fieldmaskpb.FieldMask{Paths: []string{"a"}},
			wantErr: ErrInvalidReceiverField,
		},
		{
			name: "invalid receiver field field type (pointer)",
			out:  &invalidReceiver{},
			pb: &pb.Base{
				B: 1,
			},
			mask:    &fieldmaskpb.FieldMask{Paths: []string{"b"}},
			wantErr: ErrInvalidReceiverField,
		},
		{
			name:    "invalid mask",
			pb:      &pb.Base{},
			mask:    &fieldmaskpb.FieldMask{Paths: []string{"does-not-exist"}},
			wantErr: ErrInvalidFieldMask,
		},
		{
			name:    "nil interface receiver",
			out:     nil,
			mask:    &fieldmaskpb.FieldMask{},
			pb:      &pb.Base{},
			wantErr: ErrReceiverNotPointerToStruct,
		},
		{
			name:    "nil proto message",
			pb:      nil,
			out:     &base{},
			mask:    &fieldmaskpb.FieldMask{},
			want:    &base{},
			cmpOpts: []cmp.Option{cmp.AllowUnexported(base{})},
		},
		{
			name:    "pointer to non-struct",
			out:     ref(""),
			pb:      &pb.Base{},
			mask:    &fieldmaskpb.FieldMask{},
			wantErr: ErrReceiverNotPointerToStruct,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Apply(tt.mask, tt.pb, tt.out)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == nil && tt.out != nil && !cmp.Equal(tt.out, tt.want, tt.cmpOpts...) {
				t.Errorf("Apply() got = %v, want %v\n%s", tt.out, tt.want, cmp.Diff(tt.out, tt.want, tt.cmpOpts...))
			}
		})
	}
}

func Test_typeFields(t *testing.T) {
	type res struct {
		typ   reflect.Type
		depth int
	}
	tests := []struct {
		name    string
		arg     any
		want    structFields
		wantMap map[string]res
	}{
		{
			name: "embedded non-struct pointer",
			arg: struct {
				*string
				A string
			}{},
			wantMap: map[string]res{
				"A": {typ: reflect.TypeOf(""), depth: 1},
			},
		},
		{
			name:    "Multiple fields same struct base",
			arg:     childEmbedder{},
			wantMap: map[string]res{}, // all fields get skipped since they are identical at the same level
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structType := reflect.TypeOf(tt.arg)
			tf := typeFields(structType)
			require.Len(t, tf.list, len(tt.wantMap))
			for k, want := range tt.wantMap {
				have, ok := tf.byExactName[k]
				require.True(t, ok, "key not found in result: %s", k)
				require.Equal(t, want.typ, have.typ, k)
				require.Len(t, have.index, want.depth, k)
			}
		})
	}
}

func TestNoRefMap(t *testing.T) {
	msg := &pb.ArrayMap{
		Array: []int32{1, 2, 3},
		Map:   map[int32]string{1: "test", 2: "test2"},
	}
	mask := &fieldmaskpb.FieldMask{Paths: []string{"array", "map"}}
	out := &sliceMap{}

	if err := Apply(mask, msg, out); err != nil {
		t.Errorf("Apply() error = %v", err)
	}

	msg.Array[0] = 3
	if msg.Array[0] == out.Slice[0] {
		t.Errorf("Slice was not deep copied")
	}

	msg.Map[1] = "tset"
	if msg.Map[1] == out.Map[1] {
		t.Errorf("Map was not deep copied")
	}

}

func ref[T any](v T) *T {
	return &v
}
