package main

import (
	"flag"
	"levy-generator/helpers"
	"log"
	"os"
)

func main() {
	// Populate flags from command line
	debug := flag.Bool("debug", false, "debug input file")
	inputFile := flag.String("if", "", "input file location")
	tonalCenter := flag.String("gravity", "", "tonal center")
	flag.Parse()

	if *inputFile == "" || *tonalCenter == "" {
		log.Println("Usage: main.go -if \"input-file.mid\" -gravity \"c-2\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Find the MIDI key for the user-supplied tonal center
	tonalCenterMidiKey, err := helpers.LookupMidiKey(*tonalCenter)
	if err != nil {
		log.Fatal(err)
	}

	// Read the input file
	rd, err := helpers.ReadFile(*inputFile, *debug)
	if err != nil {
		log.Fatal(err)
	}

	// Write the output file if debug is false
	if !*debug {
		helpers.WriteFile(rd, tonalCenterMidiKey, helpers.ConstructOutputFileName(*inputFile))
	}
}
