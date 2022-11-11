package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/theobitoproject/kankuro/pkg/destination"
)

func main() {
	now := time.Now()

	dst := newDestinationCsv(".") // "." means current directory
	runner := destination.NewSafeDestinationRunner(dst, os.Stdout, os.Stdin, os.Args)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(now))
}
