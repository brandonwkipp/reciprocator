package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"
)

// DebugMisc prints out any message that is not a NoteOn/NoteOff message
func DebugMisc(p *reader.Position, m midi.Message) {
	if !strings.Contains(m.String(), "channel.Note") {
		fmt.Printf("DEBUG MISC: %v\n", m)
	}
}

// DebugNote prints out NoteOn/NoteOff messages
func DebugNote(p *reader.Position, channel uint8, key uint8, velocity uint8) {
	fmt.Printf("DEBUG NOTE: Position %v, Channel %v, Key %v, Velocity %v\n", p, channel, key, velocity)
}

// DebugSMF prints out the contents of a standard midi file
func DebugSMF(fileName string) {
	rd := reader.New(
		reader.NoLogger(),
		reader.Each(DebugMisc),
		reader.NoteOff(DebugNote),
		reader.NoteOn(DebugNote),
	)

	err := reader.ReadSMFFile(rd, fileName)
	if err != nil {
		log.Printf("could not read SMF file %v\n", fileName)
		os.Exit(1)
	}
}

// DebugSMFHeader prints out the contents of a standard midi file header
func DebugSMFHeader(h smf.Header) {
	fmt.Println(h)
}
