package main

import (
	"encoding/json"
	"github.com/safetorun/PromptDefender/user_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveUserHandler_Handle(t *testing.T) {
	getUsersResponse := []user_repository.UserCore{
		{UserOrSessionId: "mockedID1"},
		{UserOrSessionId: "mockedID2"},
	}

	expected := []User{
		{UserId: &getUsersResponse[0].UserOrSessionId},
		{UserId: &getUsersResponse[1].UserOrSessionId},
	}
	handler := &RetrieveUserHandler{
		userInstance: &MockUserRepository{
			getUsersResponse: getUsersResponse,
		},
		apikeyId: "testApiKey",
		logger:   nil,
	}

	response := handler.Handle()

	var users []User
	err := json.Unmarshal([]byte(response.Body), &users)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Len(t, users, 2)
	assert.ElementsMatch(t, users, expected)
}
