use rimd::{MidiMessage, TrackEvent};

struct MiscEvent {
	message: MidiMessage,
	position: u8,
}

struct NoteEvent {
	channel: u8,
	key: u8,
	position: String,
	velocity: u8,
}

pub fn capture_note_message(e: TrackEvent) {
	println!("{}", e);
}

pub fn invert_note(key: u8, tonal_center: u8) -> u8 {
	// ignore the tonal center
	if key == tonal_center {
		return key
	}

	// handle inversion based on distance from tonal center
	if key > tonal_center {
		return tonal_center - (key - tonal_center) // transpose down
	} else {
		return tonal_center + (tonal_center - key) // transpose up
	}
}

pub fn reciprocate_note(key: u8, tonal_center: u8) -> u8 {
	// 3.5 half steps is half the distance to the 5th from the root
	// This only seems to work for the major scale for some reason??
	let axis = tonal_center as f32 + 3.5;
	let key_distance = axis - key as f32;
	(axis + key_distance) as u8
}