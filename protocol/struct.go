package protocol

import (
	"encoding/json"
)

type Auth struct {
	Auth string `json:"auth"`
}

type Payload struct {
	Type string          `json:"type"`
	Body json.RawMessage `json:"body"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}
