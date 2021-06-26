package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/middleware"
	errs "github.com/serverless/sls-go-mod/src/models/custom_errors"
	"github.com/serverless/sls-go-mod/src/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest
type RequestBody struct {
	Message string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	util.Trace("body", request.Body)

	// BodyRequest will be used to take the json r esponse from client and build it
	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		notFoundError := errs.NewBadRequest(err, "cannot unmarshall request.Body")
		return Response{Body: notFoundError.Error(), StatusCode: notFoundError.StatusCode}, notFoundError
	}

	responseBody, err := json.Marshal((map[string]string{
		"message": requestBody.Message,
	}))
	if err != nil {
		internalError := errs.NewInternalServerError(err, "cannot marshall requestBody.Message")
		return Response{Body: internalError.Error(), StatusCode: internalError.StatusCode}, internalError
	}

	response := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return response, nil
}

// Dependency injection
func injectedHandler(request Request) (Response, error) {
	var dynamoDBAdapter data.IDynamoDBAdapter = data.DynamoDBAdapter{}
	return Handler(dynamoDBAdapter, request)
}

func main() {
	wrappedHandler := middleware.Middify(injectedHandler)
	lambda.Start(wrappedHandler)

}
