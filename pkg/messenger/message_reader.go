package messenger

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// MessageReader defines an implementation to read messages
// This MessageReader should NOT be available for the connector implementations
type MessageReader interface {
	// Read will read the information from reader and
	// send it through the record channel
	Read(hub ChannelHub)
}

type messageReader struct {
	r io.Reader
}

// NewMessageReader creates a new instance of MessageReader
func NewMessageReader(r io.Reader) MessageReader {
	return &messageReader{r}
}

// Read will read the information from reader and
// send it through the record channel
func (mr *messageReader) Read(hub ChannelHub) {
	var err error

	scanner := bufio.NewScanner(mr.r)
	for scanner.Scan() {
		select {

		case <-hub.GetClosingChannel():
			close(hub.GetRecordChannel())
			return

		default:
			txt := scanner.Text()

			var msg protocol.AirbyteMessage
			err = json.Unmarshal([]byte(txt), &msg)
			if err != nil {
				hub.GetErrorChannel() <- err
				continue
			}

			if msg.Type != protocol.MessageTypeRecord {
				// ignore all messages but records
				continue
			}

			hub.GetRecordChannel() <- msg.Record
		}
	}
	if err := scanner.Err(); err != nil {
		hub.GetErrorChannel() <- err
	}

	close(hub.GetRecordChannel())
}
