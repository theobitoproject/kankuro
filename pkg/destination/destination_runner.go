package destination

import (
	"fmt"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

type DestinationRunner struct {
	dst          Destination
	msgr         messenger.Messenger
	prvtMsgr     messenger.PrivateMessenger
	configParser messenger.ConfigParser
}

func NewDestinationRunner(
	dst Destination,
	msgr messenger.Messenger,
	prvtMsgr messenger.PrivateMessenger,
	configParser messenger.ConfigParser,
) DestinationRunner {
	return DestinationRunner{
		dst,
		msgr,
		prvtMsgr,
		configParser,
	}
}

func (dr DestinationRunner) Start() (err error) {
	mainCmd, err := dr.configParser.GetMainCommand()
	if err != nil {
		return err
	}

	if mainCmd.IsZero() {
		return fmt.Errorf("main command is required")
	}

	switch mainCmd {
	case protocol.CmdSpec:
		err = dr.spec()

	case protocol.CmdCheck:
		err = dr.check()

	case protocol.CmdWrite:
		err = dr.write()

	default:
		err = fmt.Errorf("command not supported: %s", mainCmd)
	}

	return err
}

func (dr DestinationRunner) spec() error {
	spec, err := dr.dst.Spec(dr.msgr, dr.configParser)
	if err != nil {
		// TODO: handle error from dr.msgr.WriteLog
		dr.msgr.WriteLog(protocol.LogLevelError, "failed"+err.Error())
		return err
	}

	return dr.prvtMsgr.WriteSpec(&spec)
}

func (dr DestinationRunner) check() error {
	err := dr.dst.Check(dr.msgr, dr.configParser)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		// TODO: log error
		checkStatus = protocol.CheckStatusFailed
	}

	return dr.prvtMsgr.WriteConnectionStat(checkStatus)
}

func (dr DestinationRunner) write() error {
	dr.msgr.WriteLog(protocol.LogLevelInfo, "writing from dst runner...")

	var incat protocol.ConfiguredCatalog

	err := dr.configParser.UnmarshalCatalogPath(&incat)
	if err != nil {
		// TODO: log error
		return err
	}

	err = dr.dst.Write(&incat, dr.msgr, dr.configParser)
	if err != nil {
		// TODO: log error
		return err
	}

	return nil
}
