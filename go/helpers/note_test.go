package helpers

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/gomidi/midi/reader"
)

type TestMiscMessage struct{}
type TestNoteMessage struct{}

func (m *TestMiscMessage) Raw() []byte {
	return []byte{0x00, 0x00, 0x00}
}

func (m *TestMiscMessage) String() string {
	return "meta.TrackSequenceName: \"Track\""
}

func (m *TestNoteMessage) Raw() []byte {
	return []byte{0x00, 0x00, 0x00}
}

func (m *TestNoteMessage) String() string {
	return "channel.Note"
}

// TestCaptureMiscMessage tests
func TestCaptureMiscMessage(t *testing.T) {
	ClearEvents()

	p := reader.Position{Track: 0, DeltaTicks: 0, AbsoluteTicks: 0}
	CaptureMiscMessage(&p, &TestMiscMessage{})
	CaptureMiscMessage(&p, &TestNoteMessage{})

	assert.Equal(t, 1, len(Events))
	assert.Equal(t, []byte{0x00, 0x00, 0x00}, Events[0].(MiscEvent).Message.Raw())
	assert.Equal(t, "meta.TrackSequenceName: \"Track\"", Events[0].(MiscEvent).Message.String())
	assert.Equal(t, p, Events[0].(MiscEvent).Position)
}

// TestCaptureMiscMessage tests
func TestCaptureNoteMessage(t *testing.T) {
	ClearEvents()

	p := reader.Position{Track: 0, DeltaTicks: 0, AbsoluteTicks: 0}
	n := NoteEvent{Channel: 0, Key: 0, Velocity: 0, Position: p}
	CaptureNoteMessage(&p, n.Channel, n.Key, n.Velocity)

	assert.Equal(t, 1, len(Events))
	assert.Equal(t, n.Channel, Events[0].(NoteEvent).Channel)
	assert.Equal(t, n.Key, Events[0].(NoteEvent).Key)
	assert.Equal(t, n.Velocity, Events[0].(NoteEvent).Velocity)
	assert.Equal(t, p, Events[0].(NoteEvent).Position)
}

func TestInvertNote(t *testing.T) {
	tonic := uint8(50)

	// Test the tonal center
	invertedNote := InvertNote(50, tonic)
	assert.Equal(t, uint8(50), invertedNote)

	// Test transpose down
	invertedNote = InvertNote(48, tonic)
	assert.Equal(t, uint8(52), invertedNote)

	// Test transpose up
	invertedNote = InvertNote(52, tonic)
	assert.Equal(t, uint8(48), invertedNote)
}

func TestLookupMidiKey(t *testing.T) {
	_ = json.Unmarshal([]byte(toneJson), &toneMap)

	for k, v := range toneMap {
		key, err := LookupMidiKey(k)
		assert.Equal(t, v, key)
		assert.Nil(t, err)
	}

	_, err := LookupMidiKey("")
	assert.NotNil(t, err)

	toneMap = nil
	_, err = LookupMidiKey("c2")
	assert.Nil(t, err)
}

func TestReciprocateNote(t *testing.T) {
	tonic := uint8(50)

	// Test the tonal center
	reciprocatedNote := ReciprocateNote(50, tonic)
	assert.Equal(t, uint8(57), reciprocatedNote)

	reciprocatedNote = ReciprocateNote(52, tonic)
	assert.Equal(t, uint8(55), reciprocatedNote)

	reciprocatedNote = ReciprocateNote(48, tonic)
	assert.Equal(t, uint8(59), reciprocatedNote)

	reciprocatedNote = ReciprocateNote(30, 7)
	assert.Equal(t, uint8(1), reciprocatedNote)
}
