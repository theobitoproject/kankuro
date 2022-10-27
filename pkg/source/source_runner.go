package source

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// SourceRunner acts as an "orchestrator" of sorts to run your source for you
type SourceRunner struct {
	src Source

	mw  messenger.MessageWriter
	pmw messenger.PrivateMessageWriter

	cp messenger.ConfigParser

	hub messenger.ChannelHub
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(
	src Source,
	mw messenger.MessageWriter,
	pmw messenger.PrivateMessageWriter,
	cp messenger.ConfigParser,
	hub messenger.ChannelHub,
) SourceRunner {
	//  TODO: should checks be added to catch nil pointers?
	return SourceRunner{
		src,
		mw,
		pmw,
		cp,
		hub,
	}
}

// Start performs actions related to a single Airbyte command (spec, check, read, write, etc)
func (sr SourceRunner) Start() (err error) {
	mainCmd, err := sr.cp.GetMainCommand()
	if err != nil {
		return err
	}

	if mainCmd.IsZero() {
		return fmt.Errorf("main command is required")
	}

	switch mainCmd {
	// airbyte commands
	case protocol.CmdSpec:
		err = sr.spec()

	case protocol.CmdCheck:
		err = sr.check()

	case protocol.CmdDiscover:
		err = sr.discover()

	case protocol.CmdRead:
		err = sr.read()

	// kankuro dev commands
	case CmdPrintCatalog:
		err = sr.printConfiguredCatalogOnFile()

	default:
		err = fmt.Errorf("command not supported: %s", mainCmd)
	}

	return err
}

func (sr SourceRunner) spec() error {
	spec, err := sr.src.Spec(sr.mw, sr.cp)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source spec: %v", err).Error(),
		)
		return err
	}

	err = sr.pmw.WriteSpec(spec)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing spec: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) check() error {
	err := sr.src.Check(sr.mw, sr.cp)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		checkStatus = protocol.CheckStatusFailed

		// write log and don't return error
		// because we need to write success/failed connection status message
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source check: %v", err).Error(),
		)
	}

	err = sr.pmw.WriteConnectionStat(checkStatus)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing connection stat: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) discover() error {
	ct, err := sr.src.Discover(sr.mw, sr.cp)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source discover: %v", err).Error(),
		)
		return err
	}

	err = sr.pmw.WriteCatalog(ct)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing catalog: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) read() error {
	var cc protocol.ConfiguredCatalog

	err := sr.cp.UnmarshalCatalogPath(&cc)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed unmarshaling catalog: %v", err).Error(),
		)
		return err
	}

	sr.src.Read(
		&cc,
		sr.mw,
		sr.cp,
		sr.hub,
	)

	doneChannel := messenger.NewDoneChannel()

	go func() {
		for {
			select {

			case _, channelOpen := <-sr.hub.GetClosingChannel():
				if !channelOpen {
					doneChannel <- true
				}

			case err, channelOpen := <-sr.hub.GetErrorChannel():
				if channelOpen {
					sr.mw.WriteLog(
						protocol.LogLevelError,
						fmt.Errorf("failed running source read: %v", err).Error(),
					)

				} else {
					doneChannel <- true
				}

			case record, channelOpen := <-sr.hub.GetRecordChannel():
				if channelOpen {
					err = sr.pmw.WriteRecord(record)
					if err != nil {
						sr.hub.GetErrorChannel() <- err
					}

				} else {
					doneChannel <- true
				}
			}
		}
	}()

	// Wait for three channels to be closed before continue
	// - recordChannel
	// - errorChannel
	// - closinghannel
	<-doneChannel
	<-doneChannel
	<-doneChannel

	sr.mw.WriteLog(
		protocol.LogLevelInfo,
		"reading has finished",
	)

	return nil
}

func (sr *SourceRunner) printConfiguredCatalogOnFile() error {
	ct, err := sr.src.Discover(sr.mw, sr.cp)
	if err != nil {
		return err
	}

	data, err := json.Marshal(ct)
	if err != nil {
		return err
	}

	// TODO: find a good way to define the path of the file
	// where the catalog will be stored
	err = os.MkdirAll("sample_files", 0755)
	if err != nil {
		return err
	}

	return os.WriteFile("sample_files/configured_catalog.json", data, 0755)
}
