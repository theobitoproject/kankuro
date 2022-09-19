package protocol

import "io"

// PrivateMessenger defines an implementation to send messages
// This PrivateMessenger should NOT be available for the connector implementations
type PrivateMessenger interface {
	// WriteConnectionStat writes a connection check status
	WriteConnectionStat(status checkStatus) error
	// WriteLog writes a catalog
	WriteCatalog(catalog *Catalog) error
	// WriteLog writes a connector specification
	WriteSpec(connectorSpecification *ConnectorSpecification) error
}

type privateMessenger struct {
	writer io.Writer
}

// NewPrivateMessenger creates a new instance of PrivateMessenger
func NewPrivateMessenger(writer io.Writer) PrivateMessenger {
	return privateMessenger{writer}
}

// WriteConnectionStat writes a connection check status
func (pm privateMessenger) WriteConnectionStat(status checkStatus) error {
	return write(pm.writer, &Message{
		Type: msgTypeConnectionStat,
		ConnectionStatus: &connectionStatus{
			Status: status,
		},
	})
}

// WriteLog writes a catalog
func (pm privateMessenger) WriteCatalog(catalog *Catalog) error {
	return write(pm.writer, &Message{
		Type:    msgTypeCatalog,
		Catalog: catalog,
	})
}

// WriteLog writes a connector specification
func (pm privateMessenger) WriteSpec(connectorSpecification *ConnectorSpecification) error {
	return write(pm.writer, &Message{
		Type:                   msgTypeSpec,
		ConnectorSpecification: connectorSpecification,
	})
}
