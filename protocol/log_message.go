package protocol

type logMessage struct {
	Level   LogLevel `json:"level"`
	Message string   `json:"message"`
}
