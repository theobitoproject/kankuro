package protocol

// record defines a record as per airbyte - a "data point"
type record struct {
	EmittedAt int64       `json:"emitted_at"`
	Namespace string      `json:"namespace"`
	Data      interface{} `json:"data"`
	Stream    string      `json:"stream"`
}
