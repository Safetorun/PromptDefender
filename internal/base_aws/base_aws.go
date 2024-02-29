package base_aws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"time"
)

type HandlerFunc[T any, V any] interface {
	Handle(T) (*V, error)
}

func BaseHandler[T any, V any](request events.APIGatewayProxyRequest, handlerFunc HandlerFunc[T, V]) (events.APIGatewayProxyResponse, error) {

	var input T

	startTime := time.Now()

	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	response, err := handlerFunc.Handle(input)

	if err != nil {
		fmt.Printf("error processing request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing request: %v", err)
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	jsonString := string(jsonBytes)

	RequestCallbackComplete(jsonString, request, startTime)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: jsonString}, nil

}
