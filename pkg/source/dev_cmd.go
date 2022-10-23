package source

import "github.com/theobitoproject/kankuro/pkg/protocol"

// Note:
// All commands defined here should only be used for development stages

const (
	// CmdPrintCatalog references print-catalog command which takes the catalog from a source
	// and prints it out inside a json file
	CmdPrintCatalog protocol.Cmd = "print-catalog"
)
