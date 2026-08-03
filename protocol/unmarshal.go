package protocol

import (
	"encoding/json"
)

func (p Payload) UnmarshalTo(v any) error {
	return json.Unmarshal(p.Body, v)
}
