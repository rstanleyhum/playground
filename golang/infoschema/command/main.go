package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"_snippets/golang/infoschema"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	query := url.Values{}
	query.Add("database", "works")

	u := &url.URL{
		Scheme:   "sqlserver",
		Host:     "<server ip>",
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("mssql", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Create("<results filename>")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	log.SetOutput(file)

	var v infoschema.CheckConstraint

	rows, err := db.Query("select CONSTRAINT_CATALOG, CONSTRAINT_SCHEMA, CONSTRAINT_NAME, CHECK_CLAUSE from INFORMATION_SCHEMA.CHECK_CONSTRAINTS;")
	//rows, err := db.Query("select TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE from INFORMATION_SCHEMA.TABLES;")
	//rows, err := db.Query("select TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, ORDINAL_POSITION, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, CHARACTER_OCTET_LENGTH, NUMERIC_PRECISION, NUMERIC_PRECISION_RADIX, NUMERIC_SCALE, DATETIME_PRECISION, CHARACTER_SET_CATALOG, CHARACTER_SET_SCHEMA, CHARACTER_SET_NAME, COLLATION_CATALOG, COLLATION_SCHEMA, COLLATION_NAME, DOMAIN_CATALOG, DOMAIN_SCHEMA, DOMAIN_NAME from INFORMATION_SCHEMA.COLUMNS;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		err := rows.Scan(&v.Catalog, &v.Schema, &v.Name, &v.CheckClause)
		//err := rows.Scan(&v.Catalog, &v.Schema, &v.Name, &v.Type)
		//err := rows.Scan(&v.TableCatalog, &v.TableSchema, &v.TableName, &v.Name, &v.OrdinalPosition, &v.Default, &v.IsNullable, &v.DataType, &v.CharMaxLen, &v.CharOctetLen, &v.NumPrecision, &v.NumPrecisionRadix, &v.NumScale, &v.DateTimePrec, &v.CharSetCatalog, &v.CharSetSchema, &v.CharSetName, &v.CollationCatalog, &v.CollationSchema, &v.CollationName, &v.DomainCatalog, &v.DomainSchema, &v.DomainName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(file, "%v\n\n", v)
		count = count + 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	fmt.Fprintf(file, "\ncount = %v\n", count)
}
