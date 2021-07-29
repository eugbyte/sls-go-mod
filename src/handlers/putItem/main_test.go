package main

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/serverless/sls-go-mod/src/data"
	"github.com/serverless/sls-go-mod/src/lib/util"
	"github.com/serverless/sls-go-mod/src/models"
)

func TestHandler(t *testing.T) {
	book := models.Book{
		Title:  "Percy Jackson & the Olympians",
		Author: "Rick Riordan",
	}
	mockRequest := Request{}
	mockRequest.Body = util.Stringify(book)
	response, err := Handler(&data.DynamoDBAdapter{}, mockRequest)
	if err != nil {
		t.Errorf("An error occured in Handler %v", err)
		return
	}

	if response.StatusCode != 200 {
		t.Errorf("test failed. Expected status code to be %d, but got %d", 200, response.StatusCode)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", response.Body)
		return
	}

	var actualBook models.Book
	err = json.Unmarshal([]byte(response.Body), &actualBook)
	if err != nil {
		t.Error("could not unmarshall book")
		return
	}

	if !cmp.Equal(book, actualBook) {
		t.Errorf("test failed. Expected book to be %v, but got %v", book, actualBook)
		t.Error("attempting to print out response body")
		util.Trace("response.Body:", actualBook)
	}

}
