package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/middleware"
	"github.com/serverless/sls-go-mod/src/models"
)

type Response = events.APIGatewayProxyResponse
type Request = events.APIGatewayProxyRequest
type RequestBody struct {
	Message string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {

	util.Trace("body", request.Body)

	// BodyRequest will be used to take the json response from client and build it
	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		httpError := models.HttpError{Err: errors.Wrap(err, "Cannot unmarshall: "), StatusCode: http.StatusBadRequest}
		return httpError.ToResponseAndLog(), nil
	}

	message := requestBody.Message
	message = strings.ToUpper(message) + "!!"
	responseBody := util.Stringify(map[string]string{
		"message": message,
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

func main() {
	wrappedHandler := middleware.Middify(Handler)
	lambda.Start(wrappedHandler)

}
