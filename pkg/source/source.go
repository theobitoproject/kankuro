package source

import (
	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Source defines a connector to extract data
type Source interface {
	// Spec returns the schema which described how the source connector can be configured
	Spec(
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
	) (*protocol.ConnectorSpecification, error)
	// Check verifies that, given a configuration, data can be accessed properly
	Check(
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
	) error
	// Discover returns the schema which describes the structure of the data
	// that can be extracted from the source
	Discover(
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
	) (*protocol.Catalog, error)
	// Read fetches data from the source and
	// communicates all records to the record channel
	// Note: all channels from hub needs to be closed
	Read(
		cc *protocol.ConfiguredCatalog,
		mw messenger.MessageWriter,
		cp messenger.ConfigParser,
		hub messenger.ChannelHub,
	)
}
