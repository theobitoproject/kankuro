package protocol

import (
	"encoding/json"
	"errors"
)

var errInvalidTypePayload = errors.New("message type and payload are invalid")

// Message contains all information that is generated from a source and passed to a destination
// TODO: improve this description
type Message struct {
	Type                    msgType `json:"type"`
	*record                 `json:"record,omitempty"`
	*state                  `json:"state,omitempty"`
	*logMessage             `json:"log,omitempty"`
	*ConnectorSpecification `json:"spec,omitempty"`
	*ConnectionStatus       `json:"connectionStatus,omitempty"`
	*Catalog                `json:"catalog,omitempty"`
}

// message MarshalJSON is a custom marshaller which validates the messageType with the sub-struct
func (m *Message) MarshalJSON() ([]byte, error) {
	switch m.Type {
	case MsgTypeRecord:
		if m.record == nil ||
			m.state != nil ||
			m.logMessage != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	case MsgTypeState:
		if m.state == nil ||
			m.record != nil ||
			m.logMessage != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	case MsgTypeLog:
		if m.logMessage == nil ||
			m.record != nil ||
			m.state != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	}

	type m2 Message
	return json.Marshal(m2(*m))
}
