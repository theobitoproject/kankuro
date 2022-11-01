package destination

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/tools"
)

// NewSafeDestinationRunner returns an instance of DestinationRunner
// reducing the complexity of creating it for clients
func NewSafeDestinationRunner(
	dst Destination,
	w io.Writer,
	r io.Reader,
	args []string,
) *DestinationRunner {
	sw := tools.NewSafeWriter(w)

	t := tools.NewTimer()

	mw := messenger.NewMessageWriter(sw)
	pmw := messenger.NewPrivateMessageWriter(sw, t)

	mr := messenger.NewMessageReader(r)

	cp := messenger.NewConfigParser(args)

	hub := messenger.NewChannelHub(
		messenger.NewRecordChannel(),
		messenger.NewErrorChannel(),
	)

	return NewDestinationRunner(dst, mw, pmw, mr, cp, hub)
}
