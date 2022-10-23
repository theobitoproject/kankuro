package messenger

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Messenger defines an implementation to send messages
// This Messenger should be available for the connector implementations
type Messenger interface {
	// WriteState writes state information
	WriteState(stateData *protocol.StateData) error
	// WriteLog writes a log message
	WriteLog(logLevel protocol.LogLevel, logMessage string) error
}

type messenger struct {
	writer io.Writer
}

// NewMessenger creates a new instance of a Messenger
func NewMessenger(writer io.Writer) Messenger {
	return &messenger{writer}
}

// WriteState writes state information
func (m *messenger) WriteState(stateData *protocol.StateData) error {
	message, err := protocol.NewStateMessage(&protocol.State{
		Data: stateData,
	})
	if err != nil {
		return err
	}

	return write(m.writer, &message)
}

// WriteLog writes a log message
func (m *messenger) WriteLog(logLevel protocol.LogLevel, logMessage string) error {
	message, err := protocol.NewLogMessage(&protocol.Log{
		Level:   logLevel,
		Message: logMessage,
	})
	if err != nil {
		return err
	}

	return write(m.writer, &message)
}
