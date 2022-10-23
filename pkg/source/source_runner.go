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

	msgr     messenger.Messenger
	prvtMsgr messenger.PrivateMessenger

	cfgPsr messenger.ConfigParser

	chanHub messenger.ChannelHub
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(
	src Source,
	msgr messenger.Messenger,
	prvtMsgr messenger.PrivateMessenger,
	cfgPsr messenger.ConfigParser,
	chanHub messenger.ChannelHub,
) SourceRunner {
	//  TODO: should checks be added to catch nil pointers?
	return SourceRunner{
		src,
		msgr,
		prvtMsgr,
		cfgPsr,
		chanHub,
	}
}

// Start performs actions related to a single Airbyte command (spec, check, read, write, etc)
func (sr SourceRunner) Start() (err error) {
	mainCmd, err := sr.cfgPsr.GetMainCommand()
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
	spec, err := sr.src.Spec(sr.msgr, sr.cfgPsr)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source spec: %v", err).Error(),
		)
		return err
	}

	err = sr.prvtMsgr.WriteSpec(spec)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing spec: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) check() error {
	err := sr.src.Check(sr.msgr, sr.cfgPsr)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		checkStatus = protocol.CheckStatusFailed

		// write log and don't return error
		// because we need to write success/failed connection status message
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source check: %v", err).Error(),
		)
	}

	err = sr.prvtMsgr.WriteConnectionStat(checkStatus)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing connection stat: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) discover() error {
	ct, err := sr.src.Discover(sr.msgr, sr.cfgPsr)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source discover: %v", err).Error(),
		)
		return err
	}

	err = sr.prvtMsgr.WriteCatalog(ct)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing catalog: %v", err).Error(),
		)
		return err
	}

	return err
}

func (sr SourceRunner) read() error {
	var incat protocol.ConfiguredCatalog

	err := sr.cfgPsr.UnmarshalCatalogPath(&incat)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed unmarshaling catalog catalog: %v", err).Error(),
		)
		return err
	}

	go sr.src.Read(
		&incat,
		sr.msgr,
		sr.cfgPsr,
		sr.chanHub,
	)

	for {
		select {

		case <-sr.chanHub.GetDoneChannel():
			sr.src.Close(sr.chanHub)
			sr.msgr.WriteLog(
				protocol.LogLevelInfo,
				"reading has finished",
			)
			return nil

		// in case of any errors, log it and close all channels
		case err = <-sr.chanHub.GetErrorChannel():
			sr.msgr.WriteLog(
				protocol.LogLevelError,
				fmt.Errorf("failed running source read: %v", err).Error(),
			)
			sr.chanHub.GetDoneChannel() <- true
			return err

		case record := <-sr.chanHub.GetRecordChannel():
			err = sr.prvtMsgr.WriteRecord(record)
			if err != nil {
				sr.chanHub.GetErrorChannel() <- err
				return err
			}
		}
	}
}

func (sr *SourceRunner) printConfiguredCatalogOnFile() error {
	ct, err := sr.src.Discover(sr.msgr, sr.cfgPsr)
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
