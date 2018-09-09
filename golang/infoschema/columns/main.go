package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	query := url.Values{}
	query.Add("database", "works")

	u := &url.URL{
		Scheme:   "sqlserver",
		Host:     "",
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("mssql", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Create("<columns filename>")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	log.SetOutput(file)
	rows, err := db.Query("select * from INFORMATION_SCHEMA.KEY_COLUMN_USAGE;")
	//rows, err := db.Query("select CONSTRAINT_CATALOG, CONSTRAINT_SCHEMA, CONSTRAINT_NAME, CHECK_CLAUSE from INFORMATION_SCHEMA.CHECK_CONSTRAINTS;")
	//rows, err := db.Query("select TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE from INFORMATION_SCHEMA.TABLES;")
	//rows, err := db.Query("select TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, ORDINAL_POSITION, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, CHARACTER_OCTET_LENGTH, NUMERIC_PRECISION, NUMERIC_PRECISION_RADIX, NUMERIC_SCALE, DATETIME_PRECISION, CHARACTER_SET_CATALOG, CHARACTER_SET_SCHEMA, CHARACTER_SET_NAME, COLLATION_CATALOG, COLLATION_SCHEMA, COLLATION_NAME, DOMAIN_CATALOG, DOMAIN_SCHEMA, DOMAIN_NAME from INFORMATION_SCHEMA.COLUMNS;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(file, "columns = %v\n", cols)

	ctypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range ctypes {
		fmt.Fprintf(file, "&ctypes = %v\n", v)
	}

	vals := make([]interface{}, len(cols))
	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}

	count := 0
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			log.Fatal(err)
		}
		count = count + 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	fmt.Fprintf(file, "\ncount = %v\n", count)
}
