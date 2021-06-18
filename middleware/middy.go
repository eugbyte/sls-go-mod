package middleware

import (
	"fmt"

	colors "github.com/TwinProduction/go-color"
	"github.com/aws/aws-lambda-go/events"
	"github.com/serverless/sls-go-mod/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Middify(handler func(request Request) (Response, error)) func(request Request) (Response, error) {
	return func(request Request) (Response, error) {
		util.Trace("middleware", "In Middleware...")
		// Logic to preprocess request here...
		util.Trace("handler", "proceeding to handler...")
		response, err := handler(request)

		// Logic to process response and error here
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
		}

		return response, err
	}
}
