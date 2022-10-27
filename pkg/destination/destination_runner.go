package destination

import (
	"fmt"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

type DestinationRunner struct {
	dst Destination

	mw  messenger.MessageWriter
	pmw messenger.PrivateMessageWriter

	mr messenger.MessageReader

	cp messenger.ConfigParser

	hub messenger.ChannelHub
}

func NewDestinationRunner(
	dst Destination,
	mw messenger.MessageWriter,
	pmw messenger.PrivateMessageWriter,
	mr messenger.MessageReader,
	cp messenger.ConfigParser,
	hub messenger.ChannelHub,
) DestinationRunner {
	return DestinationRunner{
		dst,
		mw,
		pmw,
		mr,
		cp,
		hub,
	}
}

func (dr DestinationRunner) Start() (err error) {
	mainCmd, err := dr.cp.GetMainCommand()
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
	spec, err := dr.dst.Spec(dr.mw, dr.cp)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		dr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running destination spec: %v", err).Error(),
		)
		return err
	}

	return dr.pmw.WriteSpec(spec)
}

func (dr DestinationRunner) check() error {
	err := dr.dst.Check(dr.mw, dr.cp)

	checkStatus := protocol.CheckStatusSuccess
	if err != nil {
		checkStatus = protocol.CheckStatusFailed

		// write log and don't return error
		// because we need to write success/failed connection status message
		// TODO: is there a good way to handle error from messenger.WriteLog?
		dr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed running destination check: %v", err).Error(),
		)
	}

	err = dr.pmw.WriteConnectionStat(checkStatus)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		dr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed writing connection stat: %v", err).Error(),
		)
		return err
	}

	return err
}

func (dr DestinationRunner) write() error {
	dr.mw.WriteLog(protocol.LogLevelInfo, "writing from dst runner...")

	var cc protocol.ConfiguredCatalog

	err := dr.cp.UnmarshalCatalogPath(&cc)
	if err != nil {
		// TODO: is there a good way to handle error from messenger.WriteLog?
		dr.mw.WriteLog(
			protocol.LogLevelError,
			fmt.Errorf("failed unmarshaling catalog: %v", err).Error(),
		)
		return err
	}

	dr.dst.Write(
		&cc,
		dr.mw,
		dr.cp,
		dr.hub,
	)

	doneChannel := messenger.NewDoneChannel()

	go func() {
		for {
			select {

			case _, channelOpen := <-dr.hub.GetClosingChannel():
				if !channelOpen {
					doneChannel <- true
				}

			case err, channelOpen := <-dr.hub.GetErrorChannel():
				if channelOpen {
					dr.mw.WriteLog(
						protocol.LogLevelError,
						fmt.Errorf("failed running destination write: %v", err).Error(),
					)

				} else {
					doneChannel <- true
				}

			case _, channelOpen := <-dr.hub.GetRecordChannel():
				if !channelOpen {
					doneChannel <- true
				}
			}
		}
	}()

	go dr.mr.Read(dr.hub)

	// Wait for three channels to be closed before continue
	// - recordChannel
	// - errorChannel
	// - closinghannel
	<-doneChannel
	<-doneChannel
	<-doneChannel

	dr.mw.WriteLog(
		protocol.LogLevelInfo,
		"writing has finished",
	)

	return nil
}
