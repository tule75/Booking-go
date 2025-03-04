package logger

import (
	"ecommerce_go/pkg/setting"
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(logSetting setting.LogSetting) *LoggerZap {
	timeNow := time.Now().Format("24-12-2006")
	filePath := fmt.Sprintf(logSetting.FileLogName, timeNow)
	logLevel := logSetting.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    logSetting.MaxSize,
		MaxBackups: logSetting.MaxBackups,
		MaxAge:     logSetting.MaxAge,
		Compress:   logSetting.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1740381840.0001862 to 2025-02-24T14:23:59.999+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts => Time
	encoderConfig.TimeKey = "Time"

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
