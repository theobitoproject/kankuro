package trackers

import "github.com/theobitoproject/kankuro/protocol"

// MessageTracker is used to encap State tracking, Record tracking and Log tracking
// It's thread safe
type MessageTracker struct {
	// State will save an arbitrary JSON blob to airbyte state
	State protocol.StateWriter
	// Record will emit a record (data point) out to airbyte to sync with appropriate timestamps
	Record protocol.RecordWriter
	// Log logs out to airbyte
	Log protocol.LogWriter
}
