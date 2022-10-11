package protocol

// State is used to store data between syncs - useful for incremental syncs and state storage
type State struct {
	Data *StateData `json:"data"`
}

type StateData map[string]interface{}
