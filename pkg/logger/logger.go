package logger

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/knadh/koanf"
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	TEXT = "text"
	JSON = "json"
)
const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
	Fatal = "fatal"
)

func NewLogger(conf *koanf.Koanf) (*zap.SugaredLogger, error) {
	//logger, err := zap.NewProduction()
	//if err != nil {
	//	return nil, err
	//}
	logLevel := conf.String("log.level")
	logFormat := conf.String("log.format")
	logPath := conf.String("log.path")
	linkName := conf.String("log.linkName")

	core, err := NewCore(logLevel, logFormat, logPath, linkName)
	if err != nil {
		return nil, err
	}

	logger := zap.New(core)
	log := logger.Sugar()

	log.Infof("logger init")
	return log, nil
}

type ZapWriter struct {
	log *zap.SugaredLogger
}

func NewZapWriter(log *zap.SugaredLogger) *ZapWriter {
	return &ZapWriter{log: log.Named("panic-recover")}
}

func (z ZapWriter) Write(p []byte) (n int, err error) {
	z.log.Error(string(p))
	return len(p), nil
}

func NewTestLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func NewCore(level, format, logPath, linkName string) (zapcore.Core, error) {
	var lvl = zap.DebugLevel
	switch level {
	case Debug:
		lvl = zap.DebugLevel
	case Info:
		lvl = zap.InfoLevel
	case Warn:
		lvl = zap.WarnLevel
	case Error:
		lvl = zap.ErrorLevel
	case Fatal:
		lvl = zap.FatalLevel
	}
	writer, err := NewWrite(logPath, linkName)
	if err != nil {
		return nil, err
	}

	return zapcore.NewCore(NewEncoder(format), writer, lvl), nil
}
func NewEncoder(format string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "msg"
	encoderConfig.CallerKey = "caller"
	encoderConfig.NameKey = "name"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	switch format {
	case TEXT:
		return zapcore.NewConsoleEncoder(encoderConfig)
	case JSON:
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func NewWrite(logPath string, linkName string) (zapcore.WriteSyncer, error) {
	//return zapcore.AddSync(os.Stdout)
	win := runtime.GOOS == "windows"

	logFile := filepath.Join(logPath, "current.log")
	options := &lumberjack.Options{
		MaxAge:     time.Hour * 24 * 30,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   false,
	}
	//20MB
	var size int64 = 20 * 1024 * 1024

	roller, err := lumberjack.NewRoller(logFile, size, options)
	if err != nil {
		return nil, err
	}

	if !win {
		_ = os.Symlink(logFile, linkName)
	}

	coreWriter := io.MultiWriter(roller, os.Stdout)
	return zapcore.AddSync(coreWriter), nil
}
