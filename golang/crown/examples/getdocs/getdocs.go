package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"_snippets/golang/crown/table"

	"github.com/jmoiron/sqlx"

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

	db, err := sqlx.Open("mssql", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Create("<text file>")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	log.SetOutput(file)

	// Get documents
	docs := []table.DboDocument{}
	myquery := `SELECT TOP (1000)
		*
	FROM <database>.dbo.Document WITH (NOLOCK)
	WHERE PatientID=<patient id>;
	`
	err = db.Select(&docs, myquery)
	if err != nil {
		log.Fatal(err)
	}
	count := len(docs)
	for _, v := range docs {
		fmt.Fprintf(file, "%v\n", v)
	}

	fmt.Fprintf(file, "\ncount = %v\n", count)

	// Get instance output map
	// Not all documents have an output map
	outputmaps := []table.DboNoteFormNAWInstanceOutputMap{}
	myquery = `SELECT TOP (1000)
		*
	FROM <database>.dbo.NoteForm_NAW_Instance_Output_Map WITH (NOLOCK)
	WHERE OriginatingDocumentID=<document id>;
	`
	err = db.Select(&outputmaps, myquery)
	if err != nil {
		log.Fatal(err)
	}
	count = len(outputmaps)
	for _, v := range outputmaps {
		fmt.Fprintf(file, "%v\n", v)
	}
	fmt.Fprintf(file, "\ncount = %v\n", count)

	// Get output template
	outputtpls := []table.DboNoteFormOutputTemplate{}
	myquery = `SELECT TOP (1000)
		*
	FROM <database>.dbo.NoteForm_Output_Template WITH (NOLOCK)
	WHERE TemplateID=<template id>;
	`
	err = db.Select(&outputtpls, myquery)
	if err != nil {
		log.Fatal(err)
	}
	count = len(outputtpls)
	for _, v := range outputtpls {
		fmt.Fprintf(file, "%v\n", v)
	}
	fmt.Fprintf(file, "\ncount = %v\n", count)

	// Get instance header
	insts := []table.DboNoteFormNAWInstanceHeader{}
	myquery = `SELECT TOP (1000)
		*
	FROM <database>.dbo.NoteForm_NAW_Instance_Header WITH (NOLOCK)
	WHERE HeaderID=<headerid>;
	`
	err = db.Select(&insts, myquery)
	if err != nil {
		log.Fatal(err)
	}
	count = len(insts)
	for _, v := range insts {
		fmt.Fprintf(file, "%v\n", v)
	}
	fmt.Fprintf(file, "\ncount = %v\n", count)

	// Get input tempalte from instance header
	inputtps := []table.DboNoteFormInputTemplate{}
	myquery = `SELECT TOP (1000)
		*
	FROM <database>.dbo.NoteForm_Input_Template WITH (NOLOCK)
	WHERE TemplateID=<template_id>;
	`
	err = db.Select(&inputtps, myquery)
	if err != nil {
		log.Fatal(err)
	}
	count = len(inputtps)
	for _, v := range inputtps {
		fmt.Fprintf(file, "%v\n", v)
	}
	fmt.Fprintf(file, "\ncount = %v\n", count)

}
