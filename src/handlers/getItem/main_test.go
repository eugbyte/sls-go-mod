package main

import (
	"testing"

	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
)

func TestHandler(t *testing.T) {

	mockRequest := Request{}
	mockRequest.PathParameters = map[string]string{
		"Id": "1768dc61-e6f0-4a6c-8cfb-2f4160edff47",
	}

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
