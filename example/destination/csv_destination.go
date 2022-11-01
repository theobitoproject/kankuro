package main

import (
	"fmt"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// MaxLimitAllowed is the max amount of objects that can be fecthed from random API platform
const (
	MinLimitAllowed = 2
	MaxLimitAllowed = 100
)

type csvDestination struct{}

type sourceConfiguration struct {
	Limit int `json:"limit"`
}

func newCsvDestination() *csvDestination {
	return &csvDestination{}
}

// Spec returns the schema which described how the destination connector can be configured
func (d *csvDestination) Spec(
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
) (*protocol.ConnectorSpecification, error) {
	return &protocol.ConnectorSpecification{
		DocumentationURL:      "https://example-csv-api.com/",
		ChangeLogURL:          "https://example-csv-api.com/",
		SupportsIncremental:   false,
		SupportsNormalization: true,
		SupportsDBT:           true,
		SupportedDestinationSyncModes: []protocol.DestinationSyncMode{
			protocol.DestinationSyncModeOverwrite,
		},
		ConnectionSpecification: protocol.ConnectionSpecification{
			Title:       "Example CSV",
			Description: "Example CSV",
			Type:        "object",
			Required:    []protocol.PropertyName{"limit"},
			Properties: protocol.Properties{
				Properties: map[protocol.PropertyName]protocol.PropertySpec{
					"limit": {
						Description: fmt.Sprintf(
							"max number of element to pull per instance. Allowed values between %d and %d",
							MinLimitAllowed,
							MaxLimitAllowed,
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
func (d *csvDestination) Check(
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

	if sc.Limit < MinLimitAllowed {
		msg := fmt.Sprintf("limit configuration value must be greater than or equal to %d", MinLimitAllowed)
		mw.WriteLog(protocol.LogLevelInfo, msg)
		return fmt.Errorf(msg)
	}

	if sc.Limit > MaxLimitAllowed {
		msg := fmt.Sprintf("limit configuration value must be less than or equal to %d", MaxLimitAllowed)
		mw.WriteLog(protocol.LogLevelInfo, msg)
		return fmt.Errorf(msg)
	}

	return nil
}

// Write takes the data from the record channel
// and stores it in the destination
// Note: all channels except record channel from hub needs to be closed
func (d *csvDestination) Write(
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

	doneStreamChannel := make(chan bool)

	go func() {
		for i := 0; i < len(cc.Streams); i++ {
			<-doneStreamChannel
		}

		close(hub.GetErrorChannel())
	}()

	csvRecordChann := newCsvRecordChannel()

	rm := newCsvRecordMarshaler(hub, csvRecordChann, doneStreamChannel)
	rm.addWorker()

	cfh := newCsvFileHandler(hub, csvRecordChann, doneStreamChannel)
	cfh.addWorker()
}
