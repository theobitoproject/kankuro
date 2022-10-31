package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/theobitoproject/kankuro/pkg/messenger"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

type csvRecordMarshaler struct {
	hub               messenger.ChannelHub
	csvRecordChann    csvRecordChannel
	doneStreamChannel chan bool
	workersAmount     int
}

func newCsvRecordMarshaler(
	hub messenger.ChannelHub,
	csvRecordChann csvRecordChannel,
	doneStreamChannel chan bool,
) *csvRecordMarshaler {
	return &csvRecordMarshaler{
		hub:               hub,
		csvRecordChann:    csvRecordChann,
		doneStreamChannel: doneStreamChannel,
		workersAmount:     0,
	}
}

func (rp *csvRecordMarshaler) addWorker() {
	rp.workersAmount++

	go func() {
		for rec := range rp.hub.GetRecordChannel() {
			csvRec, err := marshal(rec)
			if err != nil {
				rp.hub.GetErrorChannel() <- err
				continue
			}

			rp.csvRecordChann <- csvRec
		}

		rp.workersAmount--

		if rp.workersAmount == 0 {
			// close(rp.csvRecordChann)
			rp.doneStreamChannel <- true
		}
	}()
}

func marshal(rec *protocol.Record) (*csvRecord, error) {
	csvRec := &csvRecord{
		streamName: rec.Stream,
	}

	for _, v := range *rec.Data {

		str, err := convertToString(v)
		if err != nil {
			panic(err)
		}

		csvRec.data = append(csvRec.data, str)
	}

	return csvRec, nil
}

func convertToString(v interface{}) (string, error) {
	switch assert := v.(type) {
	case string:
		return assert, nil

	case int:
		return strconv.Itoa(assert), nil

	case float64:
		return fmt.Sprintf("%f", assert), nil

	case []interface{}:
		strValues := []string{}

		for _, a := range assert {
			sa, err := convertToString(a)
			if err != nil {
				return "", err
			}

			strValues = append(strValues, sa)
		}

		bt, err := json.Marshal(strValues)
		if err != nil {
			return "", err
		}

		return string(bt), nil

	case map[string]interface{}:
		strValues := []string{}

		for _, a := range assert {
			sa, err := convertToString(a)
			if err != nil {
				return "", err
			}

			strValues = append(strValues, sa)
		}

		bt, err := json.Marshal(strValues)
		if err != nil {
			return "", err
		}

		return string(bt), nil

	default:
		return "", fmt.Errorf("type not supported for %v", v)
	}
}
