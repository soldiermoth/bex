package bexlib

import (
	"encoding/xml"
	"io"
)

const (
	// DefaultIndent sets default for XML Encoder
	DefaultIndent = "  "
	// AttrsBeforePad sets limit where it'll break values
	// onto new lines
	AttrsBeforePad = 5
)

// Control the color of xml types
var (
	Element = NewColorizer(LightBlue)
	Attr    = NewColorizer(LightRed)
	Comment = NewColorizer(Yellow)
)

// XMLTokener input to read from
type XMLTokener interface{ Token() (xml.Token, error) }

// Encode pipes from in to out while adding color along the way
func Encode(in XMLTokener, out io.Writer) error {
	enc := Encoder{enc: xml.NewEncoder(out)}
	enc.Indent("", DefaultIndent)
	defer enc.enc.Flush()
	for {
		tok, err := in.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		enc.EncodeToken(tok)
	}
}
