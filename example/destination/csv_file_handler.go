package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/theobitoproject/kankuro/pkg/messenger"
)

type csvFileHandler struct {
	hub               messenger.ChannelHub
	csvRecordChann    csvRecordChannel
	doneStreamChannel chan bool
	fileWriterPairs   map[string]*fileWriterPair
}

type fileWriterPair struct {
	file   *os.File
	writer *csv.Writer
}

func newCsvFileHandler(
	hub messenger.ChannelHub,
	csvRecordChann csvRecordChannel,
	doneStreamChannel chan bool,
) *csvFileHandler {
	return &csvFileHandler{
		hub:               hub,
		csvRecordChann:    csvRecordChann,
		doneStreamChannel: doneStreamChannel,
		fileWriterPairs:   map[string]*fileWriterPair{},
	}
}

func (cfh *csvFileHandler) addWorker() {
	go func() {
		for csvRec := range cfh.csvRecordChann {
			fwPair, err := cfh.getFileWriterPairForStream(csvRec.streamName)
			if err != nil {
				cfh.hub.GetErrorChannel() <- err
				continue
			}

			err = fwPair.writer.Write(csvRec.data)
			if err != nil {
				cfh.hub.GetErrorChannel() <- err
				continue
			}
		}

		cfh.closeAndFlush()
	}()
}

func (cfh *csvFileHandler) getFileWriterPairForStream(streamName string) (*fileWriterPair, error) {
	fwPair, created := cfh.fileWriterPairs[streamName]
	if created {
		return fwPair, nil
	}

	f, err := os.Create(fmt.Sprintf("%s.csv", streamName))
	if err != nil {
		return nil, err
	}

	w := csv.NewWriter(f)

	fwPair = &fileWriterPair{
		file:   f,
		writer: w,
	}

	cfh.fileWriterPairs[streamName] = fwPair

	return fwPair, nil
}

func (cfh *csvFileHandler) closeAndFlush() {
	for _, fileWriterPair := range cfh.fileWriterPairs {
		fileWriterPair.writer.Flush()
		fileWriterPair.file.Close()
	}

	cfh.doneStreamChannel <- true
}
