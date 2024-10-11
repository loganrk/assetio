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
	Level           int
	Encoding        int
	EncodingCaller  bool
	OutputPath      string
	ErrorOutputPath string
}

type zapLog struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func New(config Config) (port.Logger, error) {

	zapLevel, err := intToZapLevel(config.Level)
	if err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if config.EncodingCaller {
		encoderConfig.EncodeCaller = callerEncoder
	}

	generalCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.OutputPath,
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}),
		zap.NewAtomicLevelAt(zapLevel),
	)

	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.ErrorOutputPath,
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}),
		zap.NewAtomicLevelAt(zapcore.ErrorLevel),
	)

	core := zapcore.NewTee(generalCore, errorCore)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &zapLog{
		logger: logger,
		sugar:  logger.Sugar(),
	}, nil
}

func (l *zapLog) Debug(ctx context.Context, messages ...any) {
	l.sugar.Debug(messages...)
}

func (l *zapLog) Info(ctx context.Context, messages ...any) {
	l.sugar.Info(messages...)
}

func (l *zapLog) Warn(ctx context.Context, messages ...any) {
	l.sugar.Warn(messages...)
}

func (l *zapLog) Error(ctx context.Context, messages ...any) {
	l.sugar.Error(messages...)
}

func (l *zapLog) Fatal(ctx context.Context, messages ...any) {
	l.sugar.Fatal(messages...)
}

func (l *zapLog) Debugf(ctx context.Context, template string, args ...any) {
	l.sugar.Debugf(template, args...)
}

func (l *zapLog) Infof(ctx context.Context, template string, args ...any) {
	l.sugar.Infof(template, args...)
}

func (l *zapLog) Warnf(ctx context.Context, template string, args ...any) {
	l.sugar.Warnf(template, args...)
}

func (l *zapLog) Errorf(ctx context.Context, template string, args ...any) {
	l.sugar.Errorf(template, args...)
}

func (l *zapLog) Fatalf(ctx context.Context, template string, args ...any) {
	l.sugar.Fatalf(template, args...)
}

func (l *zapLog) Debugw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *zapLog) Infow(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *zapLog) Warnw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *zapLog) Errorw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *zapLog) Fatalw(ctx context.Context, msg string, keysAndValues ...any) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func (l *zapLog) Sync(ctx context.Context) error {
	return l.logger.Sync()
}

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
		return zapcore.InfoLevel, fmt.Errorf("invalid log level: %d", level)
	}
}

// Custom caller encoder to include function name
func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	file, line, functionName := getFunctionName(8) // Get the calling function's name
	enc.AppendString(fmt.Sprintf("%s:%d %s", file, line, functionName))
}

// Helper function to retrieve the function name
func getFunctionName(skip int) (string, int, string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0, "unknown"
	}
	function := runtime.FuncForPC(pc)
	if function == nil {
		return "unknown", 0, "unknown"
	}
	return file, line, function.Name()
}
