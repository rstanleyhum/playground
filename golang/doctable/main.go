package main

import (
	"bytes"
	"compress/flate"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "<sqlite3 databaase>")
	rows, err := db.Query("SELECT EditableChunkCompressed FROM Document LIMIT 3;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var input []byte
	for rows.Next() {
		err := rows.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(len(input))
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("hello.bin", input, 0644)
	if err != nil {
		log.Fatal(err)
	}

	b := bytes.NewReader(input)
	r := flate.NewReader(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if _, err := io.ReadFull(r, b[:1]); err != nil {
		if err == io.EOF {
			break
		}
		log.Fatal(err)
	}
	 n := int(b[0])
	 if _, err := io.ReadFull(r, b[:n]) err != nil {
		 log.Fatal(err)
	 }

	 fmt.Printf("Received")
	//r.Close()
}
