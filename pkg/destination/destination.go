package destination

import (
	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// Destination is the only interface you need to define to create your destination!
type Destination interface {
	// Spec returns the input "form" spec needed for your source
	Spec(
		msgr messenger.Messenger,
		configParser messenger.ConfigParser,
	) (protocol.ConnectorSpecification, error)
	// Check verifies the source - usually verify creds/connection etc.
	Check(
		msgr messenger.Messenger,
		configParser messenger.ConfigParser,
	) error
	// Read will read the actual data from your source and use
	// tracker.Record(), tracker.State() and tracker.Log() to sync data
	// with airbyte/destinations
	// MessageTracker is thread-safe and so it is completely find to
	// spin off goroutines to sync your data (just don't forget your waitgroups :))
	// returning an error from this will cancel the sync and returning a nil
	// from this will successfully end the sync
	Write(
		configuredCat *protocol.ConfiguredCatalog,
		msgr messenger.Messenger,
		configParser messenger.ConfigParser,
	) error
}
