package protocol

// Log defines messages that be used for debugging
type Log struct {
	Level   LogLevel `json:"level"`
	Message string   `json:"message"`
}
