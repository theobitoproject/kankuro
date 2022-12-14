package protocol

import (
	"encoding/json"
	"fmt"
)

// Message is a wrapper for the output of each method in actor interface
type Message interface {
	// MarshalMessage returns the JSON encoding of the message
	MarshalMessage() ([]byte, error)
}

type AirbyteMessage struct {
	Type                   MessageType             `json:"type"`
	Record                 *Record                 `json:"record,omitempty"`
	State                  *State                  `json:"state,omitempty"`
	Log                    *Log                    `json:"log,omitempty"`
	ConnectorSpecification *ConnectorSpecification `json:"spec,omitempty"`
	ConnectionStatus       *ConnectionStatus       `json:"connectionStatus,omitempty"`
	Catalog                *Catalog                `json:"catalog,omitempty"`
}

// NewRecordMessage creates a new instance related to a record message
func NewRecordMessage(record *Record) (Message, error) {
	if record == nil {
		return nil, fmt.Errorf("record cannot be empty")
	}

	return &AirbyteMessage{
		Type:   MessageTypeRecord,
		Record: record,
	}, nil
}

// NewStateMessage creates a new instance related to a state message
func NewStateMessage(state *State) (Message, error) {
	if state == nil {
		return nil, fmt.Errorf("state cannot be empty")
	}

	return &AirbyteMessage{
		Type:  MessageTypeState,
		State: state,
	}, nil
}

// NewLogMessage creates a new instance related to a log message
func NewLogMessage(log *Log) (Message, error) {
	if log == nil {
		return nil, fmt.Errorf("log cannot be empty")
	}

	return &AirbyteMessage{
		Type: MessageTypeLog,
		Log:  log,
	}, nil
}

// NewConnectorSpecificationMessage creates a new instance related to a connector specification message
func NewConnectorSpecificationMessage(connectorSpecification *ConnectorSpecification) (Message, error) {
	if connectorSpecification == nil {
		return nil, fmt.Errorf("connectorSpecification cannot be empty")
	}

	return &AirbyteMessage{
		Type:                   MessageTypeSpec,
		ConnectorSpecification: connectorSpecification,
	}, nil
}

// NewConnectionStatusMessage creates a new instance related to a connection status message
func NewConnectionStatusMessage(connectionStatus *ConnectionStatus) (Message, error) {
	if connectionStatus == nil {
		return nil, fmt.Errorf("connectionStatus cannot be empty")
	}

	return &AirbyteMessage{
		Type:             MessageTypeConnectionStat,
		ConnectionStatus: connectionStatus,
	}, nil
}

// NewCatalogMessage creates a new instance related to a catalog message
func NewCatalogMessage(catalog *Catalog) (Message, error) {
	if catalog == nil {
		return nil, fmt.Errorf("catalog cannot be empty")
	}

	return &AirbyteMessage{
		Type:    MessageTypeCatalog,
		Catalog: catalog,
	}, nil
}

// MarshalMessage returns the JSON encoding of the message
func (m *AirbyteMessage) MarshalMessage() ([]byte, error) {
	return json.Marshal(m)
}
