package bench

import (
	"github.com/captainlettuce/field_mask/testing/pb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"math/rand"
)

type outStructPointer struct {
	A             *string  `field_mask:"a"`
	B             *float32 `field_mask:"b"`
	C             *int32   `field_mask:"c"`
	UntaggedField *bool    `field_mask:"UntaggedField"`
}

type outStruct struct {
	A             string  `field_mask:"a"`
	B             float32 `field_mask:"b"`
	C             int32   `field_mask:"c"`
	UntaggedField bool
}

func generateMask() fieldmaskpb.FieldMask {
	fm := fieldmaskpb.FieldMask{}
	fields := []string{
		"a",
		"b",
		"c",
		"d",
	}
	for _, field := range fields {
		if rand.Intn(2) == 1 {
			fm.Paths = append(fm.Paths, field)
		}
	}

	return fm
}

func ref[T any](t T) *T {
	return &t
}

func generateSimpleMessage() pb.BenchmarkTest {
	return pb.BenchmarkTest{
		A: randomString(rand.Intn(10)),
		B: rand.Float32(),
		C: rand.Int31(),
		D: ref(true),
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func pascalCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	return string(append(t, lookupAndReplacePascalCaseWords(s, i)...))
}

// lookupAndReplacePascalCaseWords lookups for words in the string starting in position i
// and replaces the snake case format to PascalCase.
// Invariant: if the next letter is lower case, it must be converted
// to upper case.
// That is, we process a word at a time, where words are marked by _ or
// upper case letter. Digits are treated as words.
func lookupAndReplacePascalCaseWords(s string, i int) []byte {
	t := make([]byte, 0, 32)
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		t, i = appendLowercaseSequence(s, i, t)
	}
	return t
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// appendLowercaseSequence appends the lowercase sequence from s that begins at i into t
// returns the new t that contains all the chain of characters that should be lowercase
// and the new index where to start counting from.
func appendLowercaseSequence(s string, i int, t []byte) ([]byte, int) {
	for i+1 < len(s) && isASCIILower(s[i+1]) {
		i++
		t = append(t, s[i])
	}
	return t, i
}
