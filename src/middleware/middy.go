package middleware

import (
	"fmt"

	colors "github.com/TwinProduction/go-color"
	"github.com/aws/aws-lambda-go/events"
	"github.com/serverless/sls-go-mod/src/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Middify(handler func(request Request) (Response, error)) func(request Request) (Response, error) {
	return func(request Request) (Response, error) {
		// Logic to preprocess request here!...
		fmt.Println(colors.Blue, "middleware: pre-processing...", colors.Reset)
		fmt.Println(colors.Blue, "proceeding to handler...", colors.Reset)
		response, err := handler(request)

		util.Trace("middleware", "post-processing...")
		// Logic to process response and error here
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
		}

		return response, err
	}
}
