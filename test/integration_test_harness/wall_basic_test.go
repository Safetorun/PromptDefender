package integration_test_harness

import (
	"context"
	"fmt"
	"testing"
)

func TestBasicWallTest(t *testing.T) {
	client, err := CreateClient()

	if err != nil {
		t.Errorf("error: %v", err)
	}

	shieldResponse, err := client.BuildShieldWithResponse(context.Background(), WallRequest{
		Prompt:  "Test",
		ScanPii: false,
		XmlTag:  nil,
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Printf(fmt.Sprintf("%+v", *shieldResponse.JSON200))
}
