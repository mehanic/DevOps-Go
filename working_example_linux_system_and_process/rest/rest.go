package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := Exec(); err != nil {
		panic(err)
	}
}

//---

// APIClient is our custom client
type APIClient struct {
	*http.Client
}

// NewAPIClient constructor initializes the client with our
// custom Transport
func NewAPIClient(username, password string) *APIClient {
	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Transport: &t,
				username:  username,
				password:  password,
			},
		},
	}
}

// GetGoogle is an API Call - we abstract away
// the REST aspects
func (c *APIClient) GetGoogle() (int, error) {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

//----

// Exec creates an API Client and uses its
// GetGoogle method, then prints the result
func Exec() error {
	c := NewAPIClient("username", "password")

	StatusCode, err := c.GetGoogle()
	if err != nil {
		return err
	}
	fmt.Println("Result of GetGoogle:", StatusCode)
	return nil
}

//---

// APITransport does a SetBasicAuth
// for every request
type APITransport struct {
	*http.Transport
	username, password string
}

// RoundTrip does the basic auth before deferring to the
// default transport
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.Transport.RoundTrip(req)
}
