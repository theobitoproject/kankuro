package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/theobitoproject/kankuro/pkg/destination"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

type messages struct {
	Data []protocol.AirbyteMessage `json:"data"`
}

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		data, err := os.ReadFile("data.json")
		if err != nil {
			panic(err)
		}

		var msgs messages

		err = json.Unmarshal(data, &msgs)
		if err != nil {
			panic(err)
		}

		for _, msg := range msgs.Data {
			err = json.NewEncoder(w).Encode(msg)
			if err != nil {
				fmt.Println("error logging messages")
				break
			}
		}

		err = w.Close()
		if err != nil {
			panic(err)
		}
	}()

	exampleCSV := newCsvDestination()
	runner := destination.NewSafeDestinationRunner(exampleCSV, w, r, os.Args)
	err = runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
