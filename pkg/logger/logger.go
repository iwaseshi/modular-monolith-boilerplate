package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
)

type loggerKey struct{}

func getLogger() *zap.Logger {
	if instance == nil {
		instance = newLogger(zap.DebugLevel)
	}
	return instance
}

func newLogger(severity zapcore.Level) (logger *zap.Logger) {
	// fallback
	defer func() {
		if err := recover(); err != nil {
			// errorの場合も成功時と呼び出し元を合わせるためスタックをスキップする設定を追加する。
			logger, err = zap.NewProduction(zap.AddCallerSkip(1))
			if err != nil {
				panic(err)
			}
		}
	}()

	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(severity),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			TimeKey:    "timestamp",
			LevelKey:   "severity",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
				switch level {
				case zapcore.DebugLevel:
					encoder.AppendString("DEBUG")
				case zapcore.InfoLevel:
					encoder.AppendString("INFO")
				case zapcore.WarnLevel:
					encoder.AppendString("WARNING")
				case zapcore.ErrorLevel:
					encoder.AppendString("ERROR")
				case zapcore.FatalLevel:
					encoder.AppendString("FATAL")
				}
			},
		},
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	// LoggerをWrapして利用しているため、呼び出し元のスタックを1つスキップする設定を追加する。
	return logger.WithOptions(zap.AddCallerSkip(1))
}

// デフォルトのloggerを取得
// Contextを使用できないところのみ使用する
func Default() (defaultLogger *logger) {
	// fallback
	defer func() {
		if err := recover(); err != nil {
			defaultLogger = &logger{getLogger().Sugar()}
		}
	}()
	return &logger{
		logger: getLogger().Sugar(),
	}
}

func RegisterInCtx(ctx context.Context, fields ...zap.Field) (newCtx context.Context) {
	// fallback
	defer func() {
		if err := recover(); err != nil {
			newCtx = context.WithValue(ctx, loggerKey{}, &logger{
				logger: getLogger().With(fields...).Sugar(),
			})
		}
	}()

	return context.WithValue(ctx, loggerKey{}, &logger{
		logger: getLogger().With(fields...).Sugar(),
	})
}

// contextからloggerを取得
func WithCtx(ctx context.Context) *logger {
	if ctx == nil {
		return Default()
	}
	if WithCtx, acquired := ctx.Value(loggerKey{}).(*logger); acquired {
		return WithCtx
	}
	return Default()
}

type logger struct {
	logger *zap.SugaredLogger
}

func (l *logger) Debug(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *logger) Info(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *logger) Warning(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *logger) Error(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *logger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
	_ = l.Sync()
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
