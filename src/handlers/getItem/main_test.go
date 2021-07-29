package main

import (
	"encoding/json"
	"testing"

	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/models"
)

func TestHandler(t *testing.T) {

	id := "1768dc61-e6f0-4a6c-8cfb-2f4160edff47"
	mockRequest := Request{}
	mockRequest.PathParameters = map[string]string{
		"Id": id,
	}

	response, err := Handler(&data.DynamoDBAdapter{}, mockRequest)
	if err != nil {
		t.Errorf("An error occured with API Gateway %v", err)
		return
	}

	if response.StatusCode != 200 {
		t.Errorf("test failed. Expected status code to be %d, but got %d", 200, response.StatusCode)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
		return
	}
	var book models.Book
	err = json.Unmarshal([]byte(response.Body), &book)
	if book.Id != id {
		t.Errorf("test failed. Expected book id to be %s, but got %s", id, book.Id)
	}

}
