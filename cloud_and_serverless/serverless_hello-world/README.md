# AWS Lambda function handler in Go

The Lambda function handler is the method in our function code that processes events. When our function is invoked, Lambda runs the handler method. Our function runs until the handler returns a response, exits, or times out.

A Lambda function written in Go is authored as a Go executable. In our Lambda function code, you we to include the github.com/aws/aws-lambda-go/lambda package, which implements the Lambda programming model for Go. In addition, we need to implement handler function code and a main() function.

# Function Details

- `go mod init hello-world-lambda` : It creates hello-world-lambda module.

- Install go SDK with the command `go get -u github.com/aws/aws-lambda-go`.

- package main: In Go, the package containing func main() must always be named main.

- import: Use this to include the libraries our Lambda function requires. In this instance, it includes:

    - context: AWS Lambda context object in Go.

    - fmt: The Go Formatting object used to format the return value of our function.

    - github.com/aws/aws-lambda-go/lambda: As mentioned previously, implements the Lambda programming model for Go.

- func HandleRequest(ctx context.Context, event *lambdaParameters) (*string, error): This is the signature of our Lambda handler. It's the entry point for our Lambda function and contains the logic that is executed when our function is invoked. In addition, the parameters included denote the following:

    - ctx context.Context: Provides runtime information for our Lambda function invocation. ctx is the variable we declare to leverage the information available via AWS Lambda context object in Go.

    - event *lambdaParameters: This is a parameter named event that points to lambdaParameters. It represents the input to the Lambda function.

    - *string, error: The handler returns two values. The first is a pointer to a string which contains the result of the Lambda function. The second is an error type, which is nil if there's no error and contains standard error information if something goes wrong. For more information on custom error handling, see AWS Lambda function errors in Go.

- return &message, nil: Returns two values. The first is a pointer to a string message, which is a greeting constructed using the Name field from the input event. The second value, nil, indicates that the function didn't encounter any errors.

- func main(): The entry point that runs our Lambda function code. This is `required`.

- By adding lambda.Start(HandleRequest) between func main(){} code brackets, our Lambda function will be executed. Per Go language standards, the opening bracket, { must be placed directly at the end of the main function signature.

# Important Notes

- Before uploading the code to the AWS, we need to compile it, so lambda can understand our main code. To do that, type  `GOOS=linux go build main.go`. OS-only runtimes support both arm64 and x86_64. That's why we need to do that in linux.

- In order to upload our executable file to AWS, we need to zip it via `zip main.zip main`.

- Then go to AWS console and upload our zip file to the Lambda.

- In order to test the code locally, create a `.env` file in your main directory and create a parameter for `AWS_LAMBDA_FUNCTION_NAME`. This can be any value. Then call this variable via `lambdaFunctionName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")`

# Resources

- [Deploy Go Lambda functions with .zip file archives](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
- [AWS Lambda function handler in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [A Golang Lambda Function - Youtube](https://www.youtube.com/watch?v=KaydW7FO-Pw&ab_channel=CloudCasts)