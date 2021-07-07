package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
	errs "github.com/serverless/sls-go-mod/src/models/custom_errors"
	"github.com/serverless/sls-go-mod/src/services/util"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	util.Trace("body", request.Body)

	var book models.Book
	err := json.Unmarshal([]byte(request.Body), &book)
	if err != nil {
		notFoundError := errs.NewBadRequest(err, "cannot unmarshall request.Body")
		return Response{Body: notFoundError.Error(), StatusCode: notFoundError.StatusCode}, notFoundError
	}

	book.Id = uuid.New().String()
	util.Trace("book", book)

	_, err = dynamoDBAdapter.Put("Book", book)
	if err != nil {
		util.LogError(err)
		fmt.Println(err.Error() + "Fail")
		// internalError := errs.NewInternalServerError(err, "cannot put book")
		return Response{Body: "cannot put book", StatusCode: 500}, err
	}

	response := Response{
		StatusCode:      200,
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
	var dynamoDBAdapter data.IDynamoDBAdapter = data.DynamoDBAdapter{}
	return Handler(dynamoDBAdapter, request)
}

func main() {
	wrappedHandler := middleware.Middify(injectedHandler)
	lambda.Start(wrappedHandler)

}
