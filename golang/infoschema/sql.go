package infoschema

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func CreateDBUrl(dbname string, host string) string {
	query := url.Values{}
	query.Add("database", dbname)

	u := &url.URL{
		Scheme:   "sqlserver",
		Host:     host,
		RawQuery: query.Encode(),
	}

	return u.String()
}

func GetDB(dbURL string) (*sql.DB, error) {

	db, err := sql.Open("mssql", dbURL)
	if err != nil {
		return db, err
	}

	return db, nil
}

func getRowsForTable(db *sql.DB, tableName string, limit int) (*sql.Rows, error) {
	var limitClause strings.Builder
	if limit > 0 {
		fmt.Fprintf(&limitClause, "Top %d", limit)
	}

	const queryBaseFmt = "select %s * from INFORMATION_SCHEMA.%s;"
	var query strings.Builder
	fmt.Fprintf(&query, queryBaseFmt, limitClause.String(), tableName)

	rows, err := db.Query(query.String())
	if err != nil {
		return rows, err
	}
	return rows, nil
}

func GetColumnsInTable(db *sql.DB, f *os.File, tableName string) error {

	rows, err := getRowsForTable(db, tableName, 1)
	if err != nil {
		return err
	}
	defer rows.Close()

	ctypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	for _, v := range ctypes {
		fmt.Fprintf(f, "%s (%s)\n", v.Name(), v.DatabaseTypeName())
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	rows.Close()

	return nil
}
