package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "time"
)


// PR is pull request info.
type PR struct {
    Number int
    Title  string
}

func parseResponse(r io.Reader) ([]PR, error) {
    var prs []PR
    if err := json.NewDecoder(r).Decode(&prs); err != nil {
        return nil, err
    }

    return prs, nil
}


func buildURL(owner, repo string, page int) string {
    // We want only open PRs to master
    query := url.Values{}
    query.Set("state", "open")
    query.Set("base", "master")
    query.Set("per_page", "100")
    query.Set("page", fmt.Sprintf("%d", page))

    owner, repo = url.PathEscape(owner), url.PathEscape(repo)
    const format = "https://api.github.com/repos/%s/%s/pulls?%s"
    return fmt.Sprintf(format, owner, repo, query.Encode())
}


func pageOpenPRs(owner, repo string, page int) ([]PR, error) {
    url := buildURL(owner, repo, page)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Accept", "application/vnd.github.v3+json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    } else if resp.StatusCode != http.StatusOK {
        const format = "bad status: %d - %s"
        return nil, fmt.Errorf(format, resp.StatusCode, resp.Status)
    }
    log.Printf("URL: %s", url)

    defer resp.Body.Close()
    r := io.LimitReader(resp.Body, 10*(1<<20)) // Don't read more than 10MB
    return parseResponse(r)
}


func openPRs(owner, repo string) ([]PR, error) {
    var prs []PR
    for page := 1; true; page++ {
        pagePRs, err := pageOpenPRs(owner, repo, page)
        if err != nil {
            return nil, err
        }

        if len(pagePRs) == 0 {
            break
        }
        prs = append(prs, pagePRs...)
    }

    return prs, nil
}


func main() {
    prs, err := openPRs("golang", "go")
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    fmt.Printf("%d open PRs\n", len(prs))
}
