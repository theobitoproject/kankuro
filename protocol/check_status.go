package protocol

const (
	// CheckStatusSuccess defines the success status when checking the connector
	CheckStatusSuccess checkStatus = "SUCCEEDED"
	// CheckStatusFailed defines the failed status when checking the connector
	CheckStatusFailed checkStatus = "FAILED"
)

// ConnectionStatus defines the connection status object to define the state of the connector
type connectionStatus struct {
	Status checkStatus `json:"status"`
}

type checkStatus string
