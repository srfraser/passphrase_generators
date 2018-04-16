extern crate rand;

use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;
use std::path::Path;
use rand::{Rng, thread_rng};

fn read_dictionary<P>(filename: P) -> Vec<String>
where
    P: AsRef<Path>,
{
    let f = File::open(filename).expect("File not found");
    let buf = BufReader::new(f);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}

fn main() {
    let filename = "/usr/share/dict/words";
    let length = 4;
    let mut wordlist = read_dictionary(filename);

    thread_rng().shuffle(&mut wordlist);
    println!("{}", wordlist.iter().take(length).map(|s| s.clone()).collect::<Vec<_>>().join(" "));

}
