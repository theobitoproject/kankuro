package protocol

const (
	// AirbyteRaw defines the prefix that should be used for the name
	// of a stream (file, table, collection, etc) when raw data is stored
	// and no transformation is used
	AirbyteRaw = "_airbyte_raw_"
	// AirbyteAbId is the name of the column/field that represents
	// the uuid value assigned by connectors to each row of the data
	// written in the destination
	AirbyteAbId = "_airbyte_ab_id"
	// AirbyteEmittedAt is the of the column/field that represents
	// the time at which the record was emitted
	// and recorded by destination connector
	AirbyteEmittedAt = "_airbyte_emitted_at"
	// AirbyteData is the of the column/field that represents
	// the entire data for a single point/row/document
	AirbyteData = "_airbyte_data"
)

// RawRecord is the the raw format for a single data point
// without any normalization (Raw data (JSON) option)
type RawRecord struct {
	ID        string      `json:"_airbyte_ab_id"`
	EmittedAt int64       `json:"_airbyte_emitted_at"`
	Data      *RecordData `json:"_airbyte_data"`
}
