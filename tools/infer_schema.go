package kankuro

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/theobitoproject/kankuro/protocol"
	"github.com/theobitoproject/kankuro/schema"
)

// Infer schema translates golang structs to JSONSchema format
func InferSchemaFromStruct(i interface{}, messenger protocol.Messenger) protocol.Properties {
	var prop protocol.Properties

	s, err := schema.Generate(reflect.TypeOf(i))
	if err != nil {
		messenger.WriteLog(protocol.LogLevelError, fmt.Sprintf("generate schema error: %v", err))
		return prop
	}

	b, err := json.Marshal(s)
	if err != nil {
		messenger.WriteLog(protocol.LogLevelError, fmt.Sprintf("json marshal schema error: %v", err))
		return prop
	}

	err = json.Unmarshal(b, &prop)
	if err != nil {
		messenger.WriteLog(protocol.LogLevelError, fmt.Sprintf("unmarshal schema to propspec error: %v", err))
		return prop
	}

	return prop
}
