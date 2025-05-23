package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	GetResponseOutput  *http.Response
	PostResponseOutput *http.Response
}

func (m MockClient) Get(string) (resp *http.Response, err error) {
	return m.GetResponseOutput, nil
}

func (m MockClient) Post(string, string, io.Reader) (resp *http.Response, err error) {
	return m.PostResponseOutput, nil
}

func TestAPI_DoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b", "c"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}
	apiInstance := API{
		Options: Options{},
		Client: &MockClient{
			GetResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.DoGetRequest("http://localhost/words")
	if err != nil {
		t.Errorf("DoGetRequest error: %s", err)
	}
	if response == nil {
		t.Fatalf("response is empty")
	}
	if response.GetResponse() != strings.Join(words.Words.Words, ", ") {
		t.Errorf("Unexpected response: %s", response.GetResponse())
	}
}
