package main

import (
	"log"
)

type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}

func main() {
	akali, ev, kat, elise, gnar := createChampions()

	// Adding to an array, you can't add, set allocated elements
	champs := [3]champion{}
	champs[0] = akali
	champs[1] = ev
	champs[2] = kat
	//champs[3] = elise

	log.Printf("There are %d champs, %v\n\n", len(champs), champs)

	assassins := []champion{}
	log.Printf("assassins length=%d, capacity=%d\n\n", len(assassins), cap(assassins))
	assassins = append(assassins, akali, ev, kat)
	log.Printf("assassins length=%d, capacity=%d, contents=%v\n\n", len(assassins), cap(assassins), assassins)

	shapeshifters := []champion{elise, gnar}
	log.Printf(
		"shapeshifters length=%d, capacity=%d, contents=%v\n\n",
		len(shapeshifters),
		cap(shapeshifters),
		shapeshifters,
	)

	allChamps := []champion{}
	allChamps = append(allChamps, assassins...)
	allChamps = append(allChamps, shapeshifters...)
	log.Printf("allChamps length=%d, capacity=%d, contents=%v\n\n", len(allChamps), cap(allChamps), allChamps)
}

func createChampions() (champion, champion, champion, champion, champion) {
	akali := champion{
		Name:    "Akali",
		Classes: []string{"Assassin"},
		Origins: []string{"Ninja"},
		Cost:    4,
	}

	ev := champion{
		Name:    "Evelynn",
		Classes: []string{"Assassin"},
		Origins: []string{"Demon"},
		Cost:    3,
	}

	kat := champion{
		Name:    "Katarina",
		Classes: []string{"Assassin"},
		Origins: []string{"Imperial"},
		Cost:    3,
	}

	elise := champion{
		Name:    "Elise",
		Classes: []string{"Shapeshifter"},
		Origins: []string{"Demon"},
		Cost:    1,
	}

	gnar := champion{
		Name:    "Gnar",
		Classes: []string{"Shapeshifter"},
		Origins: []string{"Wild", "Yordle"},
		Cost:    4,
	}

	return akali, ev, kat, elise, gnar
}

// 2024/10/02 21:40:16 There are 3 champs, [{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3}]

// 2024/10/02 21:40:16 assassins length=0, capacity=0

// 2024/10/02 21:40:16 assassins length=3, capacity=3, contents=[{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3}]

// 2024/10/02 21:40:16 shapeshifters length=2, capacity=2, contents=[{Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4}]

// 2024/10/02 21:40:16 allChamps length=5, capacity=6, contents=[{Akali [Assassin] [Ninja] 4} {Evelynn [Assassin] [Demon] 3} {Katarina [Assassin] [Imperial] 3} {Elise [Shapeshifter] [Demon] 1} {Gnar [Shapeshifter] [Wild Yordle] 4}]
