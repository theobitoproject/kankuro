# kankuro

**Kankuro** is the go-sdk/cdk to help build Airbyte connectors quickly in Golang.

**Kankuro** is built on top of the Airbyte go-sdk/cdk built by bitstrapped (https://github.com/bitstrapped/airbyte) and it is meant to help build connectors quickly in go

Words extracted from base package:
>This package abstracts away much of the "protocol" away from the user and lets them focus on biz logic
It focuses on developer efficiency and tries to be strongly typed as much as possible to help dev's move fast without mistakes.

<div style="width: 320px;">

![Kankuro](kankuro.jpg)

</div>

> Sasori...your strength came because of your soul, not in spite of it. You tried to erase it, to become a puppet yourself, but couldn't change completely. Now you've got your immortal body, but you've fallen, sunk to the level of the puppets you used to control. You were supposed to be a top class ninja puppeteer, not a worthless nobody who lets someone else pull the strings.

**Kankuro**, shinobi of Sunagakure, the second eldest of the Three Sand Siblings and talented puppeteer.

## Installation 

```
go get github.com/theobitoproject/kankuro
```

## Usage 

### Detailed Usage

1. Define a source by implementing the `Source` interface. 

```go
// Source is the only interface you need to define to create your source!
type Source interface {
	// Spec returns the input "form" spec needed for your source
	Spec(
		messenger protocol.Messenger,
		configParser protocol.ConfigParser,
	) (*protocol.ConnectorSpecification, error)
	// Check verifies the source - usually verify creds/connection etc.
	Check(
		messenger protocol.Messenger,
		configParser protocol.ConfigParser,
	) error
	// Discover returns the schema of the data you want to sync
	Discover(
		messenger protocol.Messenger,
		configParser protocol.ConfigParser,
	) (*protocol.Catalog, error)
	// Read will read the actual data from your source and use
	// tracker.Record(), tracker.State() and tracker.Log() to sync data
	// with airbyte/destinations
	// MessageTracker is thread-safe and so it is completely find to
	// spin off goroutines to sync your data (just don't forget your waitgroups :))
	// returning an error from this will cancel the sync and returning a nil
	// from this will successfully end the sync
	Read(
		configuredCat *protocol.ConfiguredCatalog,
		messenger protocol.Messenger,
		configParser protocol.ConfigParser,
	) error
}
```

2. Inside of main, pass your source into the sourcerunner

```go
func main() {
	fsrc := filesource.NewFileSource("foobar.txt")
	runner := kankuro.NewSourceRunner(fsrc)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
```


3. Write a dockerfile (sample below)

```dockerfile
FROM golang:1.17-buster as build

WORKDIR /base
ADD . /base/
RUN go build -o /base/app .


LABEL io.airbyte.version=0.0.1
LABEL io.airbyte.name=airbyte/source

ENTRYPOINT ["/base/app"]
```

Check the [Random API source example](https://github.com/theobitoproject/airbyte_source_random_api) for a better usage example

### Contributors

I really thank [Bitstrapped connector](https://github.com/bitstrapped/airbyte) and contributors since this connector is inspired on their work
