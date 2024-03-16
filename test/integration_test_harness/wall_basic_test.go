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

	scanPii := false
	shieldResponse, err := client.BuildWallWithResponse(context.Background(), WallRequest{
		Prompt:  "Test",
		ScanPii: &scanPii,
		XmlTag:  nil,
	})

	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Printf(fmt.Sprintf("%+v", *shieldResponse.JSON200))
}
