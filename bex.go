package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/soldiermoth/bex/bexlib"
)

func main() {
	var (
		enc    = bexlib.NewEncoder(os.Stdout)
		indent = flag.String("indent", bexlib.DefaultIndent, "Controls how much of an indent to give it")
		r      io.Reader
		err    error
	)
	flag.IntVar(&enc.AttrsBeforePad, "attrs-before-pad", enc.AttrsBeforePad, "How many attrs previous node has before breaking value onto it's own line")
	flag.Var(&FlagColorizer{Colorizer: &enc.ColorAttr}, "color-attr", "Attribute Color")
	flag.Var(&FlagColorizer{Colorizer: &enc.ColorComment}, "color-comment", "Comment Color")
	flag.Var(&FlagColorizer{Colorizer: &enc.ColorElement}, "color-element", "Element Color")
	flag.Parse()
	enc.Indent("", *indent)
	args := flag.Args()
	// Detect if stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		r = os.Stdin
		if len(args) > 0 {
			log.Fatal("no arguments expected when reading from stdin")
		}
	} else if len(args) == 1 {
		if r, err = os.Open(args[0]); err != nil {
			log.Fatalf("could not open file %q err=%q", args[0], err)
		}
	} else {
		log.Fatal("expected 1 argument of the xml file to process or to read from stdin")
	}

	if err := bexlib.Pipe(xml.NewDecoder(r), enc); err != nil {
		log.Fatal(err)
	}
	enc.Flush()
	fmt.Fprintln(os.Stdout)
}

// FlagColorizer sets color from a command line string
type FlagColorizer struct{ *bexlib.Colorizer }

func (f *FlagColorizer) String() string {
	if f == nil || f.Colorizer == nil || len(*f.Colorizer) == 0 {
		return "No Color"
	}
	return f.S("Example")
}

// Set fufills flag.Value
func (f *FlagColorizer) Set(s string) error {
	if s == "" {
		return nil
	}
	color, err := bexlib.ParseColor(s)
	if err != nil {
		return err
	}
	*f.Colorizer = bexlib.NewColorizer(color)
	return nil
}
