package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/serverless/sls-go-mod/src/lib/util"
)

func TestHandler(t *testing.T) {
	response, err := Handler(Request{})
	if err != nil {
		t.Errorf("An error occured with API Gateway %v", err)
		return
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("test failed. Expected status code to be %d, but got %d", http.StatusBadRequest, response.StatusCode)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
		return
	}

	var messageBody map[string]string
	err = json.Unmarshal([]byte(response.Body), &messageBody)
	if err != nil {
		t.Error("Cannot unmarshall response.Body")
		return
	}

	message := messageBody["message"]
	expectedMessage := "Custom Error Message!!!"
	if message != expectedMessage {
		t.Errorf("test failed. Expected error message to be %s, but got %s", expectedMessage, message)
	}

}
