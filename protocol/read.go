package protocol

import (
	"encoding/json"
	"io"
)

func Read(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}
