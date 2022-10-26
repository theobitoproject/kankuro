package messenger

// ClosingChannel defines a channel to indicate that the process has finished
type ClosingChannel chan bool

// NewClosingChannel creates an instance of ClosingChannel
func NewClosingChannel() ClosingChannel {
	return make(ClosingChannel)
}
