package main

import (
	"context"
	"net/http"
)

type SuspiciousUser interface {
	CheckSuspiciousUser(ctx context.Context, userId string) (*bool, error)
}

type RemoteSuspiciousUser struct {
	url    string
	apiKey string
}

func NewRemote(apiKey string) RemoteSuspiciousUser {
	return RemoteSuspiciousUser{
		url:    "https://prompt.safetorun.com",
		apiKey: apiKey,
	}
}

func (r *RemoteSuspiciousUser) CreateClient() (*ClientWithResponses, error) {

	addApiKey := func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-api-key", r.apiKey)
			return nil
		})

		return nil
	}

	return NewClientWithResponses(r.url, addApiKey)
}

// CheckSuspiciousUser checks if a user is suspicious
func (r *RemoteSuspiciousUser) CheckSuspiciousUser(ctx context.Context, userId string) (*bool, error) {
	client, err := r.CreateClient()
	if err != nil {
		return nil, err
	}

	response, err := client.GetUserWithResponse(ctx, userId)

	if err != nil {
		return nil, err
	}

	isUserSuspicious := response.StatusCode() == 404
	return &isUserSuspicious, nil
}
