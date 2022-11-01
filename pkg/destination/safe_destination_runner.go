package destination

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/tools"
)

// NewSafeSourceRunner returns an instance of SourceRunner
// but it wraps the writer into a safe writer instance
// to properly write messages in safe manner
func NewSafeDestinationRunner(
	dst Destination,
	w io.Writer,
	r io.Reader,
	args []string,
) DestinationRunner {
	sw := newSafeWriter(w)

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
