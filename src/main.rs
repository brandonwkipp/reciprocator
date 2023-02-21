use clap::builder::ArgAction;
use clap::Parser;

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
}

fn main() {
    let args = Args::parse();

    if args.debug {
        file::debug_smf(&args.file);
        return;
    }

    file::write_file(&args.file, 60, file::construct_output_filename(&args.file, args.invert), args.invert);
}
