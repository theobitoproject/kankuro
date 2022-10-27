package messenger

import (
	"io"

	"github.com/theobitoproject/kankuro/pkg/protocol"
	"github.com/theobitoproject/kankuro/tools"
)

// PrivateMessageWriter defines an implementation to send messages
// This PrivateMessageWriter should NOT be available for the connector implementations
type PrivateMessageWriter interface {
	// WriteRecord writes a record
	WriteRecord(record *protocol.Record) error
	// WriteConnectionStat writes a connection check status
	WriteConnectionStat(status protocol.CheckStatus) error
	// WriteLog writes a catalog
	WriteCatalog(catalog *protocol.Catalog) error
	// WriteLog writes a connector specification
	WriteSpec(connectorSpecification *protocol.ConnectorSpecification) error
}

type privateMessageWriter struct {
	writer io.Writer
	timer  tools.Timer
}

// NewPrivateMessageWriter creates a new instance of PrivateMessageWriter
func NewPrivateMessageWriter(writer io.Writer, timer tools.Timer) PrivateMessageWriter {
	return &privateMessageWriter{writer, timer}
}

// WriteRecord writes a record
func (pmw *privateMessageWriter) WriteRecord(record *protocol.Record) error {

	// fallback: if emitted at is not set, the set it
	if record.EmittedAt == 0 {
		record.EmittedAt = pmw.timer.Now().UnixMilli()
	}

	message, err := protocol.NewRecordMessage(record)
	if err != nil {
		return err
	}

	return write(pmw.writer, &message)
}

// WriteConnectionStat writes a connection check status
func (pmw *privateMessageWriter) WriteConnectionStat(status protocol.CheckStatus) error {
	message, err := protocol.NewConnectionStatusMessage(&protocol.ConnectionStatus{
		Status: status,
	})
	if err != nil {
		return err
	}

	return write(pmw.writer, &message)
}

// WriteLog writes a catalog
func (pmw *privateMessageWriter) WriteCatalog(catalog *protocol.Catalog) error {
	message, err := protocol.NewCatalogMessage(catalog)
	if err != nil {
		return err
	}

	return write(pmw.writer, &message)
}

// WriteLog writes a connector specification
func (pmw *privateMessageWriter) WriteSpec(connectorSpecification *protocol.ConnectorSpecification) error {
	message, err := protocol.NewConnectorSpecificationMessage(connectorSpecification)
	if err != nil {
		return err
	}

	return write(pmw.writer, &message)
}
