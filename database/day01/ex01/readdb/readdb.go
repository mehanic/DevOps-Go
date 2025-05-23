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

func ParsFile(file_name string) (Recipes, error, string) {
	var base Recipes
	var format string
	format, err := checkFormatFile(file_name)
	if err != nil {
		fmt.Println(err.Error())

		return base, err, format
	}

	file, err := ReadFile(file_name)
	if err != nil {

		return base, fmt.Errorf("open file error: %w", err), format
	}
	if format == "json" {
		err = json.Unmarshal(file, &base)
	} else if format == "xml" {
		err = xml.Unmarshal(file, &base)
	} else {

		return base, fmt.Errorf("unknown format: %w", err), format
	}
	if err != nil {

		return base, fmt.Errorf("pars file error: %w", err), format
	}

	return base, nil, format
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

func checkFormatFile(file_name string) (string, error) {
	if file_name == "" {

		return "", fmt.Errorf("ошибка: отсутствует файл")
	}
	if filepath.Ext(file_name) == ".json" {

		return "json", nil
	}
	if filepath.Ext(file_name) == ".xml" {

		return "xml", nil
	}
	return "", fmt.Errorf("ошибка: файл должен быть json или xml")
}
