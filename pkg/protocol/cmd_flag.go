package protocol

const (
	// CatalogCmdFlag defines the flag name for catalog
	CatalogCmdFlag CmdFlag = "--catalog"
	// ConfigCmdFlag defines the flag name for config
	ConfigCmdFlag CmdFlag = "--config"
	// StateCmdFlag defines the flag name for state
	StateCmdFlag CmdFlag = "--state"
)

// CmdFlag defines the specific flags that could be used
// when airebyte runs a main command
type CmdFlag string
