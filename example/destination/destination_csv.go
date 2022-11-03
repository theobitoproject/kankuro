package main

import (
	"fmt"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

const (
	recordMarshalerWorkers = 4
	csvWriterWorkers       = 2

	minLimitAllowed = 2
	maxLimitAllowed = 100
)

type destinationCsv struct{}

type sourceConfiguration struct {
	Limit int `json:"limit"`
}

func newDestinationCsv() *destinationCsv {
	return &destinationCsv{}
}

// Spec returns the schema which described how the destination connector can be configured
func (d *destinationCsv) Spec(
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
) (*protocol.ConnectorSpecification, error) {
	return &protocol.ConnectorSpecification{
		DocumentationURL:      "https://example-csv-api.com/",
		ChangeLogURL:          "https://example-csv-api.com/",
		SupportsIncremental:   false,
		SupportsNormalization: false,
		SupportsDBT:           false,
		SupportedDestinationSyncModes: []protocol.DestinationSyncMode{
			protocol.DestinationSyncModeOverwrite,
		},
		ConnectionSpecification: protocol.ConnectionSpecification{
			Title:       "Golang - Local CSV",
			Description: "Example CSV",
			Type:        "object",
			Required:    []protocol.PropertyName{"limit"},
			Properties: protocol.Properties{
				Properties: map[protocol.PropertyName]protocol.PropertySpec{
					"limit": {
						Description: fmt.Sprintf(
							"max number of element to pull per instance. Allowed values between %d and %d",
							minLimitAllowed,
							maxLimitAllowed,
						),
						PropertyType: protocol.PropertyType{
							Type: []protocol.PropType{
								protocol.Integer,
							},
						},
					},
				},
			},
		},
	}, nil
}

// Check verifies that, given a configuration, data can be accessed properly
func (d *destinationCsv) Check(
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
) error {
	err := mw.WriteLog(protocol.LogLevelInfo, "checking random api source")
	if err != nil {
		return err
	}

	var sc sourceConfiguration
	err = cp.UnmarshalSourceConfigPath(&sc)
	if err != nil {
		return err
	}

	if sc.Limit < minLimitAllowed {
		return fmt.Errorf(
			"limit configuration value must be greater than or equal to %d",
			minLimitAllowed,
		)
	}

	if sc.Limit > maxLimitAllowed {
		return fmt.Errorf(
			"limit configuration value must be less than or equal to %d",
			maxLimitAllowed,
		)
	}

	return nil
}

// Write takes the data from the record channel
// and stores it in the destination
// Note: all channels except record channel from hub needs to be closed
func (d *destinationCsv) Write(
	cc *protocol.ConfiguredCatalog,
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
	hub messenger.ChannelHub,
) {
	err := mw.WriteLog(protocol.LogLevelInfo, "running write from csv dst")
	if err != nil {
		hub.GetErrorChannel() <- err
	}

	var sc sourceConfiguration
	err = cp.UnmarshalSourceConfigPath(&sc)
	if err != nil {
		hub.GetErrorChannel() <- err
		return
	}

	csvRecordChan := newCsvRecordChannel()
	recordMarshalerWorkersChan := make(chan bool)
	csvWriterWorkersChan := make(chan bool)

	rm := newRecordMarshaler(hub, csvRecordChan, recordMarshalerWorkersChan)
	rm.writeHeaders(cc.Streams)
	for i := 0; i < recordMarshalerWorkers; i++ {
		rm.addWorker()
	}

	cw := newCsvWriter(hub, csvRecordChan, csvWriterWorkersChan)
	for i := 0; i < csvWriterWorkers; i++ {
		cw.addWorker()
	}

	for i := 0; i < recordMarshalerWorkers; i++ {
		<-recordMarshalerWorkersChan
	}
	close(csvRecordChan)
	for i := 0; i < csvWriterWorkers; i++ {
		<-csvWriterWorkersChan
	}

	close(recordMarshalerWorkersChan)
	close(csvWriterWorkersChan)

	cw.closeAndFlush()

	close(hub.GetErrorChannel())
}
