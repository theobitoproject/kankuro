package apisource

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/theobitoproject/kankuro"
)

type APISource struct {
	baseURL string
}

type LastSyncTime struct {
	Timestamp int64 `json:"timestamp"`
}

type HTTPConfig struct {
	APIKey string `json:"apiKey"`
}

func NewAPISource(baseURL string) kankuro.Source {
	return APISource{
		baseURL: baseURL,
	}
}

func (h APISource) Spec(logTracker kankuro.LogTracker) (*kankuro.ConnectorSpecification, error) {
	if err := logTracker.Log(kankuro.LogLevelInfo, "Running Spec"); err != nil {
		return nil, err
	}
	return &kankuro.ConnectorSpecification{
		DocumentationURL:      "https://bitstrapped.com",
		ChangeLogURL:          "https://bitstrapped.com",
		SupportsIncremental:   false,
		SupportsNormalization: true,
		SupportsDBT:           true,
		SupportedDestinationSyncModes: []kankuro.DestinationSyncMode{
			kankuro.DestinationSyncModeOverwrite,
		},
		ConnectionSpecification: kankuro.ConnectionSpecification{
			Title:       "Example HTTP Source",
			Description: "This is an example http source for the docs's",
			Type:        "object",
			Required:    []kankuro.PropertyName{"apiKey"},
			Properties: kankuro.Properties{
				Properties: map[kankuro.PropertyName]kankuro.PropertySpec{
					"apiKey": {
						Description: "api key to access http source, valid uuid",
						Examples:    []string{"xxxx-xxxx-xxxx-xxxx"},
						PropertyType: kankuro.PropertyType{
							Type: []kankuro.PropType{
								kankuro.String,
							},
						},
					},
				},
			},
		},
	}, nil
}

func (h APISource) Check(srcCfgPath string, logTracker kankuro.LogTracker) error {
	if err := logTracker.Log(kankuro.LogLevelDebug, "validating api connection"); err != nil {
		return err
	}
	var srcCfg HTTPConfig
	err := kankuro.UnmarshalFromPath(srcCfgPath, &srcCfg)
	if err != nil {
		return err
	}

	resp, err := http.Get(fmt.Sprintf("%s/ping?key=%s", h.baseURL, srcCfg.APIKey))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status")
	}

	return resp.Body.Close()
}

func (h APISource) Discover(srcCfgPath string, logTracker kankuro.LogTracker) (*kankuro.Catalog, error) {
	var srcCfg HTTPConfig
	err := kankuro.UnmarshalFromPath(srcCfgPath, &srcCfg)
	if err != nil {
		return nil, err
	}

	return &kankuro.Catalog{Streams: []kankuro.Stream{{
		Name: "users",
		JSONSchema: kankuro.Properties{
			Properties: map[kankuro.PropertyName]kankuro.PropertySpec{
				"userid": {
					PropertyType: kankuro.PropertyType{
						Type:        []kankuro.PropType{kankuro.Integer, kankuro.Null},
						AirbyteType: kankuro.BigInteger},
					Description: "user ID - see the big int",
				},
				"name": {
					PropertyType: kankuro.PropertyType{
						Type: []kankuro.PropType{kankuro.String, kankuro.Null},
					},
					Description: "user name",
				},
			},
		},
		SupportedSyncModes: []kankuro.SyncMode{
			kankuro.SyncModeFullRefresh,
		},
		SourceDefinedCursor: false,
		Namespace:           "bitstrapped",
	},
		{
			Name: "payments",
			JSONSchema: kankuro.Properties{
				Properties: map[kankuro.PropertyName]kankuro.PropertySpec{
					"userid": {
						PropertyType: kankuro.PropertyType{
							Type:        []kankuro.PropType{kankuro.Integer, kankuro.Null},
							AirbyteType: kankuro.BigInteger},
						Description: "user ID - see the big int",
					},
					"paymentAmount": {
						PropertyType: kankuro.PropertyType{
							Type: []kankuro.PropType{kankuro.Integer, kankuro.Null},
						},
						Description: "payment amount",
					},
				},
			},
			SupportedSyncModes: []kankuro.SyncMode{
				kankuro.SyncModeFullRefresh,
			},
			SourceDefinedCursor: false,
			Namespace:           "bitstrapped",
		},
	}}, nil
}

type User struct {
	UserID int64  `json:"userid"`
	Name   string `json:"name"`
}

type Payment struct {
	UserID        int64 `json:"userid"`
	PaymentAmount int64 `json:"paymentAmount"`
}

func (h APISource) Read(sourceCfgPath string, prevStatePath string, configuredCat *kankuro.ConfiguredCatalog,
	tracker kankuro.MessageTracker) error {
	if err := tracker.Log(kankuro.LogLevelInfo, "Running read"); err != nil {
		return err
	}
	var src HTTPConfig
	err := kankuro.UnmarshalFromPath(sourceCfgPath, &src)
	if err != nil {
		return err
	}

	// see if there is a last sync
	var st LastSyncTime
	_ = kankuro.UnmarshalFromPath(sourceCfgPath, &st)
	if st.Timestamp <= 0 {
		st.Timestamp = -1
	}

	for _, stream := range configuredCat.Streams {
		if stream.Stream.Name == "users" {
			var u []User
			uri := fmt.Sprintf("https://api.bistrapped.com/users?apiKey=%s", src.APIKey)
			if err := httpGet(uri, &u); err != nil {
				return err
			}

			for _, ur := range u {
				err := tracker.Record(ur, stream.Stream.Name, stream.Stream.Namespace)
				if err != nil {
					return err
				}
			}
		}

		if stream.Stream.Name == "payments" {
			var p []Payment
			uri := fmt.Sprintf("%s/payments?apiKey=%s", h.baseURL, src.APIKey)
			if err := httpGet(uri, &p); err != nil {
				return err
			}

			for _, py := range p {
				err := tracker.Record(py, stream.Stream.Name, stream.Stream.Namespace)
				if err != nil {
					return err
				}
			}
		}
	}

	return tracker.State(&LastSyncTime{
		Timestamp: time.Now().UnixMilli(),
	})
}

func httpGet(uri string, v interface{}) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
