package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	baseURL    = "https://openlibrary.org"
	endpoint   = "/api/books"
	numWorkers = 2
)

type Job struct {
	URL string
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	urls := []string{
		baseURL + endpoint + "?bibkeys=ISBN:9780134190440&format=json",
		baseURL + endpoint + "?bibkeys=ISBN:978-1491941195&format=json",
		baseURL + endpoint + "?bibkeys=ISBN:0134494164&format=json",
	}
	jobs := make(chan Job, len(urls))
	results := make(chan string, len(urls))

	// Start workers
	c := new(http.Client)
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		eg.Go(func() error {
			return worker(ctx, &wg, c, jobs, results)
		})
	}

	// Sending jobs to the "jobs" channel
	wg.Add(len(urls))
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs)
	wg.Wait()

	// Wait for all requests to complete
	if err := eg.Wait(); err != nil {
		log.Printf("Error occurred: %v", err)
	}

	// Collecting results
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}

func worker(
	ctx context.Context, wg *sync.WaitGroup, c *http.Client,
	jobs <-chan Job, results chan<- string,
) error {
	for job := range jobs {
		defer wg.Done()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, job.URL, nil)
		if err != nil {
			return err
		}

		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var response interface{}
		json.Unmarshal(body, &response)
		// fmt.Printf("%T\n", response)

		results <- fmt.Sprintf("%d: %s", resp.StatusCode, response)
	}

	return nil
}

// λ go run main.go                                                                                                               0 (0.001s) < 14:08:49
// 200: map[ISBN:9780134190440:map[bib_key:ISBN:9780134190440 info_url:https://openlibrary.org/books/OL25925901M/The_Go_Programming_Language preview:noview preview_url:https://openlibrary.org/books/OL25925901M/The_Go_Programming_Language thumbnail_url:https://covers.openlibrary.org/b/id/7427319-S.jpg]]
// 200: map[ISBN:978-1491941195:map[bib_key:ISBN:978-1491941195 info_url:https://openlibrary.org/books/OL26834447M/Concurrency_in_Go_Tools_and_Techniques_for_Developers preview:noview preview_url:https://openlibrary.org/books/OL26834447M/Concurrency_in_Go_Tools_and_Techniques_for_Developers thumbnail_url:https://covers.openlibrary.org/b/id/8510272-S.jpg]]
// 200: map[ISBN:0134494164:map[bib_key:ISBN:0134494164 info_url:https://openlibrary.org/books/OL31838215M/Clean_Architecture preview:noview preview_url:https://openlibrary.org/books/OL31838215M/Clean_Architecture]]
