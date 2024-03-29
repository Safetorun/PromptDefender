package user_repository_ddb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/safetorun/PromptDefender/user_repository"
	"os"
)

const IndexKey = "ApiKeyId"

type UserRepositoryDdb struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func New() *UserRepositoryDdb {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	}))

	return &UserRepositoryDdb{
		db:        dynamodb.New(sess),
		tableName: os.Getenv("USERS_TABLE"),
	}
}

func (ddb *UserRepositoryDdb) GetUserByID(id string, apiKeyId string) (*user_repository.UserCore, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserOrSessionId": {
				S: aws.String(id),
			},
			"ApiKeyId": {
				S: aws.String(apiKeyId),
			},
		},
	}

	result, err := ddb.db.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, user_repository.ErrUserIDNotFound
	}

	user := user_repository.UserCore{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ddb *UserRepositoryDdb) GetUsers(apikeyId string) ([]user_repository.UserCore, error) {
	input := &dynamodb.QueryInput{
		TableName: aws.String(ddb.tableName),
		IndexName: aws.String("ApiKeyId-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			IndexKey: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(apikeyId),
					},
				},
			},
		},
	}

	result, err := ddb.db.Query(input)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result into a slice of UserCore objects
	users := []user_repository.UserCore{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ddb *UserRepositoryDdb) CreateUser(user user_repository.UserCore) error {
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(ddb.tableName),
		Item:      item,
	}

	_, err = ddb.db.PutItem(input)
	return err
}

func (ddb *UserRepositoryDdb) DeleteUser(userOrSessionId string, apiKeyId string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(ddb.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserOrSessionId": {
				S: aws.String(userOrSessionId),
			},
			"ApiKeyId": {
				S: aws.String(apiKeyId),
			},
		},
	}

	_, err := ddb.db.DeleteItem(input)
	return err
}
