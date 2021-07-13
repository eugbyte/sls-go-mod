package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/services/util"
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
		err = errors.Wrap(err, "Cannot unmarshall")
		log.Fatal(err)
		return Response{Body: err.Error(), StatusCode: http.StatusBadRequest}, err
	}
	Id := requestBody.Id
	util.Trace("id", Id)

	key := Attributes{
		"Id": {
			S: aws.String(Id),
		},
	}

	err = dynamoDBAdapter.Delete("Book", key)
	if err != nil {
		err = errors.Wrap(err, "cannot delete item")
		log.Fatal(err)
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	responseBody, _ := json.Marshal(map[string]string{
		"message": fmt.Sprintf("item %s succesfully deleted", Id),
	})

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
	return Handler(data.DynamoDBAdapter{}, request)
}

func main() {
	wrappedHandler := middleware.Middify(injectedHandler)
	lambda.Start(wrappedHandler)
}
