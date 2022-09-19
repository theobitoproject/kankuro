package source

import (
	"io"

	"github.com/theobitoproject/kankuro/protocol"
)

// NewSafeSourceRunner returns an instance of SourceRunner
// but it wraps the writer into a safe writer instance
// to properly write messages in safe manner
func NewSafeSourceRunner(src Source, writer io.Writer) SourceRunner {
	safeWriter := newSafeWriter(writer)

	messenger := protocol.NewMessenger(safeWriter)
	privateMessenger := protocol.NewPrivateMessenger(safeWriter)

	return NewSourceRunner(src, messenger, privateMessenger)
}
