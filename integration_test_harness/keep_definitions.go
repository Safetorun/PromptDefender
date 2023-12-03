package integration_test_harness

import (
	"context"
	"errors"
	"fmt"
)

func RequestToKeep(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, RequestKey, &KeepRequest{}), nil
}

func SendRequestKeep(ctx context.Context) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	if ctx.Value(RequestKey) == nil {
		return ctx, errors.New("request is nil")
	}

	request, ok := ctx.Value(RequestKey).(*MoatRequest)

	if ok == false {
		return ctx, errors.New("request is not castable to MoatRequest")
	}

	response, err := gClient.BuildShieldWithResponse(context.Background(), *request)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 200 {
		return ctx, errors.New("error processing request")
	}

	if response.JSON200 == nil {
		return ctx, errors.New("response is nil")
	}

	return context.WithValue(ctx, ResponseKey, response.JSON200), nil

}
