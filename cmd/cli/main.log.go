package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger := zap.NewExample()
	// logger.Info("hello", zap.String("tên: ", "Tú"))

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello development", zap.Int("tuổi", 22))

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello production", zap.Bool("nam?", true))

	encoder := getEncoderLog()
	sync := getWriter()

	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core)

	logger.Info("hello world")
	logger.Error("Lỗi nè")
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

func getWriter() zapcore.WriteSyncer {
	timeNow := time.Now().Format("22-12-2006")
	filePath := fmt.Sprintf("./logs/log-%s.txt", timeNow)
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	syncFile := zapcore.AddSync(file)

	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
