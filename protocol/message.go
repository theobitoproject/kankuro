package protocol

import (
	"encoding/json"
	"errors"
)

var errInvalidTypePayload = errors.New("message type and payload are invalid")

// Message contains all information that is generated from a source and passed to a destination
// TODO: improve this description
type Message struct {
	Type                   msgType                 `json:"type"`
	Record                 *record                 `json:"record,omitempty"`
	State                  *state                  `json:"state,omitempty"`
	LogMessage             *logMessage             `json:"log,omitempty"`
	ConnectorSpecification *ConnectorSpecification `json:"spec,omitempty"`
	ConnectionStatus       *connectionStatus       `json:"connectionStatus,omitempty"`
	Catalog                *Catalog                `json:"catalog,omitempty"`
}

// message MarshalJSON is a custom marshaller which validates the messageType with the sub-struct
func (m *Message) MarshalJSON() ([]byte, error) {
	// TODO: add missing cases and default case
	switch m.Type {
	case msgTypeRecord:
		if m.Record == nil ||
			m.State != nil ||
			m.LogMessage != nil ||
			m.ConnectorSpecification != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	case msgTypeState:
		if m.State == nil ||
			m.Record != nil ||
			m.LogMessage != nil ||
			m.ConnectorSpecification != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	case msgTypeLog:
		if m.LogMessage == nil ||
			m.Record != nil ||
			m.State != nil ||
			m.ConnectorSpecification != nil ||
			m.ConnectionStatus != nil ||
			m.Catalog != nil {
			return nil, errInvalidTypePayload
		}
	}

	type m2 Message
	return json.Marshal(m2(*m))
}
