package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"

	"levy-generator/helpers"
)

func getFilePath(fileName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fileName), nil
}

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
		fmt.Println("Usage: main.go -if \"input-file.mid\" -gravity \"c-2\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := getFilePath(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// handle debug
	if debug {
		helpers.DebugSMFFile(f)
		return
	}

	tonalCenterMidiKey, err := helpers.TranslateTonalCenterToMidiKey(tonalCenter)
	if err != nil {
		fmt.Println("wtf")
	}

	rd := reader.New(
		reader.NoLogger(),
		reader.NoteOn(helpers.NoteCapture),
		reader.NoteOff(helpers.NoteCapture),
	)

	err = reader.ReadSMFFile(rd, f)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}

	dir := ""

	wf := filepath.Join(dir, "test.mid")
	err = writer.WriteSMF(wf, 2, func(wr *writer.SMF) error {
		wr.SetChannel(0)

		writer.TrackSequenceName(wr, "Transport")
		writer.Instrument(wr, "Hardware Interface II")
		writer.DeprecatedPort(wr, 0)
		writer.TempoBPM(wr, 120.00)
		writer.TimeSig(wr, 4, 4, 24, 8)

		wr.SetDelta(61440)
		writer.EndOfTrack(wr)

		writer.DeprecatedPort(wr, 4)
		writer.TrackSequenceName(wr, "MIDI Out Ch1")
		writer.Instrument(wr, "MIDI Out Ch1")

		for _, n := range helpers.NoteEvents {
			wr.SetDelta(n.Position.DeltaTicks)
			transposedNote := helpers.TransposeNote(n.Key, tonalCenterMidiKey)

			if n.Velocity == 0 {
				writer.NoteOff(wr, transposedNote)
			} else {
				wr.SetDelta(n.Position.DeltaTicks)
				writer.NoteOn(wr, transposedNote, n.Velocity)
			}
		}

		writer.EndOfTrack(wr)

		return nil
	}, smfwriter.TimeFormat(rd.Header().TimeFormat))

	if err != nil {
		fmt.Printf("could not write SMF file %v\n", f)
		return
	}
}
