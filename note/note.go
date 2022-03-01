package note

import "gitlab.com/gomidi/midi/reader"

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

/*

C  48
C# 49
D  50
D# 51
E  52
F  53
F# 54
G  55
G# 56
A  57
A# 58
B  59

C  60

C# 61
D  62
D# 63
E  64
F  65
F# 66
G  67
G# 68
A  69
A# 70
B  71
C  72

*/

func TransposeNote(key uint8) uint8 {
	// first, disregard the octaves
	if key%12 == 0 {
		return key
	}

	if key > 60 {
		return 60 - (key - 60)
	} else {
		return 60 + (60 - key)
	}
}
