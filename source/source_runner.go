package source

import (
	"fmt"
	"time"

	"github.com/theobitoproject/kankuro/protocol"
)

// SourceRunner acts as an "orchestrator" of sorts to run your source for you
type SourceRunner struct {
	src              Source
	messenger        protocol.Messenger
	privateMessenger protocol.PrivateMessenger
	configParser     protocol.ConfigParser
}

type lastSyncTime struct {
	Timestamp int64 `json:"timestamp"`
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(
	src Source,
	messenger protocol.Messenger,
	privateMessenger protocol.PrivateMessenger,
	configParser protocol.ConfigParser,
) SourceRunner {
	//  TODO: add checks to catch nil pointers
	return SourceRunner{
		src,
		messenger,
		privateMessenger,
		configParser,
	}
}

// Start starts your source
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
	spec, err := sr.src.Spec(sr.messenger, sr.configParser)
	if err != nil {
		// TODO: handle error from sr.messenger.WriteLog
		sr.messenger.WriteLog(protocol.LogLevelError, "failed"+err.Error())
		return err
	}

	return sr.privateMessenger.WriteSpec(spec)
}

func (sr SourceRunner) check() error {
	err := sr.src.Check(sr.messenger, sr.configParser)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		// TODO: log error
		checkStatus = protocol.CheckStatusFailed
	}

	return sr.privateMessenger.WriteConnectionStat(checkStatus)
}

func (sr SourceRunner) discover() error {
	ct, err := sr.src.Discover(sr.messenger, sr.configParser)
	if err != nil {
		// TODO: log error
		return err
	}

	return sr.privateMessenger.WriteCatalog(ct)
}

func (sr SourceRunner) read() error {
	var incat protocol.ConfiguredCatalog

	err := sr.configParser.UnmarshalCatalogPath(&incat)
	if err != nil {
		// TODO: log error
		return err
	}

	err = sr.src.Read(&incat, sr.messenger, sr.configParser)
	if err != nil {
		// TODO: log error
		return err
	}

	currentLastSyncTime := getCurrentLastSyncTime()
	return sr.messenger.WriteState(&currentLastSyncTime)
}

func getCurrentLastSyncTime() lastSyncTime {
	return lastSyncTime{
		Timestamp: time.Now().UnixMilli(),
	}
}
