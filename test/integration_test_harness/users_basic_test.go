package integration_test_harness

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicUserTest(t *testing.T) {
	ctx, err := RetrieveSuspiciousUser(context.Background(), "test_user")
	assert.Nilf(t, err, "error should be nil but got %s", err)
	assert.Equal(t, 404, ctx.Value(ResponseStatus))
}
