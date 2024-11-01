package config

type Logger interface {
	GetLoggerLevel() int
	GetLoggerEncodingMethod() int
	GetLoggerEncodingCaller() bool
	GetLoggerPath() string
}

func (l logger) GetLoggerLevel() int {
	return l.Level

}

func (l logger) GetLoggerEncodingMethod() int {
	return l.Encoding.Method

}

func (l logger) GetLoggerEncodingCaller() bool {
	return l.Encoding.Caller

}

func (l logger) GetLoggerPath() string {
	return l.Path

}
