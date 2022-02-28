package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
)

type Note struct {
	Channel, Key, Velocity uint8
	Position               *reader.Position
}

var NoteEvents []Note

func NoteCapture(p *reader.Position, channel uint8, key uint8, velocity uint8) {
	NoteEvents = append(NoteEvents, Note{
		Channel:  channel,
		Key:      key,
		Velocity: velocity,
		Position: p,
	})
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		//fmt.Println(err)
		return
	}

	f := filepath.Join(dir, "clap_v1.mid")

	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		reader.NoteOn(NoteCapture),
		reader.NoteOff(NoteCapture),
	)

	err = reader.ReadSMFFile(rd, f)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", f)
	}

	wf := filepath.Join(dir, "test.mid")
	err = writer.WriteSMF(wf, 1, func(wr *writer.SMF) error {
		wr.SetChannel(0)

		for _, n := range NoteEvents {
			wr.SetDelta(n.Position.DeltaTicks)

			if n.Velocity == 0 {
				writer.NoteOff(wr, n.Key+6)
			} else {
				writer.NoteOn(wr, n.Key+6, n.Velocity)
			}
		}

		writer.EndOfTrack(wr)

		return nil
	})

	if err != nil {
		fmt.Printf("could not write SMF file %v\n", f)
		return
	}

	// to disable logging, pass mid.NoLogger() as option
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
