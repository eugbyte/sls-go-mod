package models

import (
	"log"

	colors "github.com/TwinProduction/go-color"
	"github.com/aws/aws-lambda-go/events"
	"github.com/serverless/sls-go-mod/src/lib/util"
)

type Response = events.APIGatewayProxyResponse

type HttpError struct {
	StatusCode int
	Err        error
}

func (er HttpError) Error() string {
	return er.Error()
}

func (er HttpError) Log() {
	// calling er.Err will result in stack overflow
	errorMessage := er.Err.Error()
	log.Println(colors.Red, errorMessage, colors.Reset)
}

func (er HttpError) ToResponse() Response {
	responseBody := util.Stringify(map[string]string{
		"message": er.Err.Error(),
	})
	return Response{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: er.StatusCode,
		Body:       responseBody,
	}
}

func (er HttpError) ToResponseAndLog() Response {
	er.Log()
	return er.ToResponse()
}
