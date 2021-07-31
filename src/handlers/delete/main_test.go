package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
)

func TestHandler(t *testing.T) {

	mockRequest := Request{}
	mockRequest.Body = util.Stringify(RequestBody{Id: "b8283220-8776-4401-b925-49153b5d5d0f"})

	response, err := Handler(&data.DynamoDBAdapter{}, mockRequest)
	if err != nil {
		t.Fatalf("An error occured with API Gateway %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("test failed. Expected status code to be %d, but got %d", 200, response.StatusCode)
		t.Errorf("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
		t.FailNow()
	}

	var messageBody map[string]string
	err = json.Unmarshal([]byte(response.Body), &messageBody)
	if err != nil {
		t.Fatalf("Cannot unmarshall response.Body")
	}

	message := messageBody["message"]
	id := "b8283220-8776-4401-b925-49153b5d5d0f"
	if !strings.Contains(message, id) {
		t.Errorf("test failed. Expected message to contain id %s", id)
		util.Trace("received message", message)
	}

}
