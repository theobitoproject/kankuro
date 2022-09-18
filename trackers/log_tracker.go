package trackers

import "github.com/theobitoproject/kankuro/protocol"

// LogTracker is a single struct which holds a tracker which can be used for logs
type LogTracker struct {
	// Log logs out to airbyte
	Log protocol.LogWriter
}
