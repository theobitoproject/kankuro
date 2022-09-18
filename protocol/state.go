package protocol

// state is used to store data between syncs - useful for incremental syncs and state storage
type state struct {
	Data interface{} `json:"data"`
}
