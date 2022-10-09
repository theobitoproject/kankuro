package destination

import (
	"fmt"

	"github.com/theobitoproject/kankuro/protocol"
)

type DestinationRunner struct {
	dst              Destination
	messenger        protocol.Messenger
	privateMessenger protocol.PrivateMessenger
	configParser     protocol.ConfigParser
}

func NewDestinationRunner(
	dst Destination,
	messenger protocol.Messenger,
	privateMessenger protocol.PrivateMessenger,
	configParser protocol.ConfigParser,
) DestinationRunner {
	return DestinationRunner{
		dst,
		messenger,
		privateMessenger,
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
	spec, err := dr.dst.Spec(dr.messenger, dr.configParser)
	if err != nil {
		// TODO: handle error from dr.messenger.WriteLog
		dr.messenger.WriteLog(protocol.LogLevelError, "failed"+err.Error())
		return err
	}

	return dr.privateMessenger.WriteSpec(spec)
}

func (dr DestinationRunner) check() error {
	err := dr.dst.Check(dr.messenger, dr.configParser)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		// TODO: log error
		checkStatus = protocol.CheckStatusFailed
	}

	return dr.privateMessenger.WriteConnectionStat(checkStatus)
}

func (dr DestinationRunner) write() error {
	dr.messenger.WriteLog(protocol.LogLevelInfo, "writing from dst runner...")

	var incat protocol.ConfiguredCatalog

	err := dr.configParser.UnmarshalCatalogPath(&incat)
	if err != nil {
		// TODO: log error
		return err
	}

	err = dr.dst.Write(&incat, dr.messenger, dr.configParser)
	if err != nil {
		// TODO: log error
		return err
	}

	return nil
}
