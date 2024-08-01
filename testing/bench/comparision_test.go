package bench

import (
	"fmt"
	"github.com/captainlettuce/field_mask"
	"github.com/captainlettuce/field_mask/testing/pb"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"reflect"
	"testing"
)

func BenchmarkPointerStructFields(b *testing.B) {
	var (
		simpleMessage pb.BenchmarkTest
		m             fieldmask_utils.Mask
		err           error
	)

	cases := []struct {
		name string
		out  reflect.Type
	}{
		{
			name: "pointer receiver",
			out:  reflect.TypeOf(outStructPointer{}),
		},
		{
			name: "raw receiver",
			out:  reflect.TypeOf(outStruct{}),
		},
	}

	for _, c := range cases {

		b.StopTimer()
		simpleMessage = generateSimpleMessage()
		fm := generateMask()
		b.StartTimer()

		b.Run(fmt.Sprintf("fieldmask_utils-single-generation-%s", c.name), func(b *testing.B) {
			var innerOut any

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				innerOut = reflect.New(c.out).Interface()
				err = nil

				m, err = fieldmask_utils.MaskFromProtoFieldMask(&fm, pascalCase)
				if err != nil {
					b.FailNow()
				}
				err = fieldmask_utils.StructToStruct(m, &simpleMessage, &innerOut)
				if err != nil {
					b.FailNow()
				}
			}

			out = innerOut
			outerErr = err
		})

		b.Run(fmt.Sprintf("fieldmask_utils-unique-generations-%s", c.name), func(b *testing.B) {
			var (
				innerOut      any
				inputMessages = make([]pb.BenchmarkTest, b.N)
				inputMasks    = make([]fieldmaskpb.FieldMask, b.N)
			)

			for j := 0; j < b.N; j++ {
				inputMessages[j] = generateSimpleMessage()
				inputMasks[j] = generateMask()
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				innerOut = reflect.New(c.out).Interface()
				err = nil

				m, err = fieldmask_utils.MaskFromProtoFieldMask(&inputMasks[j], pascalCase)
				if err != nil {
					b.FailNow()
				}
				err = fieldmask_utils.StructToStruct(m, &inputMessages[j], &innerOut)
				if err != nil {
					b.FailNow()
				}
			}

			out = innerOut
			outerErr = err
		})

		b.Run(fmt.Sprintf("Apply-single-generation-%s", c.name), func(b *testing.B) {
			var innerOut any

			b.ResetTimer()

			for j := 0; j < b.N; j++ {
				innerOut = reflect.New(c.out).Interface()

				err = field_mask.Apply(&fm, &simpleMessage, innerOut)
				if err != nil {
					b.FailNow()
				}
			}
		})

		b.Run(fmt.Sprintf("Apply-unique-generations-%s", c.name), func(b *testing.B) {
			var (
				innerOut      any
				inputMessages = make([]pb.BenchmarkTest, b.N)
				inputMasks    = make([]fieldmaskpb.FieldMask, b.N)
			)

			for j := 0; j < b.N; j++ {
				inputMessages[j] = generateSimpleMessage()
				inputMasks[j] = generateMask()
			}

			b.ResetTimer()

			for j := 0; j < b.N; j++ {
				innerOut = reflect.New(c.out).Interface()

				err = field_mask.Apply(&inputMasks[j], &inputMessages[j], innerOut)
				if err != nil {
					b.FailNow()
				}
			}

		})
	}

}
