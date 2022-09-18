package protocol

import (
	"io"
	"time"
)

// RecordWriter is exported for documentation purposes - only use this through MessageTracker
type RecordWriter func(v interface{}, streamName string, namespace string) error

// NewRecordWriter returns the function that implements RecordWriter func type
func NewRecordWriter(w io.Writer) RecordWriter {
	return func(s interface{}, stream string, namespace string) error {
		return Write(w, &Message{
			Type: MsgTypeRecord,
			record: &record{
				EmittedAt: time.Now().UnixMilli(),
				Data:      s,
				Namespace: namespace,
				Stream:    stream,
			},
		})
	}
}
