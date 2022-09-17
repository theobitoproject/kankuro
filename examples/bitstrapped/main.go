package main

import (
	"log"
	"os"

	"github.com/theobitoproject/kankuro"
	"github.com/theobitoproject/kankuro/examples/bitstrapped/apisource"
)

func main() {
	hsrc := apisource.NewAPISource("https://api.bitstrapped.com")
	runner := kankuro.NewSourceRunner(hsrc, os.Stdout)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
