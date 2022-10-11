package protocol

const (
	// SyncModeFullRefresh means the data will be wiped and fully synced on each run
	SyncModeFullRefresh SyncMode = "full_refresh"
	// SyncModeIncremental is used for incremental syncs
	SyncModeIncremental SyncMode = "incremental"
)

var (
	// DestinationSyncModeAppend is used for the destination to know it needs to append data
	DestinationSyncModeAppend DestinationSyncMode = "append"
	// DestinationSyncModeOverwrite is used to indicate the destination should overwrite data
	DestinationSyncModeOverwrite DestinationSyncMode = "overwrite"
)

// SyncMode defines the modes that your source is able to sync in
type SyncMode string

// DestinationSyncMode represents how the destination should interpret your data
type DestinationSyncMode string
