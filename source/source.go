package source

import (
	"github.com/theobitoproject/kankuro/protocol"
	"github.com/theobitoproject/kankuro/trackers"
)

// Source is the only interface you need to define to create your source!
type Source interface {
	// Spec returns the input "form" spec needed for your source
	Spec(logTracker trackers.LogTracker) (*protocol.ConnectorSpecification, error)
	// Check verifies the source - usually verify creds/connection etc.
	Check(srcCfgPath string, logTracker trackers.LogTracker) error
	// Discover returns the schema of the data you want to sync
	Discover(srcConfigPath string, logTracker trackers.LogTracker) (*protocol.Catalog, error)
	// Read will read the actual data from your source and use
	// tracker.Record(), tracker.State() and tracker.Log() to sync data
	// with airbyte/destinations
	// MessageTracker is thread-safe and so it is completely find to
	// spin off goroutines to sync your data (just don't forget your waitgroups :))
	// returning an error from this will cancel the sync and returning a nil
	// from this will successfully end the sync
	Read(
		srcCfgPath string,
		prevStatePath string,
		configuredCat *protocol.ConfiguredCatalog,
		tracker trackers.MessageTracker,
	) error
}
