package protocol

const (
	// CheckStatusSuccess defines the success status when checking the connector
	CheckStatusSuccess CheckStatus = "SUCCEEDED"
	// CheckStatusFailed defines the failed status when checking the connector
	CheckStatusFailed CheckStatus = "FAILED"
)

// ConnectionStatus defines the connection status object to define the state of the connector
type connectionStatus struct {
	Status CheckStatus `json:"status"`
}

// CheckStatus defines the status when checking connector
type CheckStatus string
