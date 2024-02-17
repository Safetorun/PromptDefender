package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/internal/base_aws"
	"github.com/safetorun/PromptDefender/tracer"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"log"
)

var (
	user_tracer = otel.Tracer("users")
	meter       = otel.Meter("users")
)

type TracerStruct struct {
	context context.Context
	logger  *log.Logger
}

func NewTracer(context context.Context) *TracerStruct {
	return &TracerStruct{
		context: context,
		logger:  log.Default(),
	}
}

func (t *TracerStruct) TraceDecorator(fn tracer.GenericFuncType, functionName string) tracer.GenericFuncType {
	return func(args ...interface{}) (interface{}, error) {
		t.logger.Printf("Tracing function call, args: %s\n", functionName)
		_, tr := user_tracer.Start(t.context, functionName)
		defer tr.End()

		return fn(args...)
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received a request")
	ctx, span := user_tracer.Start(ctx, "lambda_users")
	defer span.End()

	if request.HTTPMethod == "POST" {
		fmt.Println("Received a POST request")
		handler := CreateUserHandler{user_repository_ddb.New(), request.RequestContext.Identity.APIKeyID}
		return base_aws.BaseHandler[User, User](request, &handler)

	} else if request.HTTPMethod == "GET" {
		fmt.Println("Received a GET request")
		if request.PathParameters["id"] != "" {
			fmt.Println("Received a GET request with id: ", request.PathParameters["id"])
			handler := RetrieveUserHandlerSingle{user_repository_ddb.New()}
			return handler.Handle(request.PathParameters["id"]), nil
		} else {
			userHandler := RetrieveUserHandler{user_repository_ddb.New(), request.RequestContext.Identity.APIKeyID}
			return userHandler.Handle(), nil
		}
	} else if request.HTTPMethod == "DELETE" {
		fmt.Println("Received a DELETE request")
		handler := DeleteUserHandler{user_repository_ddb.New()}
		return handler.Handle(request.PathParameters["id"]), nil
	} else {
		fmt.Println("Received a request with an unsupported method")
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
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
