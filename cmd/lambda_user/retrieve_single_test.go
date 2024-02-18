package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveUserHandlerSingle_Handle(t *testing.T) {
	handler := &RetrieveUserHandlerSingle{
		userInstance: &MockUserRepository{},
	}

	id := "testID"
	response := handler.Handle(id)

	var user User
	err := json.Unmarshal([]byte(response.Body), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, id, *user.UserId)
}
