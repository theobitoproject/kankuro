package main

import (
	"log"
	"os"

	"github.com/theobitoproject/kankuro"
)

func main() {
	url := "https://random-data-api.com/api/v2/beers"
	src := NewRandomAPISource(url)

	runner := kankuro.NewSourceRunner(src, os.Stdout)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
