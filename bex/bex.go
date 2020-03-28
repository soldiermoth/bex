package bex

import (
	"bytes"
	"encoding/xml"
	"io"
)

// Control the color of xml types
var (
	Element = NewColorizer(Blue)
	Attr    = NewColorizer(Red)
	Comment = NewColorizer(Yellow)
)

func colorizeToken(i xml.Token) xml.Token {
	switch t := i.(type) {
	default:
		return t
	case xml.StartElement:
		t.Name.Local = Element.S(t.Name.Local)
		for i, a := range t.Attr {
			t.Attr[i].Name.Local = Attr.S(a.Name.Local)
		}
		return t
	case xml.EndElement:
		t.Name.Local = Element.S(t.Name.Local)
		return t
	case xml.Comment:
		b := Comment.B(t)
		return xml.Comment(b)
	case xml.CharData:
		c := bytes.TrimSpace(t)
		return xml.CharData(c)
	}
}

// XMLTokener input to read from
type XMLTokener interface{ Token() (xml.Token, error) }

// XMLTokenEncoder output to write to
type XMLTokenEncoder interface{ EncodeToken(xml.Token) error }

// Encode pipes from in to out while adding color along the way
func Encode(in XMLTokener, out XMLTokenEncoder) error {
	for {
		tok, err := in.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		out.EncodeToken(colorizeToken(tok))
	}
}
