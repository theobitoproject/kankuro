package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/theobitoproject/kankuro/pkg/destination"
	"github.com/theobitoproject/kankuro/pkg/protocol"
)

type messages struct {
	Data []protocol.AirbyteMessage `json:"data"`
}

func main() {
	// this is a workaround to create:
	// - a writer to send messages
	// - a reader to catch messages
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	// again, this is a workaround
	// in the real world, a source will write messages
	go func() {
		writeMessages(w)
		err = w.Close()
		if err != nil {
			panic(err)
		}
	}()

	now := time.Now()

	dst := newDestinationCsv(".") // "." means current directory
	runner := destination.NewSafeDestinationRunner(dst, w, r, os.Args)
	err = runner.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(now))
}

func writeMessages(w io.Writer) {
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
}
