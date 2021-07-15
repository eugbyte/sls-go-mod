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
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
	"github.com/serverless/sls-go-mod/src/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	var book models.Book
	err := json.Unmarshal([]byte(request.Body), &book)
	if err != nil {
		err = errors.Wrap(err, "Cannot unmarshall")
		log.Fatal(err)
		return Response{Body: err.Error(), StatusCode: http.StatusBadRequest}, err
	}

	book.Id = uuid.New().String()
	util.Trace("book", book)

	_, err = dynamoDBAdapter.Put("Book", book)
	if err != nil {
		err = errors.Wrap(err, "cannot put book")
		log.Fatal(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	response := Response{
		StatusCode:      201,
		IsBase64Encoded: false,
		Body:            string(request.Body),
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
