package messenger

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// MessageWriter defines an implementation to send messages
// This MessageWriter should be available for the connector implementations
type MessageWriter interface {
	// WriteState writes state information
	WriteState(stateData *protocol.StateData) error
	// WriteLog writes a log message
	WriteLog(logLevel protocol.LogLevel, logMessage string) error
}

type messageWriter struct {
	writer io.Writer
}

// NewMessageWriter creates a new instance of a MessageWriter
func NewMessageWriter(writer io.Writer) MessageWriter {
	return &messageWriter{writer}
}

// WriteState writes state information
func (mw *messageWriter) WriteState(stateData *protocol.StateData) error {
	message, err := protocol.NewStateMessage(&protocol.State{
		Data: stateData,
	})
	if err != nil {
		return err
	}

	return write(mw.writer, &message)
}

// WriteLog writes a log message
func (mw *messageWriter) WriteLog(logLevel protocol.LogLevel, logMessage string) error {
	message, err := protocol.NewLogMessage(&protocol.Log{
		Level:   logLevel,
		Message: logMessage,
	})
	if err != nil {
		return err
	}

	return write(mw.writer, &message)
}
