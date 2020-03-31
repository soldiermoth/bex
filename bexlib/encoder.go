package bexlib

import (
	"bytes"
	"encoding/xml"
	"io"
)

// Encoder wraps xml.Encoder
// * color
// * padding around comments
// * padding around char data
type Encoder struct {
	ColorElement, ColorAttr, ColorComment Colorizer
	AttrsBeforePad                        int
	// internal tracking data
	state state
	enc   *xml.Encoder
}

// NewEncoder creates a new encoder with some sensible defaults
func NewEncoder(out io.Writer) *Encoder {
	enc := &Encoder{
		ColorElement:   DefaultColorElement,
		ColorAttr:      DefaultColorAttr,
		ColorComment:   DefaultColorComment,
		AttrsBeforePad: DefaultAttrsBeforePad,
		enc:            xml.NewEncoder(out),
	}
	enc.Indent("", DefaultIndent)
	return enc
}

// Indent adds an indent to the encoder
func (e *Encoder) Indent(prefix, indent string) {
	e.state.indent = []byte(indent)
	e.enc.Indent(prefix, indent)
}

// Flush wraps underlying encoder's Flush
func (e *Encoder) Flush() error { return e.enc.Flush() }

func (e *Encoder) start(t xml.StartElement) {
	if e.state.begun && !e.state.begunNodes {
		e.string("\n")
	}
	e.state.begunNodes = true
	e.state.depth++
	e.state.directParent = &t
	t.Name.Local = e.ColorElement.S(t.Name.Local)
	for i, a := range t.Attr {
		t.Attr[i].Name.Local = e.ColorAttr.S(a.Name.Local)
	}
	e.token(t)
}

func (e *Encoder) end(t xml.EndElement) {
	e.state.depth--
	e.state.directParent = nil
	t.Name.Local = e.ColorElement.S(t.Name.Local)
	e.token(t)
}

func (e *Encoder) comment(t xml.Comment) {
	t = e.ColorComment.B(t)
	if e.state.begun {
		e.string("\n")
	}
	e.padding(e.state.depth)
	e.token(t)
}

func (e *Encoder) chardata(t xml.CharData) {
	t = bytes.TrimSpace(t)
	if len(t) == 0 {
		return
	}
	if p := e.state.directParent; p != nil && len(p.Attr) > e.AttrsBeforePad {
		e.string("\n")
		e.padding(e.state.depth)
		defer e.padding(e.state.depth - 1)
		defer e.string("\n")
	}
	e.token(t)
}

func (e *Encoder) token(t xml.Token) { e.enc.EncodeToken(t) }
func (e *Encoder) bytes(b []byte)    { e.token(xml.CharData(b)) }
func (e *Encoder) string(s string)   { e.bytes([]byte(s)) }
func (e *Encoder) padding(i int)     { e.bytes(bytes.Repeat(e.state.indent, i)) }

// EncodeToken wraps xml EncodeToken & adds customizations
func (e *Encoder) EncodeToken(token xml.Token) {
	switch t := token.(type) {
	case xml.StartElement:
		e.start(t)
	case xml.EndElement:
		e.end(t)
	case xml.Comment:
		e.comment(t)
	case xml.CharData:
		e.chardata(t)
	default:
		e.token(token)
	}
	e.state.begun = true
}

type state struct {
	begun      bool
	begunNodes bool
	depth      int
	indent     []byte
	// only accurate for first child node
	directParent *xml.StartElement
}
