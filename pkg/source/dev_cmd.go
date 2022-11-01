package source

import "github.com/theobitoproject/kankuro/pkg/protocol"

// Note:
// All commands defined here should only be used for development stages

const (
	// CmdPrintConfiguredCatalog references print-configured-catalog command
	// which takes the configured catalog from a source and prints it out inside a json file
	CmdPrintConfiguredCatalog protocol.Cmd = "print-configured-catalog"
)
