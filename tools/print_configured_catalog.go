package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/theobitoproject/kankuro/pkg/protocol"
)

// PrintConfiguredCatalogOnFile creates a configured catalog from a catalog
// and prints it out inside a json file
func PrintConfiguredCatalogOnFile(
	catalog *protocol.Catalog,
	directory string,
	filename string,
) error {
	configuredStreams := []protocol.ConfiguredStream{}

	for _, stream := range catalog.Streams {
		configuredStreams = append(configuredStreams, protocol.ConfiguredStream{
			Stream:              stream,
			SyncMode:            protocol.SyncModeFullRefresh,
			DestinationSyncMode: protocol.DestinationSyncModeOverwrite,
		})
	}

	configuredCatalog := protocol.ConfiguredCatalog{
		Streams: configuredStreams,
	}

	jsonConfiguredCatalog, err := json.Marshal(configuredCatalog)
	if err != nil {
		return err
	}

	err = os.MkdirAll(directory, 0755)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", directory, filename)
	return os.WriteFile(path, jsonConfiguredCatalog, 0755)
}
