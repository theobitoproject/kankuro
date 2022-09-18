package source

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/theobitoproject/kankuro/parsers"
	"github.com/theobitoproject/kankuro/protocol"
	"github.com/theobitoproject/kankuro/trackers"
	"github.com/theobitoproject/kankuro/writers"
)

// SourceRunner acts as an "orchestrator" of sorts to run your source for you
type SourceRunner struct {
	w          io.Writer
	src        Source
	msgTracker trackers.MessageTracker
}

type lastSyncTime struct {
	Timestamp int64 `json:"timestamp"`
}

// NewSourceRunner takes your defined Source and plugs it in with the rest of airbyte
func NewSourceRunner(src Source, w io.Writer) SourceRunner {
	w = writers.NewSafeWriter(w)
	msgTracker := trackers.MessageTracker{
		Record: protocol.NewRecordWriter(w),
		State:  protocol.NewStateWriter(w),
		Log:    protocol.NewLogWriter(w),
	}

	return SourceRunner{
		w:          w,
		src:        src,
		msgTracker: msgTracker,
	}
}

// Start starts your source
// Example usage would look like this in your main.go
//  func() main {
// 	src := newCoolSource()
// 	runner := airbyte.NewSourceRunner(src)
// 	err := runner.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	 }
//  }
// Yes, it really is that easy!
func (sr SourceRunner) Start() error {
	switch protocol.Cmd(os.Args[1]) {
	case protocol.CmdSpec:
		spec, err := sr.src.Spec(trackers.LogTracker{
			Log: sr.msgTracker.Log,
		})
		if err != nil {
			sr.msgTracker.Log(protocol.LogLevelError, "failed"+err.Error())
			return err
		}
		return protocol.Write(sr.w, &protocol.Message{
			Type:                   protocol.MsgTypeSpec,
			ConnectorSpecification: spec,
		})

	case protocol.CmdCheck:
		inP, err := parsers.GetSourceConfigPath()
		if err != nil {
			return err
		}
		err = sr.src.Check(inP, trackers.LogTracker{
			Log: sr.msgTracker.Log,
		})
		if err != nil {
			log.Println(err)
			return protocol.Write(sr.w, &protocol.Message{
				Type: protocol.MsgTypeConnectionStat,
				ConnectionStatus: &protocol.ConnectionStatus{
					Status: protocol.CheckStatusFailed,
				},
			})
		}

		return protocol.Write(sr.w, &protocol.Message{
			Type: protocol.MsgTypeConnectionStat,
			ConnectionStatus: &protocol.ConnectionStatus{
				Status: protocol.CheckStatusSuccess,
			},
		})

	case protocol.CmdDiscover:
		inP, err := parsers.GetSourceConfigPath()
		if err != nil {
			return err
		}
		ct, err := sr.src.Discover(inP, trackers.LogTracker{
			Log: sr.msgTracker.Log},
		)
		if err != nil {
			return err
		}
		return protocol.Write(sr.w, &protocol.Message{
			Type:    protocol.MsgTypeCatalog,
			Catalog: ct,
		})

	case protocol.CmdRead:
		var incat protocol.ConfiguredCatalog
		p, err := parsers.GetCatalogPath()
		if err != nil {
			return err
		}

		err = parsers.UnmarshalFromPath(p, &incat)
		if err != nil {
			return err
		}

		srp, err := parsers.GetSourceConfigPath()
		if err != nil {
			return err
		}

		stp, err := parsers.GetStatePath()
		if err != nil {
			return err
		}

		err = sr.src.Read(srp, stp, &incat, sr.msgTracker)
		if err != nil {
			return err
		}

		err = sr.msgTracker.State(&lastSyncTime{
			Timestamp: time.Now().UnixMilli(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
