package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"time"
	"xx/global"
)

var (
	// 日志文件配置
	logFilePath   string // 日志文件路径
	logNamePrefix string // 日志文件名称前缀
	logNameFormat string // 日志文件名称格式
	logFileExt    string // 日志文件扩展名称

	MaxSize    int  // 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int  // 保留旧文件的最大个数
	MaxAge     int  // 保留旧文件的最大天数
	Compress   bool // 是否压缩/归档旧文件

	hook io.Writer
)

func InitLogger() {
	logFilePath = global.ServerConfig.LogInfo.LogPath
	logNamePrefix = global.ServerConfig.AppName
	logNameFormat = global.ServerConfig.LogInfo.LogFormat
	logFileExt = global.ServerConfig.LogInfo.LogFileExt

	MaxSize = global.ServerConfig.LogInfo.LogMaxSize
	MaxBackups = global.ServerConfig.LogInfo.LogMaxBackups
	MaxAge = global.ServerConfig.LogInfo.LogMaxAge
	Compress = global.ServerConfig.LogInfo.LogCompress

	// 自定义 zap 配置
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)

	logger := zap.New(core)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger) // flushes buffer, if any

	global.SugarLogger = logger.WithOptions(zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogFileFullPath() string {
	return fmt.Sprintf("%s/%s-%s.%s", logFilePath, logNamePrefix, time.Now().Format(logNameFormat), logFileExt)
}

func getLogWriter() zapcore.WriteSyncer {

	hook = &lumberjack.Logger{
		Filename:   getLogFileFullPath(),
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
		Compress:   Compress,
	}

	return zapcore.AddSync(hook)
}

func openLogFile() *os.File {
	_, err := os.Stat(logFilePath)
	switch {
	case os.IsNotExist(err):
		makeLogDir(logFilePath)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	logFileFullPath := getLogFileFullPath()
	logFile, err := os.OpenFile(logFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return logFile
}

func makeLogDir(logFilePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+logFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func SimpleHttpGet(str string) {
	global.SugarLogger.Debugf("%s", str)
}
