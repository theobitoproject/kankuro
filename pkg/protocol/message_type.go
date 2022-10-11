package protocol

const (
	// MessageTypeRecord defines the type of message for a record
	MessageTypeRecord MessageType = "RECORD"
	// MessageTypeState defines the type of message for a state
	MessageTypeState MessageType = "STATE"
	// MessageTypeLog defines the type of message for a log
	MessageTypeLog MessageType = "LOG"
	// MessageTypeConnectionStat defines the type of message for a connection status
	MessageTypeConnectionStat MessageType = "CONNECTION_STATUS"
	// MessageTypeCatalog defines the type of message for a catalog
	MessageTypeCatalog MessageType = "CATALOG"
	// MessageTypeSpec defines the type of message for a spec
	MessageTypeSpec MessageType = "SPEC"
)

type MessageType string
