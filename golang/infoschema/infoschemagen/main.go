package main

import (
	"fmt"
	"log"
	"os"

	"_snippets/golang/infoschema"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	databaseName := "works"
	hostName := "<server ip>"
	filepath := "<results filename>"

	dbURL := infoschema.CreateDBUrl(databaseName, hostName)

	db, err := infoschema.GetDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	tableNames := []string{
		"CHECK_CONSTRAINTS",
		"COLUMN_DOMAIN_USAGE",
		"COLUMN_PRIVILEGES",
		"CONSTRAINT_COLUMN_USAGE",
		"CONSTRAINT_TABLE_USAGE",
		"DOMAIN_CONSTRAINTS",
		"DOMAINS",
		"KEY_COLUMN_USAGE",
		"REFERENTIAL_CONSTRAINTS",
		"ROUTINES",
		"ROUTINE_COLUMNS",
		"SCHEMATA",
		"TABLE_CONSTRAINTS",
		"TABLE_PRIVILEGES",
		"TABLES",
		"VIEW_COLUMN_USAGE",
		"VIEW_TABLE_USAGE",
		"PARAMTERS",
		"VIEWS",
		"COLUMNS",
	}
	for _, v := range tableNames {
		fmt.Fprintf(file, "Table: %s\n", v)
		err = infoschema.GetColumnsInTable(db, file, v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(file, "\n")
	}

}
