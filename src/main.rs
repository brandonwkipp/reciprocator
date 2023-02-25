use clap::builder::ArgAction;
use clap::Parser;
use rimd::{note_num_to_name};

mod reciprocator;
use crate::reciprocator::file;

/// Search for a pattern in a file and display the lines that contain it.
#[derive(Debug, Parser)]
struct Args {
    #[arg(long, action = ArgAction::SetTrue)]
   debug: bool,

   #[arg(long, short)]
   file: String,

    #[arg(long, action = ArgAction::SetTrue)]
   invert: bool,

   #[arg(long, short)]
   tonal_center: u8,
}

fn main() {
    let args = Args::parse();

    if args.debug {
        println!("Tonal center: {}", note_num_to_name(args.tonal_center as u32));
        file::debug_smf(&args.file);
        return;
    }

    file::write_file(&args.file, args.tonal_center, file::construct_output_filename(&args.file, args.invert), args.invert);
}
