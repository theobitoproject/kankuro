package protocol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CommandParser reads the command that runs connectors
// and returns isolated parameters
type CommandParser interface {
	// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
	GetMainCommand() (Cmd, error)
	// GetSourceConfigPath returns the source config path
	// from the commmand used to run the connector
	GetSourceConfigPath() (string, error)
	// GetStatePath returns the state path
	// from the commmand used to run the connector
	GetStatePath() (string, error)
	// GetCatalogPath returns the catalog path
	// from the commmand used to run the connector
	GetCatalogPath() (string, error)
	// UnmarshalFromPath is used to unmarshal json files into respective struct's
	// this is most commonly used to unmarshal your State between
	// runs and also unmarshal SourceConfig's
	UnmarshalFromPath(path string, v interface{}) error
}

type commandParser struct {
	args []string
}

// NewCommandParser creates an instance of CommandParser
func NewCommandParser(args []string) CommandParser {
	return commandParser{args}
}

// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
func (cp commandParser) GetMainCommand() (Cmd, error) {
	if len(cp.args) <= 1 {
		return "", fmt.Errorf("main command not found")
	}
	return Cmd(cp.args[1]), nil
}

// GetSourceConfigPath returns the source config path
// from the commmand used to run the connector
func (cp commandParser) GetSourceConfigPath() (string, error) {
	if cp.args[2] != "--config" {
		return "", fmt.Errorf("expect --config")
	}
	return cp.args[3], nil
}

// GetStatePath returns the state path
// from the commmand used to run the connector
func (cp commandParser) GetStatePath() (string, error) {
	if len(cp.args) <= 6 {
		return "", nil
	}
	if cp.args[6] != "--state" {
		return "", fmt.Errorf("expect --state")
	}
	return cp.args[7], nil
}

// GetCatalogPath returns the catalog path
// from the commmand used to run the connector
func (cp commandParser) GetCatalogPath() (string, error) {
	if cp.args[4] != "--catalog" {
		return "", fmt.Errorf("expect --catalog")
	}
	return cp.args[5], nil
}

// UnmarshalFromPath is used to unmarshal json files into respective struct's
// this is most commonly used to unmarshal your State between
// runs and also unmarshal SourceConfig's
func (cp commandParser) UnmarshalFromPath(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
