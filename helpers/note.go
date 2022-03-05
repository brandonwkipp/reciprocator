package helpers

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
)

//go:embed keymap.json
var toneJson string
var toneMap map[string]uint8
var Events []interface{}

// MiscEvent is a struct for capturing misc midi messages
type MiscEvent struct {
	Message  midi.Message
	Position reader.Position
}

// NoteEvent is a struct for capturing midi note messages
type NoteEvent struct {
	Channel, Key, Velocity uint8
	Position               reader.Position
}

// CaptureMiscMessage captures any message that is not a NoteOn/NoteOff message
func CaptureMiscMessage(p *reader.Position, m midi.Message) {
	if !strings.Contains(m.String(), "channel.Note") {
		Events = append(Events, MiscEvent{
			Message:  m,
			Position: *p,
		})
	}
}

// CaptureNoteMessage captures NoteOn/NoteOff messages
func CaptureNoteMessage(p *reader.Position, channel, key, velocity uint8) {
	Events = append(Events, NoteEvent{
		Channel:  channel,
		Key:      key,
		Velocity: velocity,
		Position: *p,
	})
}

// ClearEvents clears the Events slice
func ClearEvents() {
	Events = []interface{}{}
}

// InvertNote transposes a note to its mirror image with respect to the tonal center
func InvertNote(key, tonalCenter uint8) uint8 {
	// ignore the tonal center
	if key == tonalCenter {
		return key
	}

	// handle inversion based on distance from tonal center
	if key > tonalCenter {
		return tonalCenter - (key - tonalCenter) // transpose down
	} else {
		return tonalCenter + (tonalCenter - key) // transpose up
	}
}

// LookupMidiKey returns the midi key value from keymap.json of a given tonal center
func LookupMidiKey(tonalCenter string) (uint8, error) {
	// prepare tone for lookup
	lowercase := strings.ToLower(tonalCenter)
	err := json.Unmarshal([]byte(toneJson), &toneMap)
	if err != nil {
		return 0, errors.New("could not intialize toneMap")
	}

	// check if tone is in map
	if midiKey, ok := toneMap[lowercase]; ok {
		return midiKey, nil
	} else {
		return 0, errors.New("could not find midi key for " + tonalCenter)
	}
}
