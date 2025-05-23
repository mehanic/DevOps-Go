package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

//import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/client"

func main() {
	// secure and op!
	cli := Setup(true, false)

	if err := DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := DoOps(cli); err != nil {
		panic(err)
	}

	c := Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	// secure and noop
	// also modifies default
	Setup(true, true)

	if err := DefaultGetGolang(); err != nil {
		panic(err)
	}
}

//-----

// Setup configures our client and redefines
// the global DefaultClient
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient

	// Sometimes for testing, we want to
	// turn off SSL verification
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c
	return c
}

// NopTransport is a No-Op Transport
type NopTransport struct {
}

// RoundTrip Implements RoundTripper interface
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	// note this is an unitialized Response
	// if you're looking at headers etc
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}

//---

// DoOps takes a client, then fetches
// google.com
func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of DoOps:", resp.StatusCode)

	return nil
}

// DefaultGetGolang uses the default client
// to get golang.org
func DefaultGetGolang() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}
	fmt.Println("results of DefaultGetGolang:", resp.StatusCode)
	return nil
}

//----

// Controller embeds an http.Client
// and uses it internally
type Controller struct {
	*http.Client
}

// DoOps with a controller object
func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of client.DoOps", resp.StatusCode)
	return nil
}
