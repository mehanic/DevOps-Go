package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


func main() {
	if err := Exec(); err != nil {
		panic(err)
	}
}



// Setup initializes our ClientInterface
func Setup() *http.Client {
	c := http.Client{}

	t := Decorate(&http.Transport{},
		Logger(log.New(os.Stdout, "", 0)),
		BasicAuth("username", "password"),
	)
	c.Transport = t
	return &c
}



// TransportFunc implements the RountTripper interface
type TransportFunc func(*http.Request) (*http.Response, error)

// RoundTrip just calls the original function
func (tf TransportFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return tf(r)
}

// Decorator is a convenience function to represent our
// middleware inner function
type Decorator func(http.RoundTripper) http.RoundTripper

// Decorate is a helper to wrap all the middleware
func Decorate(t http.RoundTripper, rts ...Decorator) http.RoundTripper {
	decorated := t
	for _, rt := range rts {
		decorated = rt(decorated)
	}
	return decorated
}

//----

// Exec creates a client, calls google.com
// then prints the response
func Exec() error {
	c := Setup()

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("Response code:", resp.StatusCode)
	return nil
}

//---

// Logger is one of our 'middleware' decorators
func Logger(l *log.Logger) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			start := time.Now()
			l.Printf("started request to %s at %s", r.URL, start.Format("2006-01-02 15:04:05"))
			resp, err := c.RoundTrip(r)
			l.Printf("completed request to %s in %s", r.URL, time.Since(start))
			return resp, err
		})
	}
}

// BasicAuth is another of our 'middleware' decorators
func BasicAuth(username, password string) Decorator {
	return func(c http.RoundTripper) http.RoundTripper {
		return TransportFunc(func(r *http.Request) (*http.Response, error) {
			r.SetBasicAuth(username, password)
			resp, err := c.RoundTrip(r)
			return resp, err
		})
	}
}

// started request to https://www.google.com at 2025-01-17 21:28:34
// completed request to https://www.google.com in 142.055327ms
// Response code: 200
