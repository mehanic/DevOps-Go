package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type lambdaParameters struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

func main() {

	lambdaFunctionName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME") // This is to test locally.

	if lambdaFunctionName != "" { // if the environment variable is set, assume that we are running in lambda
		lambda.Start(HandleRequest)
	} else {
		params := lambdaParameters{
			Message: "Hello World from Berkay",
			Count:   2,
		}

		HandleRequest(params)
	}
}

func HandleRequest(event lambdaParameters) (*string, error) {

	for index := 0; index < event.Count; index++ {
		fmt.Println(event.Message)
	}

	message := fmt.Sprintf(event.Message)

	return &message, nil
}
