mod reciprocator;
use crate::reciprocator::file;

fn main() {
    let filename = "arabesque.mid";
    let invert = false;

    file::write_file(filename, 100, file::construct_output_filename(filename, invert), invert);
}
