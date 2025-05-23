package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}

func loadChampions() ([]champion, error) {
	file, err := os.Open("tft_champions.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var champions []champion
	for {
		if err := json.NewDecoder(file).Decode(&champions); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return champions, nil
}

func main() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	for index, champ := range champions {
		log.Printf("%s is at index %d\n", champ.Name, index)
	}

	log.Println("---------------------")

	for _, champ := range champions {
		log.Printf("Champion is %s\n", champ.Name)
	}

	log.Println("---------------------")

	for index := range champions {
		log.Printf("There is a champion at index %d\n", index)
	}

	log.Println("---------------------")

	for i := 0; i < len(champions); i++ {
		log.Printf("%s is at index %d\n", champions[i].Name, i)
	}

	log.Println("---------------------")

	for i := len(champions) - 1; i >= 0; i-- {
		log.Printf("%s is at index %d\n", champions[i].Name, i)
	}
}

// 2024/10/02 21:57:38 Akali is at index 0
// 2024/10/02 21:57:38 Evelynn is at index 1
// 2024/10/02 21:57:38 Katarina is at index 2
// 2024/10/02 21:57:38 Kha'Zix is at index 3
// 2024/10/02 21:57:38 Pyke is at index 4
// 2024/10/02 21:57:38 Rengar is at index 5
// 2024/10/02 21:57:38 Zed is at index 6
// 2024/10/02 21:57:38 Aatrox is at index 7
// 2024/10/02 21:57:38 Camille is at index 8
