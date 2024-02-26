package base_aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/request_log"
	"log"
	"time"
)

func RequestCallbackComplete(responseBody string, request events.APIGatewayProxyRequest, startTime time.Time) {
	queueMessage := ToRequestLog(request)
	queueMessage.Response = responseBody
	queueMessage.Time = int(time.Since(startTime).Milliseconds())
	LogSummaryMessage(queueMessage)
}

func LogSummaryMessage(message request_log.QueueMessage) {

	jsonMessage, err := json.Marshal(message)

	if err != nil {
		log.Fatalf("JSON marshaling error: %v", err)
	}

	log.Default().Println("Summary message ---- ", string(jsonMessage))
}

func ToRequestLog(request events.APIGatewayProxyRequest) request_log.QueueMessage {
	return request_log.QueueMessage{
		UserId:          request.RequestContext.Identity.APIKeyID,
		Domain:          request.RequestContext.DomainName,
		Headers:         request.Headers,
		Method:          request.RequestContext.HTTPMethod,
		QueryParams:     request.QueryStringParameters,
		HttpMethod:      request.RequestContext.HTTPMethod,
		StartedDateTime: request.RequestContext.RequestTime,
		HttpResponseHeaders: map[string]string{
			"content-type": "application/json",
		},
		HttpResponse: 200,
		Endpoint:     request.Path,
		Request:      request.Body,
	}
}
