package messenger

// ErrorChannel defines a channel to share errors
type ErrorChannel chan error

// NewErrorChannel creates an instance of ErrorChannel
func NewErrorChannel() ErrorChannel {
	return make(ErrorChannel)
}
