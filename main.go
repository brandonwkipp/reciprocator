package main

import (
	"flag"
	"log"
	"os"

	"reciprocator/helpers"
)

func main() {
	// Populate flags from command line
	var debug, help, invert bool
	var inputFile, tonic string
	flag.BoolVar(&debug, "debug", false, "debug input file")
	flag.BoolVar(&help, "help", false, "show help")
	flag.StringVar(&inputFile, "if", "", "input file location")
	flag.BoolVar(&invert, "invert", false, "invert note rather than reciprocate it")
	flag.StringVar(&tonic, "tonic", "", "the tonal center")
	flag.Parse()

	// Show help if requested and exit
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// Read the input file
	if inputFile == "" {
		log.Fatal("Usage: main.go -if \"input-file.mid\"")
	}
	rd, err := helpers.ReadFile(inputFile, debug)
	if err != nil {
		log.Fatal(err)
	}

	// Rewrite the input file if debug is not enabled
	if !debug && tonic == "" {
		log.Fatal("Usage: main.go -if \"input-file.mid\" -tonic \"c2\"")
	}

	if !debug {
		// Find the MIDI key for the user-supplied tonic
		tonicMidiKey, err := helpers.LookupMidiKey(tonic)
		if err != nil {
			log.Fatal(err)
		}

		// Write the output file if debug is false
		helpers.WriteFile(rd, tonicMidiKey, helpers.ConstructOutputFileName(inputFile, invert), invert)
	}
}
