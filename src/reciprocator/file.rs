use std::path::Path;

use rimd::{SMF};

use crate::reciprocator::note;

// ConstructOutputFileName constructs a new output file name based on the input file name
pub fn construct_output_filename(filename: &str, invert: bool) -> String {
	let ext = match Path::new(filename).extension() {
		Some(ext) => match ext.to_str() {
			Some(ext) => ext,
			None => return String::new(),
		},
		None => return String::new(),
	};

	let pos = match filename.rfind(ext) {
		Some(pos) => pos,
		None => return String::new(),
	};

	if invert {
		format!("{}-inverted.{}", &filename[..pos - 1], ext)
	} else {
		format!("{}-negative.{}", &filename[..pos - 1], ext)
	}
}

// DebugMisc prints out any message that is not a NoteOn/NoteOff message
// fn debug_misc(p: *reader.Position, m: midi.Message) {
// 	if !strings.Contains(m.String(), "channel.Note") {
// 		fmt.Printf("DEBUG MISC: %v\n", m)
// 	}
// }

// DebugNote prints out NoteOn/NoteOff messages
// fn debug_note(p: *reader.Position, channel: uint8, key: uint8, velocity: uint8) {
// 	fmt.Printf("DEBUG NOTE: Position %v, Channel %v, Key %v, Velocity %v\n", p, channel, key, velocity)
// }

// DebugSMF prints out the contents of a standard midi file
pub fn debug_smf(filename: &str) {
	// rd := reader.New(
	// 	reader.NoLogger(),
	// 	reader.Each(DebugMisc),
	// 	reader.NoteOff(DebugNote),
	// 	reader.NoteOn(DebugNote),
	// )

	let tracks = match SMF::from_file(Path::new(filename)) {
		Ok(x) => x.tracks,
		Err(e) => panic!("{}", e),
	};

	// for track in tracks {
	// 	match track.name {
	// 		Some(x) => println!("{x}"),
	// 		_ => {},
	// 	}
	// }

	for track in tracks {
		for event in track.events {
			note::capture_note_message(event);
		}
	}

	// if err != nil {
	// 	log.Printf("could not read SMF file %v\n", fileName)
	// 	os.Exit(1)
	// }
}

// DebugSMFHeader prints out the contents of a standard midi file header
// fn debug_smf_header(h: smf.Header) {
// 	fmt.Println(h)
// }

// ReadFile reads a standard midi file and returns a Reader or logs output about the file if debug is set to true
// fn read_file(f: string, debug: bool) -> *reader.Reader, error {
// 	// Pass the debug functions to the reader instead of the default functions
// 	// if debug {
// 	// 	debug_smf(f)
// 	// 	return nil, nil
// 	// }

// 	let tracks = match SMF::from_file(Path::new(filename)) {
// 		Ok(x) => x.tracks,
// 		Err(e) => panic!("{}", e),
// 	};

// 	for track in tracks {
// 		for event in events {

// 		}
// 	}

// 	// err := reader.ReadSMFFile(rd, f)
// 	// if err != nil {
// 	// 	log.Fatalf("could not read SMF file %v", f)
// 	// }

// 	// return rd, nil
// }

// WriteFile writes an inverted set of notes to a new standard midi file
// fn write_file(rd: *reader.Reader, tonal_center_midi_key: uint8, output_file: string, invert: bool) {
// 	dir := ""
// 	wf := filepath.Join(dir, output_file)
// 	err := writer.WriteSMF(wf, rd.Header().NumTracks, func(wr *writer.SMF) error {
// 		for _, e := range Events {
// 			switch e := e.(type) {
// 			case MiscEvent:
// 				wr.SetDelta(e.Position.DeltaTicks)
// 				wr.Write(e.Message)
// 			case NoteEvent:
// 				wr.SetDelta(e.Position.DeltaTicks)

// 				var transposed_note uint8
// 				if !invert {
// 					transposed_note = ReciprocateNote(e.Key, tonal_center_midi_key)
// 				} else {
// 					transposed_note = InvertNote(e.Key, tonal_center_midi_key)
// 				}

// 				if e.Velocity == 0 {
// 					writer.NoteOff(wr, transposed_note)
// 				} else {
// 					wr.SetDelta(e.Position.DeltaTicks)
// 					writer.NoteOn(wr, transposed_note, e.Velocity)
// 				}
// 			}
// 		}
// 		return nil
// 	}, smfwriter.TimeFormat(rd.Header().TimeFormat))

// 	if err != nil {
// 		log.Fatalf("could not write SMF file %v", wf)
// 	}
// }