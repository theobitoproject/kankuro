package protocol

import (
	"encoding/json"
	"io"
)

// Write emits data outbound from your src/destination to airbyte workers
func Write(w io.Writer, m *Message) error {
	return json.NewEncoder(w).Encode(m)
}
