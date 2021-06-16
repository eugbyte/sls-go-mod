package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverless/sls-go-mod/services/util"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest
type BodyRequest struct {
	Message string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {

	util.Trace("In Handler")
	util.Trace(request)

	// BodyRequest will be used to take the json response from client and build it
	var bodyRequest BodyRequest

	// Unmarshal the json, return 404 if error
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		util.LogError(err)
		return Response{Body: err.Error(), StatusCode: 400}, err
	}

	responseBody, err := json.Marshal((map[string]string{
		"message": bodyRequest.Message,
	}))

	if err != nil {
		util.LogError(err)
		return Response{Body: err.Error(), StatusCode: 400}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
