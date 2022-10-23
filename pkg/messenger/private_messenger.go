package messenger

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
	"github.com/theobitoproject/kankuro/tools"
)

// PrivateMessenger defines an implementation to send messages
// This PrivateMessenger should NOT be available for the connector implementations
type PrivateMessenger interface {
	// WriteRecord writes a record
	WriteRecord(record *protocol.Record) error
	// WriteConnectionStat writes a connection check status
	WriteConnectionStat(status protocol.CheckStatus) error
	// WriteLog writes a catalog
	WriteCatalog(catalog *protocol.Catalog) error
	// WriteLog writes a connector specification
	WriteSpec(connectorSpecification *protocol.ConnectorSpecification) error
}

type privateMessenger struct {
	writer io.Writer
	timer  tools.Timer
}

// NewPrivateMessenger creates a new instance of PrivateMessenger
func NewPrivateMessenger(writer io.Writer, timer tools.Timer) PrivateMessenger {
	return &privateMessenger{writer, timer}
}

// WriteRecord writes a record
func (pm *privateMessenger) WriteRecord(record *protocol.Record) error {

	// fallback: if emitted at is not set, the set it
	if record.EmittedAt == 0 {
		record.EmittedAt = pm.timer.Now().UnixMilli()
	}

	message, err := protocol.NewRecordMessage(record)
	if err != nil {
		return err
	}

	return write(pm.writer, &message)
}

// WriteConnectionStat writes a connection check status
func (pm *privateMessenger) WriteConnectionStat(status protocol.CheckStatus) error {
	message, err := protocol.NewConnectionStatusMessage(&protocol.ConnectionStatus{
		Status: status,
	})
	if err != nil {
		return err
	}

	return write(pm.writer, &message)
}

// WriteLog writes a catalog
func (pm *privateMessenger) WriteCatalog(catalog *protocol.Catalog) error {
	message, err := protocol.NewCatalogMessage(catalog)
	if err != nil {
		return err
	}

	return write(pm.writer, &message)
}

// WriteLog writes a connector specification
func (pm *privateMessenger) WriteSpec(connectorSpecification *protocol.ConnectorSpecification) error {
	message, err := protocol.NewConnectorSpecificationMessage(connectorSpecification)
	if err != nil {
		return err
	}

	return write(pm.writer, &message)
}
