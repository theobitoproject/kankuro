package messenger

import (
	"encoding/json"
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Write emits data outbound from source to airbyte workers
func write(w io.Writer, m *protocol.Message) error {
	return json.NewEncoder(w).Encode(m)
}
