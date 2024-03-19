package models

import (
	
)

type SucecssResponse struct {
	Code    int    `json:"code,omitempty"`    // HTTP status code
	Message string `json:"message"`           // Human-readable message
	Data		interface{} `json:"data,omitempty"` // Optional data
}



