package main

import (
	"encoding/json"
	"testing"

	"github.com/serverless/sls-go-mod/src/lib/util"
)

func TestHandler(t *testing.T) {
	mockRequest := Request{}
	mockRequest.Body = util.Stringify(RequestBody{Message: "Hello"})
	response, err := Handler(mockRequest)
	if err != nil {
		t.Fatalf("An error occured with API Gateway %v", err)
		return
	}

	var messageNap map[string]string
	err = json.Unmarshal([]byte(response.Body), &messageNap)
	if err != nil {
		t.Fatal("Error unmarshalling response body to map[string]string")
	}
	message := messageNap["message"]
	if message != "HELLO!!" {
		t.Fatalf("test failed. Expected %v, received %v", "HELLO!!", message)
	} else {
		t.Logf("test passed. Expected %v, received %v", "HELLO!!", message)
	}

}
