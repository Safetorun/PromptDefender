package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/internal/base_aws"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
)

var (
	userTracer = otel.Tracer("users")
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(fmt.Sprintf("Received a request %+v", request))

	ctx, span := userTracer.Start(ctx, "lambda_users")
	defer span.End()

	if request.HTTPMethod == "POST" {
		fmt.Println("Received a POST request")
		handler := NewCreateUserHandler(request.RequestContext.Identity.APIKeyID)
		response, err := base_aws.BaseHandler[User, User](request, handler)
		response.StatusCode = 201
		return response, err

	} else if request.HTTPMethod == "GET" {
		fmt.Println("Received a GET request")
		if request.PathParameters["userId"] != "" {
			fmt.Println("Received a GET request with id: ", request.PathParameters["userId"])
			handler := NewRetrieverHandlerSingle(request.RequestContext.Identity.APIKeyID)
			response := handler.Handle(request.PathParameters["userId"])
			return response, nil
		} else {
			fmt.Println("Received a GET request with no id")
			userHandler := NewRetrieveHandler(request.RequestContext.Identity.APIKeyID)
			return userHandler.Handle(), nil
		}
	} else if request.HTTPMethod == "DELETE" {
		fmt.Println("Received a DELETE request")
		handler := NewDeleteUserHandler(request.RequestContext.Identity.APIKeyID)

		if request.PathParameters["userId"] == "" {
			return events.APIGatewayProxyResponse{StatusCode: 400}, errors.New(fmt.Sprintf("userId is required and was not found in %v", request.PathParameters))
		}
		requestLog := base_aws.ToRequestLog(request)
		base_aws.LogSummaryMessage(requestLog)
		return handler.Handle(request.PathParameters["userId"]), nil
	} else {
		fmt.Println("Received a request with an unsupported method")
		return events.APIGatewayProxyResponse{StatusCode: 400}, errors.New("unsupported method")
	}
}

func main() {

	ctx := context.Background()

	tp, err := xrayconfig.NewTracerProvider(ctx)
	if err != nil {
		fmt.Printf("error creating tracer provider: %v", err)
	}

	defer func(ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			fmt.Printf("error shutting down tracer provider: %v", err)
		}
	}(ctx)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})

	lambda.Start(otellambda.InstrumentHandler(Handler, xrayconfig.WithRecommendedOptions(tp)...))

}
