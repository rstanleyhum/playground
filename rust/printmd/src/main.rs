use std::fs::File;
use std::path::Path;
use std::io::BufReader;
use std::io::BufRead;
use std::io::BufWriter;
use std::io::Write;

fn main() {
    let inpath = Path::new("src/main.rs");
    let f = File::open(&inpath).unwrap();
    let file = BufReader::new(&f);
    
    let outpath = Path::new("main.md");
    let of = File::create(&outpath).unwrap();

    let mut writer = BufWriter::new(&of);
    
    let s = "---\ntitle: Listings\nauthor: Stanley Hum\nlanguage: en-US\n---\n";

    writer.write(s.as_bytes()).unwrap();

    writer.write("## main.rs\n".as_bytes()).unwrap();
    
    writer.write("```\n".as_bytes()).unwrap();

    for (_, line) in file.lines().enumerate() {
        let l = line.unwrap();
        writer.write(l.as_bytes()).unwrap();
        writer.write(b"\n").unwrap(); 
    }

    writer.write("```\n".as_bytes()).unwrap();

    writer.write("## second.txt\n".as_bytes()).unwrap();

    writer.write("The quick brown fox jumps over the lazy dog.\n".as_bytes()).unwrap();

}
