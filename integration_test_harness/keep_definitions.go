package integration_test_harness

import "context"

func RequestToKeep(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, RequestKey, &KeepRequest{}), nil
}
