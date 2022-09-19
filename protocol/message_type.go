package protocol

const (
	// MsgTypeRecord defines the type of message for a record
	msgTypeRecord msgType = "RECORD"
	// MsgTypeState defines the type of message for a state
	msgTypeState msgType = "STATE"
	// MsgTypeLog defines the type of message for a log
	msgTypeLog msgType = "LOG"
	// MsgTypeConnectionStat defines the type of message for a connection status
	msgTypeConnectionStat msgType = "CONNECTION_STATUS"
	// MsgTypeCatalog defines the type of message for a catalog
	msgTypeCatalog msgType = "CATALOG"
	// MsgTypeSpec defines the type of message for a spec
	msgTypeSpec msgType = "SPEC"
)

type msgType string
