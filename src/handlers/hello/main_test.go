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
		t.Errorf("An error occured in Handler %v", err)
	}

	var messageNap map[string]string
	err = json.Unmarshal([]byte(response.Body), &messageNap)
	if err != nil {
		t.Error("Error unmarshalling response body to map[string]string")
	}
	message := messageNap["message"]
	if messageNap["message"] != "HELLO!!" {
		t.Errorf("test failed. Expected %v, received %v", message, "Hello!!")
	} else {
		t.Logf("test passed. Expected %v, received %v", message, "Hello!!")
	}

}
