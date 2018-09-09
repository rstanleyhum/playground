extern crate html5ever;

#[macro_use]
extern crate lazy_static;

use std::cell::RefCell;
use std::collections::HashMap;
use std::collections::HashSet;
use std::default::Default;
use std::env;
use std::fs;
use std::fs::DirEntry;
use std::fs::File;
use std::io::prelude::*;
use std::io::Error;
use std::path::Path;
use std::path::PathBuf;
use std::result::Result;
use std::string::String;
use std::vec::Vec;

use html5ever::parse_document;
use html5ever::rcdom::{Handle, NodeData, RcDom};
use html5ever::tendril::TendrilSink;

lazy_static! {
    static ref ATTRS_STRINGS: HashSet<&'static str> = {
        [
            "class",
            "id",
            "cols",
            "rows",
            "name",
            "for",
            "selected",
            "ng-hide",
            "ng-show",
            "ng-switch",
            "size",
            "maxlength",
            "msd-elastic",
        ].iter()
            .cloned()
            .collect()
    };
}

struct ElementNode {
    indent: usize,
    name: String,
    attrs: HashMap<String, String>,
}

fn walk(indent: usize, handle: Handle) -> Vec<ElementNode> {
    let mut data_elements: Vec<ElementNode> = Vec::new();

    let node = handle;

    match node.data {
        NodeData::Document => (),

        NodeData::Doctype { .. } => (),

        NodeData::Text { ref contents } => {
            let trimmed = contents.clone().into_inner().to_string();
            let trimmed = trimmed.trim();
            if !trimmed.is_empty() {
                let mut attrs: HashMap<String, String> = HashMap::new();
                attrs.insert("value".to_string(), trimmed.to_string());

                let e = ElementNode {
                    indent: indent,
                    name: "TEXT".to_string(),
                    attrs: attrs,
                };
                data_elements.push(e);
            }
        }

        NodeData::Comment { .. } => (),

        NodeData::Element {
            ref name,
            ref attrs,
            ..
        } => {
            let attrs_str = get_attrs(attrs);
            let e = ElementNode {
                indent: indent,
                name: name.local.to_string(),
                attrs: attrs_str,
            };
            data_elements.push(e);
        }

        NodeData::ProcessingInstruction { .. } => unreachable!(),
    }

    let es: Vec<_> = node
        .children
        .borrow()
        .iter()
        .flat_map(|c| walk(indent + 4, c.clone()))
        .collect();

    data_elements.extend(es);

    data_elements
}

pub fn get_attrs(attrs: &RefCell<Vec<html5ever::Attribute>>) -> HashMap<String, String> {
    let data: HashMap<String, String> = attrs
        .borrow()
        .iter()
        .filter(|a| !ATTRS_STRINGS.contains(&*a.name.local.to_string()))
        .map(|a| (a.name.local.to_string(), a.value.to_string()))
        .collect();

    data
}

pub fn escape_default(s: &str) -> String {
    s.chars().flat_map(|c| c.escape_default()).collect()
}

fn get_attr_str(attrs: &HashMap<String, String>) -> String {
    let data: Vec<String> = attrs
        .into_iter()
        .map(|(k, v)| [k.to_string(), v.to_string()].join("="))
        .collect();
    data.join(" ")
}

fn write_to_file(filename: &PathBuf, elements: &Vec<ElementNode>) {
    let new_filename = match filename.to_str() {
        Some(f) => vec![f, ".txt"].join(""),
        None => panic!("No new filename!"),
    };

    let mut file = match File::create(new_filename) {
        Err(e) => panic!(e),
        Ok(f) => f,
    };

    elements.into_iter().for_each(|x| {
        let mut elem_name = "";
        if x.name != "html" && x.name != "head" && x.name != "body" && x.name != "div" {
            elem_name = &x.name;
        }
        writeln!(file, "{e} {a}", e = elem_name, a = get_attr_str(&x.attrs)).expect("problem")
    });
}

fn process_file(filename: &PathBuf) -> Vec<ElementNode> {
    let dom = parse_document(RcDom::default(), Default::default())
        .from_utf8()
        .from_file(Path::new(&filename))
        .unwrap();

    walk(0, dom.document)
}

fn check_filename(p: &Result<DirEntry, Error>) -> bool {
    match p {
        Ok(d) => match d.file_name().to_str() {
            Some(f) => f.to_string().contains("edit"),
            None => false,
        },
        Err(_) => false,
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let directory = &args[1];

    let paths = fs::read_dir(directory).unwrap();

    paths.filter(|p| check_filename(p)).for_each(|path| {
        let filename = path.unwrap().path();

        let elements = process_file(&filename);

        write_to_file(&filename, &elements);
    });
}
