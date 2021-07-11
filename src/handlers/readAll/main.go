package main

import (
	"log"

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
	exp, err := expression.NewBuilder().Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	var books []models.Book, err := dynamoDBAdapter.Scan("Book", exp, []models.Book{})

	
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
