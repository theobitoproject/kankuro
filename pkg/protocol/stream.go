package protocol

// Stream defines a single "schema" you'd like to sync - think of this as a table, collection, topic, etc. In airbyte terminology these are "streams"
type Stream struct {
	Name                    string     `json:"name"`
	JSONSchema              Properties `json:"json_schema"`
	SupportedSyncModes      []SyncMode `json:"supported_sync_modes,omitempty"`
	SourceDefinedCursor     bool       `json:"source_defined_cursor,omitempty"`
	DefaultCursorField      []string   `json:"default_cursor_field,omitempty"`
	SourceDefinedPrimaryKey [][]string `json:"source_defined_primary_key,omitempty"`
	Namespace               string     `json:"namespace"`
}

// ConfiguredStream defines a single selected stream to sync
type ConfiguredStream struct {
	Stream              Stream              `json:"stream"`
	SyncMode            SyncMode            `json:"sync_mode"`
	CursorField         []string            `json:"cursor_field"`
	DestinationSyncMode DestinationSyncMode `json:"destination_sync_mode"`
	PrimaryKey          [][]string          `json:"primary_key"`
}
