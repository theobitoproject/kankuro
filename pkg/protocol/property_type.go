package protocol

const (
	// String defines string property type
	String PropType = "string"
	// Number defines number property type
	Number PropType = "number"
	// Integer defines integer property type
	Integer PropType = "integer"
	// Boolean defines boolean property type
	Boolean PropType = "boolean"
	// Object defines object property type
	Object PropType = "object"
	// Array defines array property type
	Array PropType = "array"
	// Null defines null property type
	Null PropType = "null"
)

// PropType defines the property types any field can take. See more here:  https://docs.airbyte.com/understanding-airbyte/supported-data-types
type PropType string
