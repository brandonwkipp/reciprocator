package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfwriter"
	"gitlab.com/gomidi/midi/writer"
)

// ConstructOutputFileName constructs a new output file name based on the input file name
func ConstructOutputFileName(inputFile string) string {
	ext := filepath.Ext(inputFile)
	pos := strings.LastIndex(inputFile, ext)
	return inputFile[:pos] + "-negative" + ext
}

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

// ReadFile reads a standard midi file and returns a Reader or logs output about the file if debug is set to true
func ReadFile(f string, debug bool) (*reader.Reader, error) {
	// Pass the debug functions to the reader instead of the default functions
	if debug {
		DebugSMF(f)
		return nil, nil
	}

	rd := reader.New(
		reader.NoLogger(),
		reader.Each(CaptureMiscMessage),
		reader.NoteOn(CaptureNoteMessage),
		reader.NoteOff(CaptureNoteMessage),
	)

	err := reader.ReadSMFFile(rd, f)
	if err != nil {
		log.Fatalf("could not read SMF file %v", f)
	}

	return rd, nil
}

// WriteFile writes an inverted set of notes to a new standard midi file
func WriteFile(rd *reader.Reader, tonalCenterMidiKey uint8, outputFile string) {
	dir := ""
	wf := filepath.Join(dir, outputFile)
	err := writer.WriteSMF(wf, rd.Header().NumTracks, func(wr *writer.SMF) error {
		for _, e := range Events {
			switch e := e.(type) {
			case MiscEvent:
				wr.SetDelta(e.Position.DeltaTicks)
				wr.Write(e.Message)
			case NoteEvent:
				wr.SetDelta(e.Position.DeltaTicks)
				transposedNote := InvertNote(e.Key, tonalCenterMidiKey)

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
		log.Fatalf("could not write SMF file %v", wf)
	}
}
