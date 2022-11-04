package protocol

import "encoding/json"

// RecordData defines the data of a "data point"
type RecordData map[string]interface{}

// String returns a string representation for a RecordData instance
func (rd *RecordData) String() string {
	bt, err := json.Marshal(rd)
	if err != nil {
		// TODO: is this right?
		return ""
	}

	return string(bt)
}
