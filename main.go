package main

import (
	"flag"
	"levy-generator/helpers"
	"log"
	"os"
)

func main() {
	var debug bool
	var inputFile string
	var tonalCenter string

	// flags
	flag.BoolVar(&debug, "debug", false, "debug input file")
	flag.StringVar(&inputFile, "if", "", "input file location")
	flag.StringVar(&tonalCenter, "gravity", "", "tonal center")
	flag.Parse()

	if inputFile == "" || tonalCenter == "" {
		log.Println("Usage: main.go -if \"input-file.mid\" -gravity \"c-2\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	tonalCenterMidiKey, err := helpers.LookupMidiKey(tonalCenter)
	if err != nil {
		log.Fatal(err)
	}

	rd, err := helpers.ReadFile(inputFile, debug)
	if err != nil {
		log.Fatal(err)
	}

	if !debug {
		helpers.WriteFile(rd, tonalCenterMidiKey, helpers.ConstructOutputFileName(inputFile))
	}
}
