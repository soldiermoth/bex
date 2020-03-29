package bexlib

import (
	"strconv"
)

const (
	// Reset is the control code for resetting the color
	Reset = "\x1b[0m"
)

// Color codes
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
)

// Color wraps the int value for a color
type Color int

// Control outputs the color's control code
func (c Color) Control() string { return "\x1b[0;" + strconv.Itoa(int(c)) + "m" }

// Colorizer aliases a control code
type Colorizer string

// NewColorizer creates a new colorizer from a color
func NewColorizer(c Color) Colorizer { return Colorizer(c.Control()) }

// S wraps a string with the control code
func (c Colorizer) S(s string) string { return string(c) + s + Reset }

// B wraps a byte slice with the control code
func (c Colorizer) B(b []byte) []byte {
	b = append([]byte(c), b...)
	b = append(b, Reset...)
	return b
}
