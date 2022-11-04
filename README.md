# kankuro

**Kankuro** is the go-sdk/cdk to help build Airbyte connectors quickly in Golang.

**Kankuro** is based on the Airbyte go-sdk/cdk built by Bitstrapped ([https://github.com/bitstrapped/airbyte](https://github.com/bitstrapped/airbyte))

Words extracted from Bitstrapped package:
>This package abstracts away much of the "protocol" away from the user and lets them focus on biz logic
It focuses on developer efficiency and tries to be strongly typed as much as possible to help dev's move fast without mistakes.

<div style="width: 320px;">

![Kankuro](kankuro.jpg)

</div>

> Sasori...your strength came because of your soul, not in spite of it. You tried to erase it, to become a puppet yourself, but couldn't change completely. Now you've got your immortal body, but you've fallen, sunk to the level of the puppets you used to control. You were supposed to be a top class ninja puppeteer, not a worthless nobody who lets someone else pull the strings.

**Kankuro**, shinobi of Sunagakure, the second eldest of the Three Sand Siblings and talented puppeteer.

## Contributors

I really thank [Bitstrapped connector](https://github.com/bitstrapped/airbyte) and contributors since this connector is inspired on their work

## Installation

```
go get github.com/theobitoproject/kankuro
```

## Usage

### 1. Basic Interface

#### 1.1. MessageWriter

**MessageWriter** allows to write only two types of messages: **log** and **state**. Most likely, log messages would be the only thing that is needed while building a connector.

```go
import "github.com/theobitoproject/kankuro/pkg/messenger"

func (c *MyCustomConnector) Check(
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
) error {
	mw.WriteLog(protocol.LogLevelInfo, "running read...")
}
```

#### 1.2. ConfigParser

**ConfigParser** allows to get the configuration values set by the connector. To use this, it's necessary to define a struct that defines the expected configuration by the connector.

```go
import "github.com/theobitoproject/kankuro/pkg/messenger"

type connectorConfiguration struct {
	User int `json:"user"`
	Password int `json:"password"`
}

func (c *MyCustomConnector) Check(
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
) error {
	var cc connectorConfiguration
	err = cp.UnmarshalConfigPath(&cc)
	if err != nil {
		return err
	}

	// cc.User and cc.Password ready to use
}
```

#### 1.3. ConfiguredCatalog

**ConfiguredCatalog** contains all the streams that are going to be extract and loaded during sync of a single connection. This is the base to start extracting data.

```go
func (c *MyCustomConnector) Read(
	cc *protocol.ConfiguredCatalog,
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
	hub messenger.ChannelHub,
) {
	for _, stream := range cc.Streams {
		// handle each stream
	}
}
```

#### 1.4. ChannelHub

**Kankuro** uses channels and goroutines in an attempt to make connectors more efficient, which could be translated in faster connections. **ChannelHub** allows to share **Records** and **Errors**.

```go
func (c *MyCustomConnector) Read(
	cc *protocol.ConfiguredCatalog,
	mw messenger.MessageWriter,
	cp messenger.ConfigParser,
	hub messenger.ChannelHub,
) {
	rcds, err := c.fetchData()
	if err != nil {
		// send errors to the error channel
		se.hub.GetErrorChannel() <- err
		// handle execution according to the context
		return
	}

	for _, rcd := range rcds {
		rec := &protocol.Record{
			Namespace: configuredStream.Stream.Namespace,
			Data:      rcd,
			Stream:    configuredStream.Stream.Name,
		}

		// send records to the record channel
		se.hub.GetRecordChannel() <- rec
	}


	// IMPORTANT: close channels from hub at the end of the execution

	// NOTE: record channel should be closed only when building a source connector
	// this channel should NOT be closed by a destination connector
	close(hub.GetRecordChannel())
	close(hub.GetErrorChannel())
}

func (c *MyCustomConnector) fetchData() ([]*protocol.RecordData, error) {
	// fetch data from source
	return []*protocol.RecordData{}, nil
}
```

### 2. Source Connector

#### 2.1. Define an implementation following the [`Source`](https://github.com/theobitoproject/kankuro/blob/main/pkg/source/source.go) interface.

#### 2.2. Set up `main.go`

Inside of main, pass your source into the **Source Runner**

```go
package main

import "github.com/theobitoproject/kankuro/pkg/source"

func main() {
	src := NewCustomSource()
	runner := source.NewSafeSourceRunner(src, os.Stdout, os.Args)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 2.3. Source Examples

Inside this repository there's a [working example](https://github.com/theobitoproject/kankuro/tree/main/example/source) that could be tested locally.

```sh
cd example/source/
make read
```

The above command will print results in your terminal.

Also, check out [Source Random API](https://github.com/theobitoproject/airbyte_source_random_api) connector. This is an example on how it should work for projects.

### 3. Destination Connector

#### 3.1. Define an implementation following the [`Destination`](https://github.com/theobitoproject/kankuro/blob/main/pkg/destination/destination.go) interface.

#### 2.2. Set up `main.go`

Inside of main, pass your destination into the **Destination Runner**

```go
package main

import "github.com/theobitoproject/kankuro/pkg/destination"

func main() {
	dst := NewCustomDestination()
	runner := destination.NewSafeDestinationRunner(dst, os.Stdout, os.Stdin, os.Args)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 2.3. Destination Examples

Inside this repository there's a [working example](https://github.com/theobitoproject/kankuro/tree/main/example/destination) that could be tested locally.

```sh
cd example/destination/
make write
```

The above command will create csv files inside the `example/destination/destination` directory.

Also, check out [Destination CSV](https://github.com/theobitoproject/airbyte_destination_csv) connector. This is an example on how it should work for projects.

### 4. Write a dockerfile

```dockerfile
FROM golang:1.19-buster as build

WORKDIR /base
ADD . /base/
RUN go build -o /base/app .


LABEL io.airbyte.version=0.0.1
LABEL io.airbyte.name=airbyte/source

ENTRYPOINT ["/base/app"]
```

**NOTE:** The above sample could be modified according to the context.
