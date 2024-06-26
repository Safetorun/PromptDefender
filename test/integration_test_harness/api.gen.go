// Package integration_test_harness provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package integration_test_harness

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// KeepRequest defines model for KeepRequest.
type KeepRequest struct {
	// Prompt The base prompt you want to build a keep for
	Prompt string `json:"prompt"`

	// RandomiseXmlTag Whether to randomise the XML tag that is used to escape user input in your prompt.
	RandomiseXmlTag *bool `json:"randomise_xml_tag,omitempty"`
}

// KeepResponse defines model for KeepResponse.
type KeepResponse struct {
	// Canary The canary string that is used to detect prompt injection.
	Canary *string `json:"canary,omitempty"`

	// ShieldedPrompt The shielded prompt.
	ShieldedPrompt string `json:"shielded_prompt"`

	// XmlTag The XML tag that is used to escape user input in your prompt.
	XmlTag string `json:"xml_tag"`
}

// User defines model for User.
type User struct {
	// UserId The user ID of the suspicious user.
	UserId *string `json:"userId,omitempty"`
}

// WallRequest defines model for WallRequest.
type WallRequest struct {
	// CheckBadwords Whether to scan for badwords in the prompt.
	CheckBadwords *bool `json:"check_badwords,omitempty"`

	// FastCheck Whether to perform a fast check on the prompt instead of a full check.
	FastCheck *bool `json:"fast_check,omitempty"`

	// Prompt The text prompt to be verified.
	Prompt string `json:"prompt"`

	// ScanPii Whether to scan for PII in the prompt.
	ScanPii *bool `json:"scan_pii,omitempty"`

	// SessionId The session ID of the user who is submitting the prompt. This is used to flag suspicious sessions
	SessionId *string `json:"session_id,omitempty"`

	// UserId The user ID of the user who is submitting the prompt. This is used to flag suspicious users
	UserId *string `json:"user_id,omitempty"`

	// XmlTag The XML tag that is used to escape user input in your prompt (this may have been generated with keep).
	XmlTag *string `json:"xml_tag,omitempty"`
}

// WallResponse defines model for WallResponse.
type WallResponse struct {
	// ContainsPii Whether the prompt contains PII.
	ContainsPii *bool `json:"contains_pii,omitempty"`

	// ModifiedPrompt The prompt, modified in some way - e.g. by summarising it as per the configuration specified in the request.
	ModifiedPrompt *string `json:"modified_prompt,omitempty"`

	// PotentialJailbreak Whether the prompt contains a potential jailbreak.
	PotentialJailbreak *bool `json:"potential_jailbreak,omitempty"`

	// PotentialXmlEscaping Whether the prompt contains potential XML escaping.
	PotentialXmlEscaping *bool `json:"potential_xml_escaping,omitempty"`

	// SuspiciousSession Whether the session is suspicious.
	SuspiciousSession *bool `json:"suspicious_session,omitempty"`

	// SuspiciousUser Whether the user is suspicious.
	SuspiciousUser *bool `json:"suspicious_user,omitempty"`
}

// BuildKeepJSONRequestBody defines body for BuildKeep for application/json ContentType.
type BuildKeepJSONRequestBody = KeepRequest

// AddUserJSONRequestBody defines body for AddUser for application/json ContentType.
type AddUserJSONRequestBody = User

// BuildWallJSONRequestBody defines body for BuildWall for application/json ContentType.
type BuildWallJSONRequestBody = WallRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// BuildKeepWithBody request with any body
	BuildKeepWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BuildKeep(ctx context.Context, body BuildKeepJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListUsers request
	ListUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddUserWithBody request with any body
	AddUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddUser(ctx context.Context, body AddUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RemoveUser request
	RemoveUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUser request
	GetUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BuildWallWithBody request with any body
	BuildWallWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BuildWall(ctx context.Context, body BuildWallJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) BuildKeepWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBuildKeepRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BuildKeep(ctx context.Context, body BuildKeepJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBuildKeepRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddUser(ctx context.Context, body AddUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RemoveUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRemoveUserRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BuildWallWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBuildWallRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BuildWall(ctx context.Context, body BuildWallJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBuildWallRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewBuildKeepRequest calls the generic BuildKeep builder with application/json body
func NewBuildKeepRequest(server string, body BuildKeepJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBuildKeepRequestWithBody(server, "application/json", bodyReader)
}

// NewBuildKeepRequestWithBody generates requests for BuildKeep with any type of body
func NewBuildKeepRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/keep")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListUsersRequest generates requests for ListUsers
func NewListUsersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddUserRequest calls the generic AddUser builder with application/json body
func NewAddUserRequest(server string, body AddUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddUserRequestWithBody(server, "application/json", bodyReader)
}

// NewAddUserRequestWithBody generates requests for AddUser with any type of body
func NewAddUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRemoveUserRequest generates requests for RemoveUser
func NewRemoveUserRequest(server string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUserRequest generates requests for GetUser
func NewGetUserRequest(server string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewBuildWallRequest calls the generic BuildWall builder with application/json body
func NewBuildWallRequest(server string, body BuildWallJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBuildWallRequestWithBody(server, "application/json", bodyReader)
}

// NewBuildWallRequestWithBody generates requests for BuildWall with any type of body
func NewBuildWallRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/wall")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// BuildKeepWithBodyWithResponse request with any body
	BuildKeepWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BuildKeepResponse, error)

	BuildKeepWithResponse(ctx context.Context, body BuildKeepJSONRequestBody, reqEditors ...RequestEditorFn) (*BuildKeepResponse, error)

	// ListUsersWithResponse request
	ListUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUsersResponse, error)

	// AddUserWithBodyWithResponse request with any body
	AddUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddUserResponse, error)

	AddUserWithResponse(ctx context.Context, body AddUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AddUserResponse, error)

	// RemoveUserWithResponse request
	RemoveUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*RemoveUserResponse, error)

	// GetUserWithResponse request
	GetUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*GetUserResponse, error)

	// BuildWallWithBodyWithResponse request with any body
	BuildWallWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BuildWallResponse, error)

	BuildWallWithResponse(ctx context.Context, body BuildWallJSONRequestBody, reqEditors ...RequestEditorFn) (*BuildWallResponse, error)
}

type BuildKeepResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *KeepResponse
}

// Status returns HTTPResponse.Status
func (r BuildKeepResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BuildKeepResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]User
}

// Status returns HTTPResponse.Status
func (r ListUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r AddUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RemoveUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r RemoveUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RemoveUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r GetUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BuildWallResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WallResponse
}

// Status returns HTTPResponse.Status
func (r BuildWallResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BuildWallResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BuildKeepWithBodyWithResponse request with arbitrary body returning *BuildKeepResponse
func (c *ClientWithResponses) BuildKeepWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BuildKeepResponse, error) {
	rsp, err := c.BuildKeepWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBuildKeepResponse(rsp)
}

func (c *ClientWithResponses) BuildKeepWithResponse(ctx context.Context, body BuildKeepJSONRequestBody, reqEditors ...RequestEditorFn) (*BuildKeepResponse, error) {
	rsp, err := c.BuildKeep(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBuildKeepResponse(rsp)
}

// ListUsersWithResponse request returning *ListUsersResponse
func (c *ClientWithResponses) ListUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListUsersResponse, error) {
	rsp, err := c.ListUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListUsersResponse(rsp)
}

// AddUserWithBodyWithResponse request with arbitrary body returning *AddUserResponse
func (c *ClientWithResponses) AddUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddUserResponse, error) {
	rsp, err := c.AddUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddUserResponse(rsp)
}

func (c *ClientWithResponses) AddUserWithResponse(ctx context.Context, body AddUserJSONRequestBody, reqEditors ...RequestEditorFn) (*AddUserResponse, error) {
	rsp, err := c.AddUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddUserResponse(rsp)
}

// RemoveUserWithResponse request returning *RemoveUserResponse
func (c *ClientWithResponses) RemoveUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*RemoveUserResponse, error) {
	rsp, err := c.RemoveUser(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRemoveUserResponse(rsp)
}

// GetUserWithResponse request returning *GetUserResponse
func (c *ClientWithResponses) GetUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*GetUserResponse, error) {
	rsp, err := c.GetUser(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserResponse(rsp)
}

// BuildWallWithBodyWithResponse request with arbitrary body returning *BuildWallResponse
func (c *ClientWithResponses) BuildWallWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BuildWallResponse, error) {
	rsp, err := c.BuildWallWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBuildWallResponse(rsp)
}

func (c *ClientWithResponses) BuildWallWithResponse(ctx context.Context, body BuildWallJSONRequestBody, reqEditors ...RequestEditorFn) (*BuildWallResponse, error) {
	rsp, err := c.BuildWall(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBuildWallResponse(rsp)
}

// ParseBuildKeepResponse parses an HTTP response from a BuildKeepWithResponse call
func ParseBuildKeepResponse(rsp *http.Response) (*BuildKeepResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BuildKeepResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest KeepResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseListUsersResponse parses an HTTP response from a ListUsersWithResponse call
func ParseListUsersResponse(rsp *http.Response) (*ListUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAddUserResponse parses an HTTP response from a AddUserWithResponse call
func ParseAddUserResponse(rsp *http.Response) (*AddUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AddUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseRemoveUserResponse parses an HTTP response from a RemoveUserWithResponse call
func ParseRemoveUserResponse(rsp *http.Response) (*RemoveUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RemoveUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetUserResponse parses an HTTP response from a GetUserWithResponse call
func ParseGetUserResponse(rsp *http.Response) (*GetUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseBuildWallResponse parses an HTTP response from a BuildWallWithResponse call
func ParseBuildWallResponse(rsp *http.Response) (*BuildWallResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BuildWallResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WallResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
