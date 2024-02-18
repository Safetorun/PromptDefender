package integration_test_harness

import (
	"context"
	"errors"
	"fmt"
)

const UsersKey = "users"
const ResponseStatus = "ResponseStatus"

func RetrieveSuspiciousUser(ctx context.Context, userId string) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	response, err := gClient.GetUserWithResponse(ctx, userId)
	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}
	return context.WithValue(ctx, ResponseStatus, response.StatusCode()), nil
}

func Validate404Error(ctx context.Context) error {
	if ctx.Value(ResponseStatus) != 404 {
		return errors.New("response status is not 404 it is " + fmt.Sprint(ctx.Value(ResponseStatus)))
	}

	return nil
}

func CreateSuspiciousUser(ctx context.Context, userId string) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	user := User{UserId: &userId}

	response, err := gClient.AddUserWithResponse(context.Background(), user)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 201 {
		return ctx, errors.New(fmt.Sprintf("error processing request response is: %s", response.Status()))
	}

	return ctx, nil
}

func RetrieveSuspiciousUsers(ctx context.Context) (context.Context, error) {
	gClient, err := CreateClient()
	if err != nil {
		return nil, err
	}

	response, err := gClient.ListUsersWithResponse(ctx)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when listing users", err)
	}

	if response.StatusCode() != 200 {
		return ctx, errors.New(fmt.Sprintf("error processing request response (%s) is %s", response.Status(), string(response.Body)))
	}

	return context.WithValue(ctx, UsersKey, *response.JSON200), nil
}

func DeleteSuspiciousUser(ctx context.Context, userId string) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	response, err := gClient.RemoveUserWithResponse(context.Background(), userId)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 204 {
		return ctx, errors.New(fmt.Sprintf("error processing request response (%s) is %s", response.Status(), string(response.Body)))
	}

	return ctx, nil
}

func ValidateUserIdContains(ctx context.Context, userId string) error {
	users := ctx.Value(UsersKey).([]User)

	for _, user := range users {
		if *user.UserId == userId {
			return nil
		}
	}

	return errors.New("user id not found")
}

func ValidateUserNotInList(ctx context.Context, userId string) error {
	users := ctx.Value(UsersKey).([]User)

	for _, user := range users {
		if *user.UserId == userId {
			return errors.New("user id found")
		}
	}

	return nil
}
