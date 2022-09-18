package protocol

const (
	// MsgTypeRecord defines the type of message for a record
	MsgTypeRecord msgType = "RECORD"
	// MsgTypeState defines the type of message for a state
	MsgTypeState msgType = "STATE"
	// MsgTypeLog defines the type of message for a log
	MsgTypeLog msgType = "LOG"
	// MsgTypeConnectionStat defines the type of message for a connection status
	MsgTypeConnectionStat msgType = "CONNECTION_STATUS"
	// MsgTypeCatalog defines the type of message for a catalog
	MsgTypeCatalog msgType = "CATALOG"
	// MsgTypeSpec defines the type of message for a spec
	MsgTypeSpec msgType = "SPEC"
)

type msgType string
