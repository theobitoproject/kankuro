package messenger

// DoneChannel defines a channel to indicate that the process has finished
type DoneChannel chan bool

// NewDoneChannel creates an instance of DoneChannel
func NewDoneChannel() DoneChannel {
	return make(DoneChannel)
}
