package bexlib

import (
	"encoding/xml"
	"io"
)

const (
	// DefaultIndent sets default for XML Encoder
	DefaultIndent = "  "
	// DefaultAttrsBeforePad sets default limit
	// where it'll break values onto new lines
	DefaultAttrsBeforePad = 5
)

// Control the color of xml types
var (
	DefaultColorElement = NewColorizer(LightBlue)
	DefaultColorAttr    = NewColorizer(LightRed)
	DefaultColorComment = NewColorizer(Yellow)
)

// Tokener input to read from
type Tokener interface{ Token() (xml.Token, error) }

// TokenEncoder encodes provided token
type TokenEncoder interface{ EncodeToken(xml.Token) }

// Pipe pipes from in to out while adding color along the way
func Pipe(in Tokener, out TokenEncoder) error {
	// enc := Encoder{enc: xml.NewEncoder(out)}
	// enc.Indent("", DefaultIndent)
	// defer enc.enc.Flush()
	for {
		tok, err := in.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		out.EncodeToken(tok)
	}
}
