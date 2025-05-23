package readdb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type Recipes struct {
	Cake []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			IngredientName  string `xml:"itemname" json:"ingredient_name"`
			IngredientCount string `xml:"itemcount" json:"ingredient_count"`
			IngredientUnit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

func ParsFile(file []byte, format string) (db Recipes, err error) {
	if format == "json" {
		err = json.Unmarshal(file, &db)
	} else if format == "xml" {
		err = xml.Unmarshal(file, &db)
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

func ReadFile(file_name string) ([]byte, error) {
	file, err := os.ReadFile(file_name)
	if err != nil {

		return nil, fmt.Errorf("open file error: %w", err)
	}

	return file, nil
}

func PrintRecipes(res Recipes, format string) {
	if format == "xml" {
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	}
	if format == "json" {
		b, err := xml.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	}
}

func CheckFormatFile(file_name string) (string, error) {
	if file_name == "" {

		return "", fmt.Errorf("no file found")
	}
	if filepath.Ext(file_name) == ".json" {

		return "json", nil
	}
	if filepath.Ext(file_name) == ".xml" {

		return "xml", nil
	}
	return "", fmt.Errorf("only json or xml")
}
