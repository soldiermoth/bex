package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/soldiermoth/bex/bex"
)

func main() {
	var (
		d = xml.NewDecoder(os.Stdin)
		e = xml.NewEncoder(os.Stdout)
	)
	e.Indent("", "  ")
	if err := bex.Encode(d, e); err != nil {
		log.Fatal(err)
	}
	e.Flush()
	fmt.Fprintln(os.Stdout)
}
