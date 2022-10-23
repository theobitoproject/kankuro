package messenger

import "github.com/theobitoproject/kankuro/pkg/protocol"

// RecordChannel defines a channel to share records
type RecordChannel chan *protocol.Record

// NewRecordChannel creates an instance of RecordChannel
func NewRecordChannel() RecordChannel {
	return make(RecordChannel)
}
