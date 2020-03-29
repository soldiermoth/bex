package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/soldiermoth/bex/bexlib"
)

func main() {
	if len(os.Args) > 1 {
		log.Fatal("Too many arguments")
	}
	var (
		d = xml.NewDecoder(os.Stdin)
	)
	if len(os.Args) == 1 {

	}
	if err := bexlib.Encode(d, os.Stdout); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout)
}
