package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var SugarLogger *zap.SugaredLogger
// func InitializeLogger() {
// 	logger, _ := zap.NewProduction()
// 	SugarLogger = logger.Sugar()
// }
func InitializeLogger() {
    writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
    // Time function can be customized
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // Log name
		MaxSize:    1,            // File content size, MB
		MaxBackups: 5,            // Maximum number of old files retained
		MaxAge:     30,           // Maximum number of days to keep old files
		Compress:   false,        // Is the file compressed
	}
	return zapcore.AddSync(lumberJackLogger)
}
