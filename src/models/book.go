package models

type Book struct {
	Id     string `json:"id" dynamodbav:"Id" `
	Title  string `json:"title" dynamodbav:"Title" `
	Author string `json:"author" dynamodbav:"Author" `
}
