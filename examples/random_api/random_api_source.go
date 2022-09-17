package main

import "github.com/theobitoproject/kankuro"

// RandomAPISource defines the source from which data will come
// from random api platform
// See: https://random-data-api.com/
type RandomAPISource struct {
	url string
}

// NewRandomAPISource creates a new instance of RandomAPISource
func NewRandomAPISource(url string) kankuro.Source {
	return RandomAPISource{url}
}

// Spec returns the input "form" spec needed for your source
func (s RandomAPISource) Spec(
	logTracker kankuro.LogTracker,
) (*kankuro.ConnectorSpecification, error) {
	return &kankuro.ConnectorSpecification{
		DocumentationURL:      "https://random-data-api.com/",
		ChangeLogURL:          "https://random-data-api.com/",
		SupportsIncremental:   false,
		SupportsNormalization: true,
		SupportsDBT:           true,
		SupportedDestinationSyncModes: []kankuro.DestinationSyncMode{
			kankuro.DestinationSyncModeOverwrite,
		},
		ConnectionSpecification: kankuro.ConnectionSpecification{
			Title:       "Random Data API",
			Description: "Random Data Source API",
			Type:        "object",
			Required:    []kankuro.PropertyName{"limit"},
			Properties: kankuro.Properties{
				Properties: map[kankuro.PropertyName]kankuro.PropertySpec{
					"limit": {
						Description: "max number of element to pull per instance",
						Examples:    []string{"1", "5"},
						PropertyType: kankuro.PropertyType{
							Type: []kankuro.PropType{
								kankuro.Integer,
							},
						},
					},
				},
			},
		},
	}, nil
}

// Check verifies the source - usually verify creds/connection etc.
func (s RandomAPISource) Check(
	srcCfgPath string,
	logTracker kankuro.LogTracker,
) error {
	return nil
}

// Discover returns the schema of the data you want to sync
func (s RandomAPISource) Discover(
	srcConfigPath string,
	logTracker kankuro.LogTracker,
) (*kankuro.Catalog, error) {
	return nil, nil
}

// Read will read the actual data from your source and use
// tracker.Record(), tracker.State() and tracker.Log() to sync data
// with airbyte/destinations
// MessageTracker is thread-safe and so it is completely find to
// spin off goroutines to sync your data (just don't forget your waitgroups :))
// returning an error from this will cancel the sync and returning a nil
// from this will successfully end the sync
func (s RandomAPISource) Read(
	sourceCfgPath string,
	prevStatePath string,
	configuredCat *kankuro.ConfiguredCatalog,
	tracker kankuro.MessageTracker,
) error {
	return nil
}
