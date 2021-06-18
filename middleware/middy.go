package middleware

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/serverless/sls-go-mod/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Middify(handler func(request Request) (Response, error)) func(request Request) (Response, error) {
	return func(request Request) (Response, error) {
		// Logic to preprocess request here...
		response, err := handler(request)

		// Logic to process response and error here
		if err != nil {
			util.LogError(err.Error())
		}

		return response, err
	}
}
