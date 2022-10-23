package messenger

// ChannelHub defines a single object that holds all channels to share data
type ChannelHub interface {
	// GetRecordChannel returns the record channel
	GetRecordChannel() RecordChannel
	// GetErrorChannel returns the error channel
	GetErrorChannel() ErrorChannel
	// GetDoneChannel returns the done channel
	GetDoneChannel() DoneChannel
}

type channelHub struct {
	recordChannel RecordChannel
	errorChannel  ErrorChannel
	doneChannel   DoneChannel
}

// NewChannelHub creates a instance of ChannelHub
func NewChannelHub(
	recordChannel RecordChannel,
	errorChannel ErrorChannel,
	doneChannel DoneChannel,
) ChannelHub {
	return &channelHub{
		recordChannel,
		errorChannel,
		doneChannel,
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

// GetDoneChannel returns the done channel
func (ch *channelHub) GetDoneChannel() DoneChannel {
	return ch.doneChannel
}
