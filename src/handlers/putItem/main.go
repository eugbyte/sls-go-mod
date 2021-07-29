package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	var book models.Book
	err := json.Unmarshal([]byte(request.Body), &book)
	if err != nil {
		err = errors.Wrap(err, "Cannot unmarshall")
		log.Fatal(err)
		httpError := models.HttpError{Err: errors.Wrap(err, "Cannot unmarshall"), StatusCode: http.StatusBadRequest}
		return httpError.ToResponse(), nil
	}

	book.Id = uuid.New().String()
	util.Trace("book", book)

	// To create - additional check otherwise dynamodb will just overwrite
	// conditionExpression := "attribute_not_exists(Id)"

	_, err = dynamoDBAdapter.Put("Book", book, nil)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "cannot put book"), StatusCode: http.StatusInternalServerError}
		return httpError.ToResponseAndLog(), nil
	}

	response := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            request.Body,
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
