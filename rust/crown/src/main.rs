extern crate odbc;
extern crate env_logger;

use odbc::*;
use std::io;

fn main() {

    println!("Starting logger ...");    
    env_logger::init();

    match connect() {
        Ok(()) => println!("\nSuccess"),
        Err(diag) => println!("\nError {}", diag),
    }
}

 fn connect() -> std::result::Result<(), DiagnosticRecord> {
    let env = create_environment_v3().map_err(|e| e.unwrap())?;
    let conn_str = String::from("Driver={ODBC Driver 13 for SQL Server};Server=tcp:<server ip>;Database=<database>;Trusted_Connection=yes;");
    // let mut buffer = String::new();
    // println!("Please enter connection string: ");
    // io::stdin().read_line(&mut buffer).unwrap();

    let conn = env.connect_with_connection_string(&conn_str)?;
    execute_statement(&conn)
 }

 fn execute_statement<'env>(conn: &Connection<'env>) -> Result<()> {
    let stmt = Statement::with_parent(conn)?;

    let mut sql_text = String::new();
    println!("Please enter SQL statement string: ");
    io::stdin().read_line(&mut sql_text).unwrap();

    match stmt.exec_direct(&sql_text)? {
        Data(mut stmt) =>{
            let cols = stmt.num_result_cols()?;
            while let Some(mut cursor) = stmt.fetch()? {
                for i in 1..(cols + 1) {
                    match cursor.get_data::<&str>(i as u16)? {
                        Some(val) => print!(" {}", val),
                        None => print!(" NULL"),
                    }
                }
                print!("\n");
            }
            print!("\n");
        },
        NoData(_) => println!("Query executed, no data returned"),
    }

    Ok(())
 }