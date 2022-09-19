package source

import (
	"io"
	"os"
	"time"

	"github.com/theobitoproject/kankuro/protocol"
)

// SourceRunner acts as an "orchestrator" of sorts to run your source for you
type SourceRunner struct {
	src              Source
	messenger        protocol.Messenger
	privateMessenger protocol.PrivateMessenger
}

type lastSyncTime struct {
	Timestamp int64 `json:"timestamp"`
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(src Source, w io.Writer) SourceRunner {
	w = newSafeWriter(w)

	return SourceRunner{
		src:              src,
		messenger:        protocol.NewMessenger(w),
		privateMessenger: protocol.NewPrivateMessenger(w),
	}
}

// Start starts your source
// Example usage would look like this in your main.go
//  func() main {
// 	src := newCoolSource()
//  writer := newWriter()
// 	runner := airbyte.NewSourceRunner(src, writer)
// 	err := runner.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	 }
//  }
// Yes, it really is that easy!
func (sr SourceRunner) Start() (err error) {
	switch protocol.Cmd(os.Args[1]) {
	case protocol.CmdSpec:
		err = sr.spec()

	case protocol.CmdCheck:
		err = sr.check()

	case protocol.CmdDiscover:
		err = sr.discover()

	case protocol.CmdRead:
		err = sr.read()
	}

	return err
}

func (sr SourceRunner) spec() error {
	spec, err := sr.src.Spec(sr.messenger)
	if err != nil {
		sr.messenger.WriteLog(protocol.LogLevelError, "failed"+err.Error())
		return err
	}

	return sr.privateMessenger.WriteSpec(spec)
}

func (sr SourceRunner) check() error {
	inP, err := protocol.GetSourceConfigPath()
	if err != nil {
		return err
	}

	err = sr.src.Check(inP, sr.messenger)
	if err != nil {
		return sr.privateMessenger.WriteConnectionStat(protocol.CheckStatusFailed)
	}

	return sr.privateMessenger.WriteConnectionStat(protocol.CheckStatusSuccess)
}

func (sr SourceRunner) discover() error {
	inP, err := protocol.GetSourceConfigPath()
	if err != nil {
		return err
	}

	ct, err := sr.src.Discover(inP, sr.messenger)
	if err != nil {
		return err
	}

	return sr.privateMessenger.WriteCatalog(ct)
}

func (sr SourceRunner) read() error {
	var incat protocol.ConfiguredCatalog
	p, err := protocol.GetCatalogPath()
	if err != nil {
		return err
	}

	err = protocol.UnmarshalFromPath(p, &incat)
	if err != nil {
		return err
	}

	srp, err := protocol.GetSourceConfigPath()
	if err != nil {
		return err
	}

	stp, err := protocol.GetStatePath()
	if err != nil {
		return err
	}

	err = sr.src.Read(srp, stp, &incat, sr.messenger)
	if err != nil {
		return err
	}

	return sr.messenger.WriteState(&lastSyncTime{
		Timestamp: time.Now().UnixMilli(),
	})
}
