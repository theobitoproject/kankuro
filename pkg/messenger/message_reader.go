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
			msg := scanner.Text()

			var rec *protocol.Record
			err = json.Unmarshal([]byte(msg), rec)
			if err != nil {
				hub.GetErrorChannel() <- err
			}

			hub.GetRecordChannel() <- rec
		}
	}
	if err := scanner.Err(); err != nil {
		hub.GetErrorChannel() <- err
	}

	close(hub.GetRecordChannel())
}
