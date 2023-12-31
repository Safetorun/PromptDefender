// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package main

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
	// ShieldedPrompt The shielded prompt.
	ShieldedPrompt string `json:"shielded_prompt"`

	// XmlTag The XML tag that is used to escape user input in your prompt.
	XmlTag string `json:"xml_tag"`
}

// MoatRequest defines model for MoatRequest.
type MoatRequest struct {
	// Prompt The text prompt to be verified.
	Prompt string `json:"prompt"`

	// ScanPii Whether to scan for PII in the prompt.
	ScanPii bool `json:"scan_pii"`

	// XmlTag The XML tag that is used to escape user input in your prompt (this may have been generated with keep).
	XmlTag *string `json:"xml_tag,omitempty"`
}

// MoatResponse defines model for MoatResponse.
type MoatResponse struct {
	// ContainsPii Whether the prompt contains PII.
	ContainsPii *bool `json:"contains_pii,omitempty"`

	// PotentialJailbreak Whether the prompt contains a potential jailbreak.
	PotentialJailbreak *bool `json:"potential_jailbreak,omitempty"`

	// PotentialXmlEscaping Whether the prompt contains potential XML escaping.
	PotentialXmlEscaping *bool `json:"potential_xml_escaping,omitempty"`
}

// WallRequest defines model for WallRequest.
type WallRequest struct {
	// Prompt The text prompt to be verified.
	Prompt string `json:"prompt"`
}

// WallResponse defines model for WallResponse.
type WallResponse struct {
	// InjectionScore The score indicating the likelihood of prompt injection.
	InjectionScore *float32 `json:"injection_score,omitempty"`
}

// BuildKeepJSONRequestBody defines body for BuildKeep for application/json ContentType.
type BuildKeepJSONRequestBody = KeepRequest

// BuildShieldJSONRequestBody defines body for BuildShield for application/json ContentType.
type BuildShieldJSONRequestBody = MoatRequest

// WallPromptJSONRequestBody defines body for WallPrompt for application/json ContentType.
type WallPromptJSONRequestBody = WallRequest
