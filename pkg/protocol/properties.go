package protocol

// Properties defines the property map which is used to define any single "field name" along with its specification
type Properties struct {
	Properties map[PropertyName]PropertySpec `json:"properties"`
}

// PropertyName is a alias for a string to make it clear to the user that the "key" in the map is the name of the property
type PropertyName string

// PropertySpec defines the specification for a single property
type PropertySpec struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PropertyType `json:",omitempty"`
	Examples     []string                      `json:"examples,omitempty"`
	Items        map[string]interface{}        `json:"items,omitempty"`
	Properties   map[PropertyName]PropertySpec `json:"properties,omitempty"`
	IsSecret     bool                          `json:"airbyte_secret,omitempty"`
	Order        int                           `json:"order,omitempty"`
}

// PropertyType defines the type of property for a property specification
type PropertyType struct {
	Type        []PropType      `json:"type,omitempty"`
	AirbyteType AirbytePropType `json:"airbyte_type,omitempty"`
}
