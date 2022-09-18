package protocol

import (
	"io"
)

// LogWriter is exported for documentation purposes - only use this through LogTracker or MessageTracker
// to ensure thread-safe behavior with the writer
type LogWriter func(level LogLevel, s string) error

// NewLogWriter returns the function that implements LogWriter func type
func NewLogWriter(w io.Writer) LogWriter {
	return func(lvl LogLevel, s string) error {
		return Write(w, &Message{
			Type: MsgTypeLog,
			logMessage: &logMessage{
				Level:   lvl,
				Message: s,
			},
		})
	}
}
