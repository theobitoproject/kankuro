package source

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/tools"
)

// NewSafeSourceRunner returns an instance of SourceRunner
// but it wraps the writer into a safe writer instance
// to properly write messages in safe manner
func NewSafeSourceRunner(
	src Source,
	w io.Writer,
	args []string,
) SourceRunner {
	sw := tools.NewSafeWriter(w)

	t := tools.NewTimer()

	mw := messenger.NewMessageWriter(sw)

	pmw := messenger.NewPrivateMessageWriter(sw, t)

	cp := messenger.NewConfigParser(args)

	hub := messenger.NewChannelHub(
		messenger.NewRecordChannel(),
		messenger.NewErrorChannel(),
	)

	return NewSourceRunner(src, mw, pmw, cp, hub)
}
