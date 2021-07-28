package middleware

import (
	"fmt"

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

		// Logic to process response and error here...
		fmt.Println(colors.Green, "middleware: post-processing...", colors.Reset)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
		}

		return response, err
	}
}
