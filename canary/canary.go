// Package canary provides utilities for canary detection.
package canary

import (
	"errors"
)

// Mode represents the type of canary checking to be performed.
type Mode int

// Enumeration of available Modes.
const (
	None         Mode = iota // None No canary checking is performed.
	Basic                    // Basic mode performs an exact match check.
	Intermediate             // Intermediate mode is not implemented yet.
	Advanced                 // Advanced mode is not implemented yet.
)

// Result holds the result of a canary check.
type Result struct {
	Detected bool // Detected indicates if the canary was detected.
}

// Checker encapsulates the Mode and Canary string for performing canary checks.
type Checker struct {
	Mode   Mode   // Mode represents the canary checking mode.
	Canary string // Canary holds the canary string to be checked.
}

// New creates a new instance of Checker.
// It takes a Mode and a canary string as parameters.
func New(mode Mode, canary string) *Checker {
	return &Checker{Mode: mode, Canary: canary}
}

// Check performs a canary check based on the mode.
// It returns a Result and an error.
// The only implemented Mode as of now is Basic, which performs an exact match check.
func (c *Checker) Check(output string) (Result, error) {
	switch c.Mode {
	case None:
		return Result{Detected: false}, nil
	case Basic:
		return Result{Detected: output == c.Canary}, nil
	default:
		return Result{}, errors.New("mode not implemented")
	}
}
