package protocol

import (
	"encoding/json"
)

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Auth struct {
	Auth string `json:"auth"`
}

type Payload struct {
	Type string          `json:"type"`
	Body json.RawMessage `json:"body"`
}
