package protocol

import "github.com/google/uuid"

// Record defines a record as per airbyte - a "data point"
// to be sent inside a message
type Record struct {
	EmittedAt int64       `json:"emitted_at"`
	Namespace string      `json:"namespace"`
	Data      *RecordData `json:"data"`
	Stream    string      `json:"stream"`
}

// GetRawRecord returns a RawRecord instance from a record instance
func (r *Record) GetRawRecord() *RawRecord {
	return &RawRecord{
		ID:        uuid.New().String(),
		EmittedAt: r.EmittedAt,
		Data:      r.Data,
	}
}
