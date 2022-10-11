package source

import (
	"fmt"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// SourceRunner acts as an "orchestrator" of sorts to run your source for you
type SourceRunner struct {
	src          Source
	msgr         messenger.Messenger
	prvtMsgr     messenger.PrivateMessenger
	configParser messenger.ConfigParser
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(
	src Source,
	msgr messenger.Messenger,
	prvtMsgr messenger.PrivateMessenger,
	configParser messenger.ConfigParser,
) SourceRunner {
	//  TODO: should checks be added to catch nil pointers?
	return SourceRunner{
		src,
		msgr,
		prvtMsgr,
		configParser,
	}
}

// Start performs actions related to a single Airbyte command (spec, check, read, write, etc)
// Example usage would look like this in your main.go
//  func() main {
// 	src := newCoolSource()
// 	runner := airbyte.NewSourceRunner(src, os.Stdout, os.Args)
// 	err := runner.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	 }
//  }
// Yes, it really is that easy!
func (sr SourceRunner) Start() (err error) {
	mainCmd, err := sr.configParser.GetMainCommand()
	if err != nil {
		return err
	}

	if mainCmd.IsZero() {
		return fmt.Errorf("main command is required")
	}

	switch mainCmd {
	case protocol.CmdSpec:
		err = sr.spec()

	case protocol.CmdCheck:
		err = sr.check()

	case protocol.CmdDiscover:
		err = sr.discover()

	case protocol.CmdRead:
		err = sr.read()

	default:
		err = fmt.Errorf("command not supported: %s", mainCmd)
	}

	return err
}

func (sr SourceRunner) spec() error {
	spec, err := sr.src.Spec(sr.msgr, sr.configParser)
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
	err := sr.src.Check(sr.msgr, sr.configParser)

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
	ct, err := sr.src.Discover(sr.msgr, sr.configParser)
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

	err := sr.configParser.UnmarshalCatalogPath(&incat)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed unmarshaling catalog catalog: %v", err).Error(),
		)
		return err
	}

	err = sr.src.Read(&incat, sr.msgr, sr.configParser)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		sr.msgr.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running source read: %v", err).Error(),
		)
		return err
	}

	return nil
}
