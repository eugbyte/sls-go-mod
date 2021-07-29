package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/serverless/sls-go-mod/src/models"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

// The lambda handler can return 2 values. interface{} and error
// Note that for error handling, you must still return a Response
// https://stackoverflow.com/a/48462676/6514532
func Handler(request Request) (Response, error) {
	customErr := models.CustomError{StatusCode: http.StatusBadRequest, Err: errors.New("Custom Error Message!!!")}
	return customErr.ToResponse(), nil
}

func main() {
	lambda.Start(Handler)
}
