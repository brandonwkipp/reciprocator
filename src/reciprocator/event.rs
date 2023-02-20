use rimd::{MidiMessage, Status};

pub fn handle_message(msg: MidiMessage, invert: bool, tonal_center_midi_key: u8) -> MidiMessage {
	let channel = match msg.channel() {
		Some(x) => x,
		None => 0
	};

	let altered_note: u8;
	if !invert {
		altered_note = reciprocate_note(msg.data[2], tonal_center_midi_key);
	} else {
		altered_note = invert_note(msg.data[2], tonal_center_midi_key);
	}

	match msg.status() {
		Status::NoteOn => MidiMessage::note_on(msg.data[1], altered_note, channel),
		Status::NoteOn => MidiMessage::note_off(msg.data[1], altered_note, channel),
		_ => msg,
	}
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