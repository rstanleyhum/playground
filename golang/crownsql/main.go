package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	username := "<username>"
	password := "<password>"
	hostname := "<server ip>"
	u := &url.URL{
		Scheme: "sqlserver",
		User: url.UserPassword(username, password),
		Host: hostname,
	}

	db, err := sql.Open("mssql", u.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Opened db connection\n")

	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		fmt.Println("Error in Ping")
		log.Println(err)
	}

	rows, err := db.Query("SELECT name FROM works.sys.tables ORDER BY name")
	if err != nil {
		fmt.Println("Error in Select")
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("Error in scan")
			log.Println(err)
			continue
		}

		fmt.Println(name)
	}

		

	fmt.Printf("Before closing\n")
}

