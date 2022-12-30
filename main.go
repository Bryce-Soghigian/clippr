package main

import (
	"clippr/cmd"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	level := zap.DebugLevel
	if os.Getenv("LOGGING_LEVEL_DEBUG") == "true" {
		level = zap.DebugLevel
	}
	atom := zap.NewAtomicLevelAt(level)

	return zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
}
func main() {
	logger := setupLogger()
	cmd.Execute(logger)
	os.Exit(0)
}
