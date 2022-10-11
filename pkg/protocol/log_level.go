package protocol

const (
	// LogLevelFatal defines the level of log for fatal
	LogLevelFatal LogLevel = "FATAL"
	// LogLevelError defines the level of log for error
	LogLevelError LogLevel = "ERROR"
	// LogLevelWarn defines the level of log for warning
	LogLevelWarn LogLevel = "WARN"
	// LogLevelInfo defines the level of log for information
	LogLevelInfo LogLevel = "INFO"
	// LogLevelDebug defines the level of log for debugging
	LogLevelDebug LogLevel = "DEBUG"
	// LogLevelTrace defines the level of log for traces
	LogLevelTrace LogLevel = "TRACE"
)

// LogLevel defines the log levels that can be emitted with airbyte logs
type LogLevel string
