package messenger

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// ConfigParser reads the command that runs connectors
// and returns isolated parameters
type ConfigParser interface {
	// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
	GetMainCommand() (protocol.Cmd, error)
	// UnmarshalConfigPath unmarshals source config json file into a struct
	UnmarshalConfigPath(v interface{}) error
	// UnmarshalStatePath unmarshals state json file into a struct
	UnmarshalStatePath(v interface{}) error
	// UnmarshalCatalogPath unmarshals catalog json file into a struct
	UnmarshalCatalogPath(v interface{}) error
}

type configParser struct {
	args []string
}

// NewConfigParser creates an instance of ConfigParser
func NewConfigParser(args []string) ConfigParser {
	return &configParser{args}
}

// GetMainCommand returns the comand to run the connector (spec, check, discover, read)
func (cp *configParser) GetMainCommand() (protocol.Cmd, error) {
	if len(cp.args) <= 1 {
		return "", fmt.Errorf("main command not found")
	}
	return protocol.Cmd(cp.args[1]), nil
}

// UnmarshalConfigPath unmarshals source config json file into a struct
func (cp *configParser) UnmarshalConfigPath(v interface{}) error {
	path, err := cp.getFlagConfigValue(protocol.ConfigCmdFlag)
	if err != nil {
		return nil
	}

	return cp.unmarshalFromPath(path, v)
}

// UnmarshalStatePath unmarshals source config json file into a struct
func (cp *configParser) UnmarshalStatePath(v interface{}) error {
	path, err := cp.getFlagConfigValue(protocol.StateCmdFlag)
	if err != nil {
		return nil
	}

	return cp.unmarshalFromPath(path, v)
}

// UnmarshalCatalogPath unmarshals source config json file into a struct
func (cp *configParser) UnmarshalCatalogPath(v interface{}) error {
	path, err := cp.getFlagConfigValue(protocol.CatalogCmdFlag)
	if err != nil {
		return nil
	}

	return cp.unmarshalFromPath(path, v)
}

// UnmarshalFromPath is used to unmarshal json files into respective struct's
// this is most commonly used to unmarshal your State between
// runs and also unmarshal SourceConfig's
func (cp *configParser) unmarshalFromPath(path string, v interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func (cp *configParser) getFlagConfigValue(flag protocol.CmdFlag) (string, error) {
	for argIndex, argValue := range cp.args {
		if argValue != string(flag) {
			continue
		}

		return cp.args[argIndex+1], nil
	}

	return "", fmt.Errorf("flag was not found in arguments: %s", flag)
}
