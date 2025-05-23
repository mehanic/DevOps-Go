package main
import (
    "context"
    "io"
    "fmt"
    "net/http"
    "github.com/aws/aws-lambda-go/lambda"
)    
type Response struct {
    StatusCode int  `json:"statusCode"`
    Body    string `json:"body"`
}
func HandleRequest(ctx context.Context, _ interface{}) (Response, error) {
    resp, err := http.Get("http://localhost:2345")
    fmt.Println("Got response", resp)
    if err != nil {
     return Response{StatusCode: 500, Body: "Something bad happened"}, err
    }
    if resp.StatusCode != 200 {
     return Response{StatusCode: resp.StatusCode, Body: resp.Status}, err
    }
    defer resp.Body.Close()
    b, err := io.ReadAll(resp.Body)
    if err != nil {
     return Response{StatusCode: 500, Body: "Error decoding body"}, err
    }
    return Response{StatusCode: 200, Body: string(b)}, nil
}
func main() {
    lambda.Start(HandleRequest)
}
