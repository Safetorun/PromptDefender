package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserHandler_Handle(t *testing.T) {

	var called = false

	onCalled := func() {
		called = true
	}

	handler := &CreateUserHandler{
		userInstance: MockUserRepository{
			onCreateUserCalled: &onCalled,
		},
		apiKey: "testApiKey",
	}

	user := User{UserId: new(string)}
	*user.UserId = "testID"

	returnedUser, err := handler.Handle(user)

	assert.Nil(t, err)
	assert.Equal(t, user, *returnedUser)
	assert.True(t, called, "CreateUser should have been called")
}
