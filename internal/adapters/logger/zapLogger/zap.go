package zapLogger

import (
	"assetio/internal/port"
	"context"
	"fmt"
	"runtime"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ENCODING_TYPE_JSON    = 1
	ENCODING_TYPE_CONSOLE = 2
)

type Config struct {
	Level          int
	Encoding       int
	EncodingCaller bool
	OutputPath     string
}

type zapLog struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// New creates a new instance of zapLog using the provided config for logging.
// It sets up the logger with a specified log level, encoder configuration, and output path.
// It returns a port.Logger instance and any error encountered.
func New(config Config) (port.Logger, error) {

	// Convert the provided log level to a zapcore.Level
	zapLevel, err := intToZapLevel(config.Level)
	if err != nil {
		// Return error if the log level is invalid
		return nil, err
	}

	// Define the encoder configuration for structuring log output
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// If configured, use a custom encoder for the caller's information
	if config.EncodingCaller {
		encoderConfig.EncodeCaller = callerEncoder
	}

	// Set up the general core for logging with JSON encoding and output to a log file
	generalCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.OutputPath,
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     28,
		}),
		zap.NewAtomicLevelAt(zapLevel),
	)

	// Combine the general core with any other cores, if necessary
	core := zapcore.NewTee(generalCore)

	// Create the logger instance with caller and stacktrace options for errors
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Return the custom zapLog instance wrapped around the logger
	return &zapLog{
		logger: logger,
		sugar:  logger.Sugar(),
	}, nil
}

// Debug logs a debug-level message with the provided context and messages
func (l *zapLog) Debug(ctx context.Context, messages ...any) {
	l.sugar.Debug(messages...)
}

// Info logs an info-level message with the provided context and messages
func (l *zapLog) Info(ctx context.Context, messages ...any) {
	l.sugar.Info(messages...)
}

// Warn logs a warn-level message with the provided context and messages
func (l *zapLog) Warn(ctx context.Context, messages ...any) {
	l.sugar.Warn(messages...)
}

// Error logs an error-level message with the provided context and messages
func (l *zapLog) Error(ctx context.Context, messages ...any) {
	l.sugar.Error(messages...)
}

// Fatal logs a fatal-level message and terminates the program with the provided context and messages
func (l *zapLog) Fatal(ctx context.Context, messages ...any) {
	l.sugar.Fatal(messages...)
}

// Debugf logs a formatted debug-level message with the provided context and arguments
func (l *zapLog) Debugf(ctx context.Context, template string, args ...any) {
	l.sugar.Debugf(template, args...)
}

// Infof logs a formatted info-level message with the provided context and arguments
func (l *zapLog) Infof(ctx context.Context, template string, args ...any) {
	l.sugar.Infof(template, args...)
}

// Warnf logs a formatted warn-level message with the provided context and arguments
func (l *zapLog) Warnf(ctx context.Context, template string, args ...any) {
	l.sugar.Warnf(template, args...)
}

// Errorf logs a formatted error-level message with the provided context and arguments
func (l *zapLog) Errorf(ctx context.Context, template string, args ...any) {
	l.sugar.Errorf(template, args...)
}

// Fatalf logs a formatted fatal-level message and terminates the program with the provided context and arguments
func (l *zapLog) Fatalf(ctx context.Context, template string, args ...any) {
	l.sugar.Fatalf(template, args...)
}

// Debugw logs a debug-level message with additional context (key-value pairs) and the provided context
func (l *zapLog) Debugw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Debugw(msg, keysAndValues...)
}

// Infow logs an info-level message with additional context (key-value pairs) and the provided context
func (l *zapLog) Infow(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Infow(msg, keysAndValues...)
}

// Warnw logs a warn-level message with additional context (key-value pairs) and the provided context
func (l *zapLog) Warnw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Warnw(msg, keysAndValues...)
}

// Errorw logs an error-level message with additional context (key-value pairs) and the provided context
func (l *zapLog) Errorw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Errorw(msg, keysAndValues...)
}

// Fatalw logs a fatal-level message with additional context (key-value pairs) and the provided context
func (l *zapLog) Fatalw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entries to their destination (i.e., output file) and synchronizes the logger
func (l *zapLog) Sync(ctx context.Context) error {
	return l.logger.Sync()
}

// intToZapLevel converts an integer log level to a zapcore.Level, or returns an error if the level is invalid
func intToZapLevel(level int) (zapcore.Level, error) {
	switch level {
	case 1:
		return zapcore.FatalLevel, nil
	case 2:
		return zapcore.ErrorLevel, nil
	case 3:
		return zapcore.WarnLevel, nil
	case 4:
		return zapcore.InfoLevel, nil
	case 5:
		return zapcore.DebugLevel, nil
	default:
		// Return error for invalid log level
		return zapcore.InfoLevel, fmt.Errorf("invalid log level: %d", level)
	}
}

// callerEncoder is a custom function to encode the caller information in the logs,
// including the file, line, and function name.
func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// Get the function name by skipping the current stack frame
	file, line, functionName := getFunctionName(8)
	// Append the formatted caller information to the log entry
	enc.AppendString(fmt.Sprintf("%s:%d %s", file, line, functionName))
}

// getFunctionName retrieves the file, line, and function name from the current stack trace,
// helping to add detailed caller information to logs.
func getFunctionName(skip int) (string, int, string) {
	// Retrieve the program counter, file, and line from the stack trace
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0, "unknown"
	}
	// Retrieve the function name from the program counter
	function := runtime.FuncForPC(pc)
	if function == nil {
		return "unknown", 0, "unknown"
	}
	// Return the file, line, and function name
	return file, line, function.Name()
}
