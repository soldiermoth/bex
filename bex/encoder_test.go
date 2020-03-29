package bex_test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/soldiermoth/bex/bex"
)

func BenchmarkEncoder(b *testing.B) {
	fixture, err := ioutil.ReadFile("../examples/generic.xml")
	if err != nil {
		b.Fatalf("could not read fixture file %q", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dec := xml.NewDecoder(bytes.NewReader(fixture))
		err := bex.Encode(dec, ioutil.Discard)
		if err != nil {
			b.Fatalf("problem running Encode %q", err)
		}
	}
}
