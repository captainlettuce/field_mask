package bench

import (
	"github.com/captainlettuce/field_mask"
	"github.com/captainlettuce/field_mask/testing/pb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"testing"
)

var (
	out      any
	outerErr error
)

func BenchmarkApplyPointerReceiver(b *testing.B) {
	var (
		simpleMessage pb.BenchmarkTest
		o             outStructPointer
		err           error
	)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		simpleMessage = generateSimpleMessage()
		fm := generateMask()
		o = outStructPointer{}
		b.StartTimer()
		err = field_mask.Apply(&fm, &simpleMessage, &o)
	}

	out = o
	outerErr = err
}

func BenchmarkApplyRawReceiver(b *testing.B) {
	var (
		simpleMessage pb.BenchmarkTest
		o             outStruct
		err           error
	)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		simpleMessage = generateSimpleMessage()
		fm := generateMask()
		o = outStruct{}
		b.StartTimer()
		err = field_mask.Apply(&fm, &simpleMessage, &o)
	}

	out = o
	outerErr = err
}

func BenchmarkMemoryAllocations(b *testing.B) {
	msg := pb.BenchmarkTest{
		A: "test",
		B: 1.1,
		C: 1,
		D: ref(true),
	}
	mask := fieldmaskpb.FieldMask{Paths: []string{"a", "b", "c", "d"}}
	var (
		err error
		//	out any
	)

	cases := []struct {
		name string
		out  any
	}{
		{
			name: "pointer receivers",
			out:  &outStructPointer{},
		},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err = field_mask.Apply(&mask, &msg, c.out)
				if err != nil {
					b.Fatal(err.Error())
				}
			}
		})
	}
}
