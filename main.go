package main

import (
	"flag"
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
		log.Println("Usage: main.go -if \"input-file.mid\" -gravity \"c-2\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := getFilePath(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// handle debug
	if debug {
		helpers.DebugSMF(f)
		return
	}

	tonalCenterMidiKey, err := helpers.LookupMidiKey(tonalCenter)
	if err != nil {
		log.Fatal(err)
	}

	rd := reader.New(
		reader.NoLogger(),
		reader.Each(helpers.CaptureMiscMessage),
		reader.NoteOn(helpers.CaptureNoteMessage),
		reader.NoteOff(helpers.CaptureNoteMessage),
	)

	err = reader.ReadSMFFile(rd, f)
	if err != nil {
		log.Fatalf("could not read SMF file %v", f)
	}

	dir := ""

	wf := filepath.Join(dir, "test.mid")
	err = writer.WriteSMF(wf, rd.Header().NumTracks, func(wr *writer.SMF) error {
		for _, e := range helpers.Events {
			switch e := e.(type) {
			case helpers.MiscEvent:
				wr.SetDelta(e.Position.DeltaTicks)
				wr.Write(e.Message)
			case helpers.NoteEvent:
				wr.SetDelta(e.Position.DeltaTicks)
				transposedNote := helpers.InvertNote(e.Key, tonalCenterMidiKey)

				if e.Velocity == 0 {
					writer.NoteOff(wr, transposedNote)
				} else {
					wr.SetDelta(e.Position.DeltaTicks)
					writer.NoteOn(wr, transposedNote, e.Velocity)
				}
			}
		}
		return nil
	}, smfwriter.TimeFormat(rd.Header().TimeFormat))

	if err != nil {
		log.Fatalf("could not write SMF file %v", f)
	}
}
