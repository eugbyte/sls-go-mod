package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	var Id string = request.PathParameters["Id"]
	util.Trace("Id", Id)

	key := Attributes{
		"Id": {
			S: aws.String(Id),
		},
	}

	var book models.Book
	err := dynamoDBAdapter.GetItem("Book", key, &book)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "Cannot find book: "+Id), StatusCode: http.StatusBadRequest}
		return httpError.ToResponseAndLog(), nil
	}

	responseBody := util.Stringify(book)

	response := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            responseBody,
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
