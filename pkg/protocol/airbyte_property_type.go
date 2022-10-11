package protocol

const (
	// TimestampWithTZ defines the airbyte specific property type for timestamps with zone
	TimestampWithTZ AirbytePropType = "timestamp_with_timezone"
	// TimestampWOTZ defines the airbyte specific property type for timestamps without zone
	TimestampWOTZ AirbytePropType = "timestamp_without_timezone"
	// BigInteger defines the airbyte specific property type for big integers
	BigInteger AirbytePropType = "big_integer"
	// BigNumber defines the airbyte specific property type for big numbers
	BigNumber AirbytePropType = "big_number"
)

// AirbytePropType is used to define airbyte specific property types.
// See more here: https://docs.airbyte.com/understanding-airbyte/supported-data-types
type AirbytePropType string
