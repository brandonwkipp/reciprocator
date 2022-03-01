package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"

	"levy-generator/note"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "if", "", "# of iterations")
	flag.Parse()

	if inputFile == "" {
		fmt.Println("Usage: main.go -if \"file.mid\"")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	f := filepath.Join(dir, inputFile)

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		reader.NoteOn(note.NoteCapture),
		reader.NoteOff(note.NoteCapture),
	)

	err = reader.ReadSMFFile(rd, f)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}

	wf := filepath.Join(dir, "test.mid")
	err = writer.WriteSMF(wf, 1, func(wr *writer.SMF) error {
		wr.SetChannel(0)

		for _, n := range note.NoteEvents {
			wr.SetDelta(n.Position.DeltaTicks)

			transposedNote := note.TransposeNote(n.Key)

			if n.Velocity == 0 {
				writer.NoteOff(wr, transposedNote)
			} else {
				writer.NoteOn(wr, transposedNote, n.Velocity)
			}
		}

		writer.EndOfTrack(wr)

		return nil
	})

	if err != nil {
		fmt.Printf("could not write SMF file %v\n", f)
		return
	}

	wrd := reader.New()

	err = reader.ReadSMFFile(wrd, f)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}

	log.Print("\n")

	err = reader.ReadSMFFile(wrd, wf)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}
}
