package main

import (
	"flag"
	"log"
)

type flagValues struct {
	isDecode bool
	fileName string
}

func main() {
	vals := &flagValues{}
	flag.BoolVar(&vals.isDecode, "d", false, "Decode a given announcement file")
	flag.StringVar(&vals.fileName, "f", "", "File name")
	flag.Parse()

	if err := mainErr(vals); err != nil {
		log.Fatalln(err)
	}
}

func mainErr(vals *flagValues) error { _ = "STUB: not implemented"; return nil }
