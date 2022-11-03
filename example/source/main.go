package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/theobitoproject/kankuro/pkg/source"
)

func main() {
	now := time.Now()

	src := newSourceRandomAPI("https://random-data-api.com/api/v2")
	runner := source.NewSafeSourceRunner(src, os.Stdout, os.Args)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(now))
}
