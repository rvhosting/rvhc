package protocol

import (
	"encoding/json"
	"io"
)

func Write(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}
