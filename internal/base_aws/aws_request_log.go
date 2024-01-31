package base_aws

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/pkg/request_log"
)

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
	}
}
