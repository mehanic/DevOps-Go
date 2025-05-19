package main

import (
	"encoding/json"
	"fmt"
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

func (c champion) hasAllClasses(classes ...string) bool {
	classCount := len(classes)
	foundCount := 0
	for _, class := range classes {
		if found := c.hasClass(class); found {
			foundCount++
		}
	}

	return foundCount == classCount
}

func (c champion) hasSomeClasses(classes ...string) bool {
	for _, class := range classes {
		if found := c.hasClass(class); found {
			return true
		}
	}

	return false
}

func (c champion) hasClass(class string) bool {
	for _, champClass := range c.Classes {
		if champClass == class {
			return true
		}
	}

	return false
}

func (c champion) hasAllOrigins(origins ...string) bool {
	originCount := len(origins)
	foundCount := 0
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			foundCount++
		}
	}

	return foundCount == originCount
}

func (c champion) hasSomeOrigins(origins ...string) bool {
	for _, origin := range origins {
		if found := c.hasOrigin(origin); found {
			return true
		}
	}

	return false
}

func (c champion) hasOrigin(origin string) bool {
	for _, champOrigin := range c.Origins {
		if champOrigin == origin {
			return true
		}
	}

	return false
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
	all, err := loadChampions()
	if err != nil {
		log.Fatalf("An error occurred loading/parsing champions, err=%v", err)
	}

	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	first := all[0]
	fmt.Printf("The first champion is %v\n\n", first)

	some := all[0:20]
	some = all[:20]
	some = all[10:]
	some = all[:]
	fmt.Printf("The first %d champions are %v\n\n", len(some), some)

	// Remove an element from the slice by index
	most := append(all[0:10], all[11:]...)
	fmt.Printf(
		"There are %d/%d champions (most): %v\n\n",
		len(most),
		cap(most),
		most,
	)
	fmt.Printf("There are %d/%d champions (all): %v\n\n", len(all), cap(all), all)

	// Remove elements by altering the slice in place
	index := 0
	for _, champ := range all {
		if champ.hasClass("Sorcerer") {
			continue
		}

		all[index] = champ
		index++
	}

	all = all[:index]

	fmt.Printf(
		"There are %d champions after removing Sorcerers, %v\n\n",
		len(all),
		all,
	)
}

// There are 56/74 champions (all): [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Fiora [Blademaster] [Noble] 1} {Gangplank [Blademaster Gunslinger] [Pirate] 3} {Shen [Blademaster] [Ninja] 2} {Yasuo [Blademaster] [Exile] 5} {Blitzcrank [Brawler] [Robot] 2} {Cho'gath [Brawler] [Void] 4} {Rek'Sai [Brawler] [Void] 2} {Vi [Brawler] [Hextech] 3} {Volibear [Brawler] [Glacial] 3} {Warwick [Brawler] [Wild] 1} {Anivia [Elementalist] [Glacial] 5} {Brand [Elementalist] [Demon] 4} {Kennen [Elementalist] [Ninja Yordle] 3} {Lissandra [Elementalist] [Glacial] 2} {Braum [Guardian] [Glacial] 2} {Leona [Guardian] [Noble] 4} {Pantheon [Guardian] [Dragon] 5} {Graves [Gunslinger] [Pirate] 1} {Lucian [Gunslinger] [Noble] 2} {Jinx [Gunslinger] [Hextech] 4} {Miss Fortune [Gunslinger] [Pirate] 5} {Tristana [Gunslinger] [Yordle] 1} {Darius [Knight] [Imperial] 1} {Garen [Knight] [Noble] 1} {Kayle [Knight] [Noble] 5} {Mordekaiser [Knight] [Phantom] 1} {Poppy [Knight] [Yordle] 3} {Sejuani [Knight] [Glacial] 4} {Ashe [Ranger] [Glacial] 3} {Kindred [Ranger] [Phantom] 4} {Varus [Ranger] [Demon] 2} {Vayne [Ranger] [Noble] 1} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4} {Jayce [Shapeshifter] [Hextech] 2} {Nidalee [Shapeshifter] [Wild] 1} {Shyvanna [Shapeshifter] [Dragon] 3} {Swain [Shapeshifter] [Demon Imperial] 5} {Ahri [Sorcerer] [Wild] 2} {Aurelion Sol [Sorcerer] [Dragon] 4} {Karthus [Sorcerer] [Phantom] 5} {Kassadin [Sorcerer] [Void] 1} {LuLu [Sorcerer] [Yordle] 2} {Morgana [Sorcerer] [Demon] 3} {Twisted Fate [Sorcerer] [Pirate] 2} {Veigar [Sorcerer] [Yordle] 3}]

// The first champion is {Akali [Assassin] [Ninja] 4}

// The first 56 champions are [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Fiora [Blademaster] [Noble] 1} {Gangplank [Blademaster Gunslinger] [Pirate] 3} {Shen [Blademaster] [Ninja] 2} {Yasuo [Blademaster] [Exile] 5} {Blitzcrank [Brawler] [Robot] 2} {Cho'gath [Brawler] [Void] 4} {Rek'Sai [Brawler] [Void] 2} {Vi [Brawler] [Hextech] 3} {Volibear [Brawler] [Glacial] 3} {Warwick [Brawler] [Wild] 1} {Anivia [Elementalist] [Glacial] 5} {Brand [Elementalist] [Demon] 4} {Kennen [Elementalist] [Ninja Yordle] 3} {Lissandra [Elementalist] [Glacial] 2} {Braum [Guardian] [Glacial] 2} {Leona [Guardian] [Noble] 4} {Pantheon [Guardian] [Dragon] 5} {Graves [Gunslinger] [Pirate] 1} {Lucian [Gunslinger] [Noble] 2} {Jinx [Gunslinger] [Hextech] 4} {Miss Fortune [Gunslinger] [Pirate] 5} {Tristana [Gunslinger] [Yordle] 1} {Darius [Knight] [Imperial] 1} {Garen [Knight] [Noble] 1} {Kayle [Knight] [Noble] 5} {Mordekaiser [Knight] [Phantom] 1} {Poppy [Knight] [Yordle] 3} {Sejuani [Knight] [Glacial] 4} {Ashe [Ranger] [Glacial] 3} {Kindred [Ranger] [Phantom] 4} {Varus [Ranger] [Demon] 2} {Vayne [Ranger] [Noble] 1} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4} {Jayce [Shapeshifter] [Hextech] 2} {Nidalee [Shapeshifter] [Wild] 1} {Shyvanna [Shapeshifter] [Dragon] 3} {Swain [Shapeshifter] [Demon Imperial] 5} {Ahri [Sorcerer] [Wild] 2} {Aurelion Sol [Sorcerer] [Dragon] 4} {Karthus [Sorcerer] [Phantom] 5} {Kassadin [Sorcerer] [Void] 1} {LuLu [Sorcerer] [Yordle] 2} {Morgana [Sorcerer] [Demon] 3} {Twisted Fate [Sorcerer] [Pirate] 2} {Veigar [Sorcerer] [Yordle] 3}]

// There are 55/74 champions (most): [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Gangplank [Blademaster Gunslinger] [Pirate] 3} {Shen [Blademaster] [Ninja] 2} {Yasuo [Blademaster] [Exile] 5} {Blitzcrank [Brawler] [Robot] 2} {Cho'gath [Brawler] [Void] 4} {Rek'Sai [Brawler] [Void] 2} {Vi [Brawler] [Hextech] 3} {Volibear [Brawler] [Glacial] 3} {Warwick [Brawler] [Wild] 1} {Anivia [Elementalist] [Glacial] 5} {Brand [Elementalist] [Demon] 4} {Kennen [Elementalist] [Ninja Yordle] 3} {Lissandra [Elementalist] [Glacial] 2} {Braum [Guardian] [Glacial] 2} {Leona [Guardian] [Noble] 4} {Pantheon [Guardian] [Dragon] 5} {Graves [Gunslinger] [Pirate] 1} {Lucian [Gunslinger] [Noble] 2} {Jinx [Gunslinger] [Hextech] 4} {Miss Fortune [Gunslinger] [Pirate] 5} {Tristana [Gunslinger] [Yordle] 1} {Darius [Knight] [Imperial] 1} {Garen [Knight] [Noble] 1} {Kayle [Knight] [Noble] 5} {Mordekaiser [Knight] [Phantom] 1} {Poppy [Knight] [Yordle] 3} {Sejuani [Knight] [Glacial] 4} {Ashe [Ranger] [Glacial] 3} {Kindred [Ranger] [Phantom] 4} {Varus [Ranger] [Demon] 2} {Vayne [Ranger] [Noble] 1} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4} {Jayce [Shapeshifter] [Hextech] 2} {Nidalee [Shapeshifter] [Wild] 1} {Shyvanna [Shapeshifter] [Dragon] 3} {Swain [Shapeshifter] [Demon Imperial] 5} {Ahri [Sorcerer] [Wild] 2} {Aurelion Sol [Sorcerer] [Dragon] 4} {Karthus [Sorcerer] [Phantom] 5} {Kassadin [Sorcerer] [Void] 1} {LuLu [Sorcerer] [Yordle] 2} {Morgana [Sorcerer] [Demon] 3} {Twisted Fate [Sorcerer] [Pirate] 2} {Veigar [Sorcerer] [Yordle] 3}]

// There are 56/74 champions (all): [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Gangplank [Blademaster Gunslinger] [Pirate] 3} {Shen [Blademaster] [Ninja] 2} {Yasuo [Blademaster] [Exile] 5} {Blitzcrank [Brawler] [Robot] 2} {Cho'gath [Brawler] [Void] 4} {Rek'Sai [Brawler] [Void] 2} {Vi [Brawler] [Hextech] 3} {Volibear [Brawler] [Glacial] 3} {Warwick [Brawler] [Wild] 1} {Anivia [Elementalist] [Glacial] 5} {Brand [Elementalist] [Demon] 4} {Kennen [Elementalist] [Ninja Yordle] 3} {Lissandra [Elementalist] [Glacial] 2} {Braum [Guardian] [Glacial] 2} {Leona [Guardian] [Noble] 4} {Pantheon [Guardian] [Dragon] 5} {Graves [Gunslinger] [Pirate] 1} {Lucian [Gunslinger] [Noble] 2} {Jinx [Gunslinger] [Hextech] 4} {Miss Fortune [Gunslinger] [Pirate] 5} {Tristana [Gunslinger] [Yordle] 1} {Darius [Knight] [Imperial] 1} {Garen [Knight] [Noble] 1} {Kayle [Knight] [Noble] 5} {Mordekaiser [Knight] [Phantom] 1} {Poppy [Knight] [Yordle] 3} {Sejuani [Knight] [Glacial] 4} {Ashe [Ranger] [Glacial] 3} {Kindred [Ranger] [Phantom] 4} {Varus [Ranger] [Demon] 2} {Vayne [Ranger] [Noble] 1} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4} {Jayce [Shapeshifter] [Hextech] 2} {Nidalee [Shapeshifter] [Wild] 1} {Shyvanna [Shapeshifter] [Dragon] 3} {Swain [Shapeshifter] [Demon Imperial] 5} {Ahri [Sorcerer] [Wild] 2} {Aurelion Sol [Sorcerer] [Dragon] 4} {Karthus [Sorcerer] [Phantom] 5} {Kassadin [Sorcerer] [Void] 1} {LuLu [Sorcerer] [Yordle] 2} {Morgana [Sorcerer] [Demon] 3} {Twisted Fate [Sorcerer] [Pirate] 2} {Veigar [Sorcerer] [Yordle] 3} {Veigar [Sorcerer] [Yordle] 3}]

// There are 47 champions after removing Sorcerers, [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Kha'Zix [Assassin] [Void] 1} {Pyke [Assassin] [Pirate] 2} {Rengar [Assassin] [Void] 3} {Zed [Assassin] [Ninja] 2} {Aatrox [Blademaster] [Demon] 3} {Camille [Blademaster] [Hextech] 1} {Draven [Blademaster] [Imperial] 4} {Gangplank [Blademaster Gunslinger] [Pirate] 3} {Shen [Blademaster] [Ninja] 2} {Yasuo [Blademaster] [Exile] 5} {Blitzcrank [Brawler] [Robot] 2} {Cho'gath [Brawler] [Void] 4} {Rek'Sai [Brawler] [Void] 2} {Vi [Brawler] [Hextech] 3} {Volibear [Brawler] [Glacial] 3} {Warwick [Brawler] [Wild] 1} {Anivia [Elementalist] [Glacial] 5} {Brand [Elementalist] [Demon] 4} {Kennen [Elementalist] [Ninja Yordle] 3} {Lissandra [Elementalist] [Glacial] 2} {Braum [Guardian] [Glacial] 2} {Leona [Guardian] [Noble] 4} {Pantheon [Guardian] [Dragon] 5} {Graves [Gunslinger] [Pirate] 1} {Lucian [Gunslinger] [Noble] 2} {Jinx [Gunslinger] [Hextech] 4} {Miss Fortune [Gunslinger] [Pirate] 5} {Tristana [Gunslinger] [Yordle] 1} {Darius [Knight] [Imperial] 1} {Garen [Knight] [Noble] 1} {Kayle [Knight] [Noble] 5} {Mordekaiser [Knight] [Phantom] 1} {Poppy [Knight] [Yordle] 3} {Sejuani [Knight] [Glacial] 4} {Ashe [Ranger] [Glacial] 3} {Kindred [Ranger] [Phantom] 4} {Varus [Ranger] [Demon] 2} {Vayne [Ranger] [Noble] 1} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4} {Jayce [Shapeshifter] [Hextech] 2} {Nidalee [Shapeshifter] [Wild] 1} {Shyvanna [Shapeshifter] [Dragon] 3} {Swain [Shapeshifter] [Demon Imperial] 5}]
