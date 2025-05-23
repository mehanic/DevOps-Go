package main

import (
	"day01/readdb"
	"flag"
	"fmt"
)

type DBReader interface {
	ReadDb() (db readdb.Recipes, err error)
}

type XmlDb struct {
	filename string
}

func (b XmlDb) ReadDb() (db readdb.Recipes, err error) {
	file, err := readdb.ReadFile(b.filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	db, err = readdb.ParsFile(file, "xml")

	return
}

type JsonDb struct {
	filename string
}

func (b JsonDb) ReadDb() (db readdb.Recipes, err error) {
	file, err := readdb.ReadFile(b.filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	db, err = readdb.ParsFile(file, "json")

	return
}

func main() {

	filename := flag.String("f", "", "json or xml file")
	flag.Parse()

	format, db, err := OpenDB(*filename)

	if err != nil {
		fmt.Println(err.Error())

		return
	}
	readdb.PrintRecipes(db, format)
}

func OpenDB(filename string) (format string, db readdb.Recipes, err error) {
	format, err = readdb.CheckFormatFile(filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if format == "json" {
		json := JsonDb{filename}
		db, err = ThrouInterface(json)

	} else if format == "xml" {
		xml := XmlDb{filename}
		db, err = ThrouInterface(xml)

	} else {
		err = fmt.Errorf("unknown format: %w", err)
		return
	}

	if err != nil {
		err = fmt.Errorf("pars file error: %w", err)
		return
	}

	return
}

func ThrouInterface(s DBReader) (db readdb.Recipes, err error) {
	db, err = s.ReadDb()
	if err != nil {
		err = fmt.Errorf("open file error: %w", err)

		return
	}

	return s.ReadDb()
}
