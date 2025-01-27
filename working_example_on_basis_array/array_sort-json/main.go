package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
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

func sortPrimitives() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	var names []string
	for _, champ := range champions {
		names = append(names, champ.Name)
	}

	names = names[:10]

	log.Printf(
		"The first %d names are %v, sorted = %t\n\n",
		len(names),
		names,
		sort.StringsAreSorted(names),
	)

	sort.Strings(names)
	log.Printf(
		"The sorted names (ascending) are %v, sorted = %t\n\n",
		names,
		sort.StringsAreSorted(names),
	)

	var gold []int
	for _, champ := range champions {
		gold = append(gold, champ.Cost)
	}

	gold = gold[:10]

	log.Printf(
		"The first %d gold costs are %v, sorted = %t\n\n",
		len(gold),
		gold,
		sort.IntsAreSorted(gold),
	)

	sort.Ints(gold)
	log.Printf(
		"The sorted gold are %v, sorted = %t\n\n",
		gold,
		sort.IntsAreSorted(gold),
	)
}

func sortSlices() {
	champions, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	champs := champions[:10]

	log.Printf("The first %d champs are %v\n\n", len(champs), champs)

	// Sort champions in ascending alphabetical order by Name
	sort.Slice(champs, func(i, j int) bool {
		return champs[i].Name < champs[j].Name
	})
	log.Printf("The sorted champions (ascending by Name) are %v\n\n", champs)
}

func main() {
	sortSlices()
}

// 2024/10/02 22:09:53 The first 10 champs are [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4}]

// 2024/10/02 22:09:53 The sorted champions (ascending by Name) are [{Aatrox [Blademaster] [Demon] 3} {Akali [Assassin] [Ninja] 4} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2}]
