package protocol

import "io"

// StateWriter is exported for documentation purposes - only use this through MessageTracker
type StateWriter func(v interface{}) error

// NewStateWriter returns the function that implements StateWriter func type
func NewStateWriter(w io.Writer) StateWriter {
	return func(s interface{}) error {
		return Write(w, &Message{
			Type: MsgTypeState,
			state: &state{
				Data: s,
			},
		})
	}
}
