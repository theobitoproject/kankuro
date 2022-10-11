package protocol

// Record defines a record as per airbyte - a "data point"
type Record struct {
	EmittedAt int64       `json:"emitted_at"`
	Namespace string      `json:"namespace"`
	Data      *RecordData `json:"data"`
	Stream    string      `json:"stream"`
}

// RecordData defines the data of the record
type RecordData map[string]interface{}
