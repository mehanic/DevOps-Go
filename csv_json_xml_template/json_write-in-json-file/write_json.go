package main

import (
	"encoding/json"
	"os"
)

func generateDictionary(monsterName string, title string, price int, scariness int) map[string]interface{} {
	return map[string]interface{}{
		"monster_name": monsterName,
		"title":        title,
		"price":        price,
		"scariness":    scariness,
	}
}

func writeToJSON(dictionaryData []map[string]interface{}) {
	file, err := os.Create("monsters.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(dictionaryData)
	if err != nil {
		panic(err)
	}
}

func main() {
	monsterOne := generateDictionary("Filooooo", "Baker by Day, Techie by Night", 29, 3)
	monsterTwo := generateDictionary("Timber", "Database Expert", 19, 2)
	monsterThree := generateDictionary("Blade", "Monster APPetite", 29, 5)
	writeToJSON([]map[string]interface{}{monsterOne, monsterTwo, monsterThree})
}
