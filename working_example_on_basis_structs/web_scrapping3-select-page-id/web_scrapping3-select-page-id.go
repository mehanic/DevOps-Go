package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Custom error to indicate when the solution is found
type SolutionFound struct {
	Message string
}

func (e *SolutionFound) Error() string {
	return e.Message
}

// Connect to MySQL database
func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wikipedia")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

// Retrieve links from the database based on fromPageId
func getLinks(db *sql.DB, fromPageId int) []int {
	rows, err := db.Query("SELECT toPageId FROM links WHERE fromPageId = ?", fromPageId)
	if err != nil {
		log.Fatal("Query execution failed:", err)
	}
	defer rows.Close()

	var links []int
	for rows.Next() {
		var toPageId int
		if err := rows.Scan(&toPageId); err != nil {
			log.Fatal("Error reading row:", err)
		}
		links = append(links, toPageId)
	}
	if len(links) == 0 {
		return nil
	}
	return links
}

// Construct a dictionary for the current page's links
func constructDict(db *sql.DB, currentPageId int) map[int]map[int]struct{} {
	links := getLinks(db, currentPageId)
	if links == nil {
		return nil
	}
	linkTree := make(map[int]map[int]struct{})
	for _, link := range links {
		linkTree[link] = make(map[int]struct{})
	}
	return linkTree
}

// Search through the link tree recursively for the targetPageId
func searchDepth(db *sql.DB, targetPageId, currentPageId int, linkTree map[int]map[int]struct{}, depth int) (map[int]map[int]struct{}, error) {
	if depth == 0 {
		return linkTree, nil
	}
	if linkTree == nil {
		linkTree = constructDict(db, currentPageId)
		if linkTree == nil {
			return nil, nil
		}
	}
	if _, found := linkTree[targetPageId]; found {
		fmt.Println("TARGET", targetPageId, "FOUND!")
		return nil, &SolutionFound{Message: fmt.Sprintf("PAGE: %d", currentPageId)}
	}
	for branchKey := range linkTree {
		var err error
		linkTree[branchKey], err = searchDepth(db, targetPageId, branchKey, linkTree[branchKey], depth-1)
		if err != nil {
			if _, ok := err.(*SolutionFound); ok {
				fmt.Println(err.Error())
				return nil, &SolutionFound{Message: fmt.Sprintf("PAGE: %d", currentPageId)}
			} else {
				return nil, err
			}
		}
	}
	return linkTree, nil
}

func main() {
	db := connectDB()
	defer db.Close()

	// Search the tree for the target page
	targetPageId := 134951
	startPageId := 1
	depth := 4

	linkTree := make(map[int]map[int]struct{})
	err := func() error {
		_, err := searchDepth(db, targetPageId, startPageId, linkTree, depth)
		if err != nil {
			if _, ok := err.(*SolutionFound); ok {
				return nil
			}
			return err
		}
		fmt.Println("No solution found")
		return nil
	}()
	if err != nil {
		fmt.Println("Error during search:", err)
		os.Exit(1)
	}
}

