package integration_test_harness

import (
	"context"
	"fmt"
	"testing"
)

func TestBasicUserTest(t *testing.T) {
	client, err := CreateClient()

	if err != nil {
		t.Errorf("error: %v", err)
	}

	re, err := client.ListUsers(context.Background())

	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Printf(fmt.Sprintf("%+v", *re))

	if re.StatusCode != 200 {
		t.Errorf("error: %v", re.Body)
	}
}
