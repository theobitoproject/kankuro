package destination

import (
	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Destination is the only interface you need to define to create your destination!
type Destination interface {
	// Spec returns the schema which described how the destination connector can be configured
	Spec(
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
	) (*protocol.ConnectorSpecification, error)
	// Check verifies that, given a configuration, data can be accessed properly
	Check(
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
	) error
	// Write takes the data from the record channel
	// and stores it in the destination
	// Note: all channels from hub needs to be closed
	Write(
		cc *protocol.ConfiguredCatalog,
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
		hub messenger.ChannelHub,
	)
}
