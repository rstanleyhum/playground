package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	_ "github.com/denisenkom/go-mssqldb"
)

func convertToColumnName(n string) string {
	sb := strings.Builder{}

	first := true
	count := 0
	for _, c := range n {
		upper := unicode.IsUpper(c)

		if upper && (count == 1 || count == 2 || count == 3) {
			count = count + 1
		}

		if upper && (count == 0 || count == 4) {
			if !first {
				sb.WriteRune('_')
			}
			count = 1
		}

		if !upper {
			count = 0
		}
		sb.WriteRune(c)
		first = false
	}
	return sb.String()
}

func main() {

	// pfile, err := ioutil.ReadFile("./passwd.txt")
	// if err != nil {
	// 	log.Fatalf("Must have a password file (%v)\n", err)
	// }
	// password := string(pfile)

	// ufile, err := ioutil.ReadFile("./username.txt")
	// if err != nil {
	// 	log.Fatalf("Must have a username file (%v)\n", err)
	// }
	// username := string(ufile)

	// query := url.Values{}
	// query.Add("database", "works")

	// u := &url.URL{
	// 	Scheme:   "sqlserver",
	// 	User:     url.UserPassword(username, password),
	// 	Host:     "<server ip>",
	// 	RawQuery: query.Encode(),
	// }

	// db, err := sql.Open("mssql", u.String())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	file, err := os.Create("<results filename>")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	log.SetOutput(file)

	tableNames := []string{
		"CHECK_CONSTRAINTS",
		"COLUMN_DOMAIN_USAGE",
		"COLUMN_PRIVILEGES",
		"COLUMNS",
		"CONSTRAINT_COLUMN_USAGE",
		"CONSTRAINT_TABLE_USAGE",
		"DOMAIN_CONSTRAINTS",
		"DOMAINS",
		"KEY_COLUMN_USAGE",
		"PARAMETERS",
		"REFERENTIAL_CONSTRAINTS",
		"ROUTINE_COLUMNS",
		"ROUTINES",
		"SCHEMATA",
		"TABLE_CONSTRAINTS",
		"TABLE_PRIVILEGES",
		"TABLES",
		"VIEW_COLUMN_USAGE",
		"VIEW_TABLE_USAGE",
		"VIEWS",
	}

	for _, v := range tableNames {
		fmt.Fprintf(file, "Table: %s\n", v)
		//querystring := fmt.Sprintf("select COUNT(*) from works.INFORMATION_SCHEMA.%s;", v)

		fmt.Printf("%s\n", convertToColumnName("SQLName"))
		fmt.Printf("%v\n", convertToColumnName("ThisIsConstraintName"))

		// rows, err := db.Query(querystring)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer rows.Close()

		// var l int
		// for rows.Next() {
		// 	err := rows.Scan(&l)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Fprintf(file, "%v\n", l)
		// }

		// err = rows.Err()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// rows.Close()

	}
}
