package main

import (
	"testing"

	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
)

func TestHandler(t *testing.T) {

	mockRequest := Request{}
	mockRequest.Body = util.Stringify(RequestBody{Id: "b8283220-8776-4401-b925-49153b5d5d0f"})

	response, err := Handler(&data.DynamoDBAdapter{}, mockRequest)
	if err != nil {
		t.Errorf("An error occured in Handler %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("test failed. Expected status code to be %d, but got %d", 200, response.StatusCode)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
	}

}
