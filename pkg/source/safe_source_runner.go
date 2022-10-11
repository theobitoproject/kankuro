package source

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/tools"
)

// NewSafeSourceRunner returns an instance of SourceRunner
// but it wraps the writer into a safe writer instance
// to properly write messages in safe manner
func NewSafeSourceRunner(src Source, writer io.Writer, args []string) SourceRunner {
	safeWriter := newSafeWriter(writer)

	timer := tools.NewTimer()

	msgr := messenger.NewMessenger(safeWriter, timer)

	prvtMsgr := messenger.NewPrivateMessenger(safeWriter)

	configParser := messenger.NewConfigParser(args)

	return NewSourceRunner(src, msgr, prvtMsgr, configParser)
}
