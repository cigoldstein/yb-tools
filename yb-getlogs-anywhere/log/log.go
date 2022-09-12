package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"yb-get/structs"
)

var Logger structs.Logger

func CreateLogger(debugFlag bool, verboseFlag bool) zap.SugaredLogger {

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var defaultConsoleLogLevel zapcore.Level

	// configure verbosity and log-level
	switch {
	case debugFlag:
		defaultConsoleLogLevel = zap.DebugLevel
		config.FunctionKey = "true"
	case verboseFlag:
		defaultConsoleLogLevel = zap.InfoLevel
		config.FunctionKey = "true"
	default:
		defaultConsoleLogLevel = zap.InfoLevel
		config.TimeKey = ""
		config.CallerKey = ""
	}

	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultConsoleLogLevel)
	zapNew := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.FatalLevel))
	return *zapNew.Sugar()
}
