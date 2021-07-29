package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pkg/errors"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest
type Attributes = map[string]*dynamodb.AttributeValue

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	// unfortunately, doing just expression.NewBuilder().Build(), without need at least one expression, e.g. .WithFilter(), returns an error
	// this error can be ignored
	// https://github.com/aws/aws-sdk-go/blob/7798c2e0edc02ba058f7672d32f4ebf6603b5fc6/service/dynamodb/expression/expression.go#L101
	expr, _ := expression.NewBuilder().Build()

	var books []models.Book
	err := dynamoDBAdapter.Scan("Book", expr, &books)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "cannot scan"), StatusCode: http.StatusBadRequest}
		return httpError.ToResponse(), nil
	}

	responseBody := util.Stringify(books)

	response := Response{
		StatusCode: 200,
		Body:       responseBody,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return response, nil
}

// Dependency injection
func injectedHandler(request Request) (Response, error) {
	return Handler(&data.DynamoDBAdapter{}, request)
}

func main() {
	wrappedHandler := middleware.Middify(injectedHandler)
	lambda.Start(wrappedHandler)
}
