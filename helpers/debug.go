package helpers

import (
	"fmt"
	"os"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"
)

func HeaderDebug(h smf.Header) {
	fmt.Println(h)
}

func NoteDebug(p *reader.Position, channel uint8, key uint8, velocity uint8) {
	fmt.Printf("Note @ %v: Channel %v, Key %v, Velocity %v\n", p, channel, key, velocity)
}

func DebugFile(fileName string) {
	rd := reader.New(
		reader.NoLogger(),
		reader.NoteOn(NoteDebug),
		reader.NoteOff(NoteDebug),
		reader.SMFHeader(HeaderDebug),
	)
	err := reader.ReadSMFFile(rd, fileName)
	if err != nil {
		fmt.Printf("could not read SMF file %v\n", fileName)
		os.Exit(1)
	}
}
