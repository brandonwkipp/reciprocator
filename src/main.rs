mod reciprocator;
use crate::reciprocator::file;

fn main() {
    // println!("{}", file::construct_output_filename("arabesque.mid", false));
    // file::debug_smf("arabesque.mid");
    file::write_file("arabesque.mid", 100, "", false);
}
