package protocol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConfigParser reads the command that runs connectors
// and returns isolated parameters
type ConfigParser interface {
	// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
	GetMainCommand() (Cmd, error)
	// UnmarshalSourceConfigPath unmarshals source config json file into a struct
	UnmarshalSourceConfigPath(v interface{}) error
	// UnmarshalSourceConfigPath unmarshals state json file into a struct
	UnmarshalStatePath(v interface{}) error
	// UnmarshalSourceConfigPath unmarshals catalog json file into a struct
	UnmarshalCatalogPath(v interface{}) error
}

type configParser struct {
	args []string
}

// NewConfigParser creates an instance of ConfigParser
func NewConfigParser(args []string) ConfigParser {
	return configParser{args}
}

// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
func (cp configParser) GetMainCommand() (Cmd, error) {
	if len(cp.args) <= 1 {
		return "", fmt.Errorf("main command not found")
	}
	return Cmd(cp.args[1]), nil
}

// UnmarshalSourceConfigPath unmarshals source config json file into a struct
func (cp configParser) UnmarshalSourceConfigPath(v interface{}) error {
	if cp.args[2] != "--config" {
		return fmt.Errorf("expect --config")
	}
	return cp.unmarshalFromPath(cp.args[3], v)
}

// UnmarshalStatePath unmarshals source config json file into a struct
func (cp configParser) UnmarshalStatePath(v interface{}) error {
	if len(cp.args) <= 6 {
		return nil
	}
	if cp.args[6] != "--state" {
		return fmt.Errorf("expect --state")
	}
	return cp.unmarshalFromPath(cp.args[7], v)
}

// UnmarshalCatalogPath unmarshals source config json file into a struct
func (cp configParser) UnmarshalCatalogPath(v interface{}) error {
	if cp.args[4] != "--catalog" {
		return fmt.Errorf("expect --catalog")
	}
	return cp.unmarshalFromPath(cp.args[5], v)
}

// UnmarshalFromPath is used to unmarshal json files into respective struct's
// this is most commonly used to unmarshal your State between
// runs and also unmarshal SourceConfig's
func (cp configParser) unmarshalFromPath(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
