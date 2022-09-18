package protocol

const (
	// Date defines the type of format for dates
	Date FormatType = "date"
	// Date defines the type of format for date-times
	DateTime FormatType = "datetime"
)

// FormatType is used to define data type formats supported by airbyte where needed
// (usually for strings formatted as dates). See more here: https://docs.airbyte.com/understanding-airbyte/supported-data-types
type FormatType string
