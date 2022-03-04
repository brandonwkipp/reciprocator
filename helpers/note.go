package helpers

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/gomidi/midi/reader"
)

//go:embed keymap.json
var toneJson string
var toneMap map[string]uint8
var NoteEvents []NoteEvent

// A NoteEvent contains all properties of any NoteOn/NoteOff messages rendered from an SMF
type NoteEvent struct {
	Channel, Key, Velocity uint8
	Position               reader.Position
}

func NoteCapture(p *reader.Position, channel uint8, key uint8, velocity uint8) {
	NoteEvents = append(NoteEvents, NoteEvent{
		Channel:  channel,
		Key:      key,
		Velocity: velocity,
		Position: *p,
	})
}

// TransposeNote transposes a note to its mirror image
func TransposeNote(key, tonalCenter uint8) uint8 {
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

func TranslateTonalCenterToMidiKey(tone string) (uint8, error) {
	lowercase := strings.ToLower(tone)

	err := json.Unmarshal([]byte(toneJson), &toneMap)
	if err != nil {
		return 0, errors.New("could not intialize toneMap")
	}

	if value, ok := toneMap[lowercase]; ok {
		return value, nil
	}

	return 0, errors.New("couldn't do it")
}
