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

func TransposeNote(key, tonalCenter uint8) uint8 {
	// first, disregard the octaves
	if key%12 == 0 {
		return key
	}

	if key > tonalCenter {
		return tonalCenter - (key - tonalCenter)
	} else {
		return tonalCenter + (tonalCenter - key)
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
