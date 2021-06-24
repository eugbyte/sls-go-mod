package data

import (
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pkg/errors"
)

type Person struct{}

type Attributes = map[string]*dynamodb.AttributeValue

type DynamoDBAdapter struct {
	TableName string
}

type IDynamoDBAdapter interface {
	Create(obj interface{}) (interface{}, error)
	Read(key Attributes, out interface{}) error
	Update(updateItemInput dynamodb.UpdateItemInput) error
	Delete(deleteInput *dynamodb.DeleteItemInput) error
}

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var currentSession = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var client = dynamodb.New(currentSession)

func (adapter DynamoDBAdapter) Create(obj interface{}) (interface{}, error) {
	item, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(adapter.TableName),
	}

	_, err = client.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return obj, nil
}

func (adapter DynamoDBAdapter) Read(key Attributes, outPointer interface{}) error {

	if reflect.ValueOf(outPointer).Kind() == reflect.Ptr {
		err := errors.New("out argument must be a pointer")
		log.Fatalf(err.Error())
		return err
	}

	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(adapter.TableName),
		Key:       key,
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
		return err
	}

	var item Attributes = result.Item

	if item == nil {
		log.Fatalf("Could not find item")
		return err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &outPointer)
	if err != nil {
		log.Fatalf("Failed to unmarshal Record")
		return err
	}

	return nil
}

func (adapter DynamoDBAdapter) Update(updateInput *dynamodb.UpdateItemInput) error {

	if adapter.TableName != *updateInput.TableName {
		error := errors.New("tablename provided must match that initialized with DynamoDBAdapter")
		return error
	}

	_, err := client.UpdateItem(updateInput)
	if err != nil {
		error := errors.Errorf("Got error calling UpdateItem")
		return error
	}
	return nil
}

func (adapter DynamoDBAdapter) Delete(deleteInput *dynamodb.DeleteItemInput) error {
	if len(*deleteInput.TableName) > 0 && adapter.TableName != *deleteInput.TableName {
		error := errors.New("tablename provided must match that initialized with DynamoDBAdapter")
		return error
	}

	_, err := client.DeleteItem(deleteInput)
	if err != nil {
		error := errors.Errorf("Got error calling UpdateItem")
		return error
	}
	return nil
}

func (adapter DynamoDBAdapter) Scan() error {
	expression.NewBuilder()
	return nil
}
