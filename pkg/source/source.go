package source

import (
	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Source defines a connector to extract data
type Source interface {
	// Spec returns the schema which described how the source connector can be configured
	Spec(
		msgr messenger.Messenger,
		cfgPsr messenger.ConfigParser,
	) (*protocol.ConnectorSpecification, error)
	// Check verifies that, given a configuration, data can be accessed properly
	Check(
		msgr messenger.Messenger,
		cfgPsr messenger.ConfigParser,
	) error
	// Discover returns the schema which describes the structure of the data
	// that can be extracted from the source
	Discover(
		msgr messenger.Messenger,
		cfgPsr messenger.ConfigParser,
	) (*protocol.Catalog, error)
	// Read fetches data from the source and
	// communicates all records to the record channel
	// Note: To stop execution, do not use Close method inside the implementation
	// Instead, send a value to the done channel (doneChannel <- true)
	Read(
		cfgdCtg *protocol.ConfiguredCatalog,
		msgr messenger.Messenger,
		cfgPsr messenger.ConfigParser,
		chanHub messenger.ChannelHub,
	)
}
