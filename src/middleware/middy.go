package middleware

import (
	"fmt"
	"log"

	colors "github.com/TwinProduction/go-color"
	"github.com/aws/aws-lambda-go/events"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Middify(handler func(request Request) (Response, error)) func(request Request) (Response, error) {
	return func(request Request) (Response, error) {
		// Logic to preprocess request here...
		fmt.Println(colors.Green, "middleware: pre-processing...", colors.Reset)

		response, err := handler(request)

		// Logic to log response and error here...
		// Note that for lambda, the error should return nil, as the response body should already contaib the expected errors
		// https://stackoverflow.com/a/48462676/6514532
		fmt.Println(colors.Green, "middleware: post-processing...", colors.Reset)
		if err != nil {
			log.Println(colors.Red, err, colors.Reset)
		}

		return response, nil
	}
}
