package main

import (
    "fmt"
    "net/http"
    "os"
    "regexp"
    "sync"
    "time"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    taskQueue := make(chan string, 100)
    foundUrlsQueue := make(chan []string, 100)
    var wg sync.WaitGroup

    // Initialize with a task for each process
    visited := map[string]bool{
        "/wiki/Kevin_Bacon": true,
        "/wiki/Monty_Python": true,
    }
    taskQueue <- "/wiki/Kevin_Bacon"
    taskQueue <- "/wiki/Monty_Python"

    wg.Add(1)
    go taskDelegator(taskQueue, foundUrlsQueue, visited, &wg)

    for i := 0; i < 2; i++ {
        wg.Add(1)
        go scrapeArticle(taskQueue, foundUrlsQueue, &wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()
}

func taskDelegator(taskQueue chan string, foundUrlsQueue chan []string, visited map[string]bool, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case links := <-foundUrlsQueue:
            for _, link := range links {
                if !visited[link] {
                    visited[link] = true
                    taskQueue <- link
                }
            }
        }
    }
}

func getLinks(doc *goquery.Document) []string {
    var links []string
    doc.Find("div#bodyContent a[href^='/wiki/']").Each(func(i int, s *goquery.Selection) {
        href, _ := s.Attr("href")
        if matched, _ := regexp.MatchString(`^(/wiki/)((?!:).)*$`, href); matched {
            links = append(links, href)
        }
    })
    return links
}

func scrapeArticle(taskQueue chan string, foundUrlsQueue chan []string, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        path, ok := <-taskQueue
        if !ok {
            return
        }

        resp, err := http.Get("http://en.wikipedia.org" + path)
        if err != nil {
            fmt.Println("Error fetching:", err)
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode != 200 {
            fmt.Println("Error: Status code", resp.StatusCode)
            continue
        }

        // Parse the HTML response
        doc, err := goquery.NewDocumentFromReader(resp.Body)
        if err != nil {
            fmt.Println("Error parsing document:", err)
            continue
        }

        title := doc.Find("h1").Text()
        fmt.Printf("Scraping %s in process %d\n", title, os.Getpid())

        links := getLinks(doc)
        foundUrlsQueue <- links
        time.Sleep(5 * time.Second) // Simulating the same delay as in Python
    }
}

