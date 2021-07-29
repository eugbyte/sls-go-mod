package main

import (
	"encoding/json"
	"testing"

	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/models"
)

func TestHandler(t *testing.T) {

	mockRequest := Request{}

	response, err := Handler(&data.DynamoDBAdapter{}, mockRequest)
	if err != nil {
		t.Errorf("An error occured in Handler %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("test failed. Expected status code to be %d, but got %d", 200, response.StatusCode)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
	}

	var actualBooks []models.Book
	err = json.Unmarshal([]byte(response.Body), &actualBooks)
	if err != nil {
		t.Error("failed to unmarshall books")
	}
	util.Trace("actualBooks", actualBooks)

	if len(actualBooks) < 1 {
		t.Errorf("test failed. Expected number of books to 1 or more, but got %d", len(actualBooks))
	}

}
