use clap::builder::ArgAction;
use clap::Parser;

mod reciprocator;
use crate::reciprocator::file;

/// Search for a pattern in a file and display the lines that contain it.
#[derive(Parser)]
struct Args {
    #[arg(long, action = ArgAction::SetTrue)]
   invert: bool,
}


fn main() {
    let args = Args::parse();

    let filename = "arabesque.mid";

    file::write_file(filename, 100, file::construct_output_filename(filename, args.invert), args.invert);
}
