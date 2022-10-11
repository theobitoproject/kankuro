package messenger

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
	"github.com/theobitoproject/kankuro/tools"
)

// Messenger defines an implementation to send messages
// This Messenger should be available for the connector implementations
type Messenger interface {
	// WriteRecord writes a record
	WriteRecord(recordData *protocol.RecordData, stream string, namespace string) error
	// WriteState writes state information
	WriteState(stateData *protocol.StateData) error
	// WriteLog writes a log message
	WriteLog(logLevel protocol.LogLevel, logMessage string) error
}

type messenger struct {
	writer io.Writer
	timer  tools.Timer
}

// NewMessenger creates a new instance of a Messenger
func NewMessenger(writer io.Writer, timer tools.Timer) Messenger {
	return &messenger{writer, timer}
}

// WriteRecord writes a record
func (m *messenger) WriteRecord(recordData *protocol.RecordData, stream string, namespace string) error {
	message, err := protocol.NewRecordMessage(&protocol.Record{
		EmittedAt: m.timer.Now().UnixMilli(),
		Data:      recordData,
		Namespace: namespace,
		Stream:    stream,
	})
	if err != nil {
		return err
	}

	return write(m.writer, &message)
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
