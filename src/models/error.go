package models

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/serverless/sls-go-mod/src/lib/util"
)

type Response = events.APIGatewayProxyResponse

type CustomError struct {
	StatusCode int
	Err        error
}

func (er *CustomError) Error() string {
	return er.Error()
}

func (er *CustomError) ToResponse() Response {
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
