package bexlib_test

import (
	"testing"

	"github.com/soldiermoth/bex/bexlib"
)

func BenchmarkColor_S(b *testing.B) {
	var (
		in     = "hello world"
		result string
		red    = bexlib.NewColorizer(bexlib.Red)
	)
	for n := 0; n < b.N; n++ {
		result = red.S(in)
	}
	if result == "" {
		b.Fatal("Should be non-empty")
	}
}

func BenchmarkColor_B(b *testing.B) {
	var (
		in     = []byte("hello world")
		result []byte
		red    = bexlib.NewColorizer(bexlib.Red)
	)
	for n := 0; n < b.N; n++ {
		result = red.B(in)
	}
	if len(result) == 0 {
		b.Fatal("Should be non-empty")
	}
}
