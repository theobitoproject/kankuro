package messenger

// ChannelHub defines a single object that holds all channels to share data
type ChannelHub interface {
	// GetRecordChannel returns the record channel
	GetRecordChannel() RecordChannel
	// GetErrorChannel returns the error channel
	GetErrorChannel() ErrorChannel
}

type channelHub struct {
	recordChannel RecordChannel
	errorChannel  ErrorChannel
}

// NewChannelHub creates a instance of ChannelHub
func NewChannelHub(
	recordChannel RecordChannel,
	errorChannel ErrorChannel,
) ChannelHub {
	return &channelHub{
		recordChannel,
		errorChannel,
	}
}

// GetRecordChannel returns the record channel
func (ch *channelHub) GetRecordChannel() RecordChannel {
	return ch.recordChannel
}

// GetErrorChannel returns the error channel
func (ch *channelHub) GetErrorChannel() ErrorChannel {
	return ch.errorChannel
}
