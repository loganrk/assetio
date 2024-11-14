package config

// Logger defines the interface for configuring and retrieving logger properties such as logging level,
// encoding method, whether to include caller information, and the log file path.
type Logger interface {
	// GetLoggerLevel returns the log level for the logger (e.g., DEBUG, INFO, ERROR).
	GetLoggerLevel() int

	// GetLoggerEncodingMethod returns the encoding method for the log messages (e.g., JSON, plaintext).
	GetLoggerEncodingMethod() int

	// GetLoggerEncodingCaller returns whether the logger should include the caller information (file name, line number).
	GetLoggerEncodingCaller() bool

	// GetLoggerPath returns the path where the log files will be stored.
	GetLoggerPath() string
}

// GetLoggerLevel returns the log level for the logger. The log level determines the severity of logs to be recorded.
func (l logger) GetLoggerLevel() int {
	return l.Level // Accesses the Level field from the logger struct and returns the log level
}

// GetLoggerEncodingMethod returns the encoding method for the logger's output.
// This can be used to choose between different formats (e.g., JSON, plaintext).
func (l logger) GetLoggerEncodingMethod() int {
	return l.Encoding.Method // Accesses the Encoding.Method field from the logger struct and returns the encoding method
}

// GetLoggerEncodingCaller returns whether the logger includes caller information in the log messages.
// This is useful for debugging, as it helps identify where the log messages are coming from.
func (l logger) GetLoggerEncodingCaller() bool {
	return l.Encoding.Caller // Accesses the Encoding.Caller field from the logger struct and returns the status (true/false)
}

// GetLoggerPath returns the file path where the log files are saved.
// This is used to configure the destination of log files.
func (l logger) GetLoggerPath() string {
	return l.Path // Accesses the Path field from the logger struct and returns the log file path
}
