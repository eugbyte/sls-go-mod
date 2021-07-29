package main

import (
	"encoding/json"
	"fmt"
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
type RequestBody struct {
	Id string `json:"id"`
}

func Handler(dynamoDBAdapter data.IDynamoDBAdapter, request Request) (Response, error) {

	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "Cannot unmarshall"), StatusCode: http.StatusBadRequest}
		return httpError.ToResponseAndLog(), nil
	}
	Id := requestBody.Id
	util.Trace("id", Id)

	key := Attributes{
		"Id": {
			S: aws.String(Id),
		},
	}

	conditionExpression := "attribute_exists(Id)"
	err = dynamoDBAdapter.Delete("Book", key, &conditionExpression)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "cannot delete item"), StatusCode: http.StatusInternalServerError}
		return httpError.ToResponseAndLog(), nil
	}

	responseBody := util.Stringify(map[string]string{
		"message": fmt.Sprintf("item %s succesfully deleted", Id),
	})

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
