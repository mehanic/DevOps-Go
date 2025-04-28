package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Les Misérables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter and the Philisopher's Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	Book{
		ID:            6,
		Title:         "I, Robot",
		Author:        "Isaac Asimov",
		YearPublished: 1950,
	},
	Book{
		ID:            7,
		Title:         "The Gods Themselves",
		Author:        "Isaac Asimov",
		YearPublished: 1973,
	},
	Book{
		ID:            8,
		Title:         "The Moon is a Harsh Mistress",
		Author:        "Robert A. Heinlein",
		YearPublished: 1966,
	},
	Book{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	Book{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John Scalzi",
		YearPublished: 2006,
	},
}

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		// create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()

	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()

			return b, true
		}
	}

	return Book{}, false
}

// ─λ go run main.go                                                                        1 (0.058s) < 13:11:46
// from database
// Title:		"I, Robot"
// Author:		"Isaac Asimov"
// Published	1950

// from database
// Title:		"The Gods Themselves"
// Author:		"Isaac Asimov"
// Published	1973

// from database
// Title:		"On Basilisk Station"
// Author:		"David Weber"
// Published	1993

// from cache
// Title:		"On Basilisk Station"
// Author:		"David Weber"
// Published	1993

// from cache
// Title:		"I, Robot"
// Author:		"Isaac Asimov"
// Published	1950

// from database
// Title:		"The Hitchhiker's Guide to the Galaxy"
// Author:		"Douglas Adams"
// Published	1979

// from cache
// Title:		"The Hitchhiker's Guide to the Galaxy"
// Author:		"Douglas Adams"
// Published	1979

// from cache
// Title:		"On Basilisk Station"
// Author:		"David Weber"
// Published	1993

// from database
// Title:		"The Android's Dream"
// Author:		"John Scalzi"
// Published	2006

// from database
// Title:		"The Moon is a Harsh Mistress"
// Author:		"Robert A. Heinlein"
// Published	1966
