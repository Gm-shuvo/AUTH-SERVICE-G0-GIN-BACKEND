package models

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`    // HTTP status code
	Type    string `json:"type,omitempty"`    // A short error type or code
	Message string `json:"message"`           // Human-readable message
	Details []string `json:"details,omitempty"` // Optional details
}

// Error makes ErrorResponse satisfy the error interface.
func (e *ErrorResponse) Error() string {
	// You can customize the message format as needed
	return fmt.Sprintf("[%d %s] %s", e.Code, e.Type, e.Message)
}
