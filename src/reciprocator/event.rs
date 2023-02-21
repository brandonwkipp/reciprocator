use rimd::{MidiMessage, Status};

pub fn handle_message(msg: MidiMessage, invert: bool, tonal_center_midi_key: u8) -> MidiMessage {
	let channel = match msg.channel() {
		Some(x) => x,
		None => 0
	};

	match msg.status() {
		Status::NoteOff => MidiMessage::note_off(handle_operation(msg.data[1], invert, tonal_center_midi_key), msg.data[2], channel),
		Status::NoteOn => MidiMessage::note_on(handle_operation(msg.data[1], invert, tonal_center_midi_key), msg.data[2], channel),
		_ => msg,
	}
}

fn handle_operation(note: u8, invert: bool, tonal_center_midi_key: u8) -> u8 {
	if !invert {
		return reciprocate_note(note, tonal_center_midi_key)
	} else {
		return invert_note(note, tonal_center_midi_key)
	}
}

fn invert_note(key: u8, tonal_center: u8) -> u8 {
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

fn reciprocate_note(key: u8, tonal_center: u8) -> u8 {
	// 3.5 half steps is half the distance to the 5th from the root
	// This only seems to work for the major scale for some reason??
	let axis = tonal_center as f32 + 3.5;
	let key_distance = axis - key as f32;
	(axis + key_distance) as u8
}