package protocol

import (
	"io"
	"time"
)

// Messenger defines an implementation to send messages
// This Messenger should be available for the connector implementations
type Messenger interface {
	// WriteRecord writes a record
	WriteRecord(data interface{}, stream string, namespace string) error
	// WriteState writes state information
	WriteState(data interface{}) error
	// WriteLog writes a log message
	WriteLog(level LogLevel, message string) error
}

type messenger struct {
	writer io.Writer
}

// NewMessenger creates a new instance of a Messenger
func NewMessenger(writer io.Writer) Messenger {
	return messenger{writer}
}

// WriteRecord writes a record
func (m messenger) WriteRecord(data interface{}, stream string, namespace string) error {
	return write(m.writer, &Message{
		Type: msgTypeRecord,
		Record: &record{
			EmittedAt: time.Now().UnixMilli(),
			Data:      data,
			Namespace: namespace,
			Stream:    stream,
		},
	})
}

// WriteState writes state information
func (m messenger) WriteState(data interface{}) error {
	return write(m.writer, &Message{
		Type: msgTypeState,
		State: &state{
			Data: data,
		},
	})
}

// WriteLog writes a log message
func (m messenger) WriteLog(level LogLevel, message string) error {
	return write(m.writer, &Message{
		Type: msgTypeLog,
		LogMessage: &logMessage{
			Level:   level,
			Message: message,
		},
	})
}
