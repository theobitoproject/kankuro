package destination

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/tools"
)

// NewSafeSourceRunner returns an instance of SourceRunner
// but it wraps the writer into a safe writer instance
// to properly write messages in safe manner
func NewSafeDestinationRunner(dst Destination, writer io.Writer, args []string) DestinationRunner {
	safeWriter := newSafeWriter(writer)

	timer := tools.NewTimer()

	msgr := messenger.NewMessenger(safeWriter)
	prvtMsgr := messenger.NewPrivateMessenger(safeWriter, timer)

	configParser := messenger.NewConfigParser(args)

	return NewDestinationRunner(dst, msgr, prvtMsgr, configParser)
}
