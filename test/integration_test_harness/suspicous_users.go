package integration_test_harness

import (
	"context"
	"errors"
	"fmt"
)

func CreateSuspiciousUser(ctx context.Context) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	if ctx.Value(RequestKey) == nil {
		return ctx, errors.New("request is nil")
	}

	request, ok := ctx.Value(RequestKey).(*User)

	if ok == false {
		return ctx, errors.New("request is not castable to MoatRequest")
	}

	response, err := gClient.AddUserWithResponse(context.Background(), *request)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 201 {
		return ctx, errors.New("error processing request")
	}

	return ctx, nil
}
