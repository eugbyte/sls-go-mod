package data

import (
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pkg/errors"
	localConfig "github.com/serverless/sls-go-mod/src/lib/config"
)

type Attributes = map[string]*dynamodb.AttributeValue

type DynamoDBAdapter struct {
}

type IDynamoDBAdapter interface {
	Put(tableName string, obj interface{}, condtionExpression *string) (interface{}, error)
	GetItem(tableName string, key Attributes, outPointer interface{}) error
	Update(updateInput *dynamodb.UpdateItemInput) error
	Delete(tableName string, key Attributes, conditionExpression *string) error
	Scan(tableName string, expr expression.Expression, outslice interface{}) error
}

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.

// session.NewSessionWithOptions(session.Options{
// 	SharedConfigState: session.SharedConfigEnable,
// })
var currentSession = session.Must(session.NewSession())

var config = localConfig.Config

var awsConfig = aws.NewConfig().
	WithRegion(config.REGION).
	WithCredentials(
		credentials.NewStaticCredentials(config.ACCESS_KEY_ID, config.SECRET_ACCESS_KEY, ""),
	).
	WithEndpoint(config.DYNAMO_DB_ENDPOINT)

// Create DynamoDB Client
var Client = dynamodb.New(currentSession, awsConfig)

// When an existing item found, Put replaces it with the new one
func (adapter *DynamoDBAdapter) Put(tableName string, obj interface{}, condtionExpression *string) (interface{}, error) {

	item, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:                item,
		TableName:           aws.String(tableName),
		ConditionExpression: condtionExpression,
	}

	_, err = Client.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return nil, err
	}

	return obj, nil
}

func (adapter *DynamoDBAdapter) GetItem(tableName string, key Attributes, outPointer interface{}) error {

	if reflect.ValueOf(outPointer).Kind() != reflect.Ptr {
		err := errors.New("out argument must be a pointer")
		log.Fatalf(err.Error())
		return err
	}

	result, err := Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
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

	err = dynamodbattribute.UnmarshalMap(result.Item, outPointer)
	if err != nil {
		log.Fatalf("Failed to unmarshal Record")
		return err
	}

	return nil
}

func (adapter *DynamoDBAdapter) Update(updateInput *dynamodb.UpdateItemInput) error {

	_, err := Client.UpdateItem(updateInput)
	if err != nil {
		error := errors.Errorf("Got error calling UpdateItem: %v", err)
		return error
	}
	return nil
}

// conditionExpression - optional, can set to nil
// expressionAttributeNames - optional, can set to nil
func (adapter *DynamoDBAdapter) Delete(tableName string, key Attributes, conditionExpression *string) error {
	deleteInput := &dynamodb.DeleteItemInput{
		Key:                 key,
		TableName:           &tableName,
		ConditionExpression: conditionExpression,
	}
	_, err := Client.DeleteItem(deleteInput)
	if err != nil {
		error := errors.Errorf("Got error calling Delete: %v", err)
		return error
	}
	return nil
}

func (adapter *DynamoDBAdapter) Scan(tableName string, expr expression.Expression, outslice interface{}) error {

	if reflect.ValueOf(outslice).Kind() != reflect.Ptr {
		err := errors.New("out argument must be a pointer")
		log.Fatalf(err.Error())
		return err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	result, err := Client.Scan(params)
	if err != nil {
		error := errors.Errorf("Got error calling Scan")
		return error
	}

	items := result.Items
	err = dynamodbattribute.UnmarshalListOfMaps(items, outslice)
	if err != nil {
		error := errors.Wrap(err, "Got error calling Scan")
		return error
	}

	return err
}
