package protocol

const (
	// CmdSpec references spec command which declares the user-provided credentials
	// or configuration needed to run the connector
	CmdSpec Cmd = "spec"
	// CmdCheck references check command which tests if with the user-provided configuration
	// the connector can connect with the underlying data source
	CmdCheck Cmd = "check"
	// CmdDiscover references discover command which declares the different streams of data
	// that this connector can output
	CmdDiscover Cmd = "discover"
	// CmdRead references read command which reads data from the underlying data source
	CmdRead Cmd = "read"
	// CmdWrite references write command which writes data to the underlying destination
	CmdWrite Cmd = "write"
)

// Cmd defines the specific name of the command run for the connector
type Cmd string

// IsZero if the instance is empty
func (c *Cmd) IsZero() bool {
	return *c == ""
}
