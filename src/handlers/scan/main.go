package main

import (
	"encoding/json"
	"log"

	errs "github.com/serverless/sls-go-mod/src/models/custom_errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest
type Attributes = map[string]*dynamodb.AttributeValue

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {
	exprsn, err := expression.NewBuilder().Build()
	if err != nil {
		log.Fatal("Got error building expression" + err.Error())
		internalError := errs.NewInternalServerError(err, "Got error building expression")
		return Response{Body: internalError.Error(), StatusCode: 500}, err
	}

	var books []models.Book
	err = dynamoDBAdapter.Scan("Book", exprsn, &books)

	responseBody, err := json.Marshal(books)
	if err != nil {
		internalError := errs.NewInternalServerError(err, "cannot marshall requestBody")
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
