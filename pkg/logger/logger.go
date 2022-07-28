package logger

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/knadh/koanf"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
	var opts []rotatelogs.Option
	opts = append(opts, rotatelogs.WithMaxAge(time.Hour*24*7))
	opts = append(opts, rotatelogs.WithRotationTime(time.Hour*24))
	opts = append(opts, rotatelogs.WithRotationSize(1024*1024*10))

	if !win {
		opts = append(opts, rotatelogs.WithLinkName(linkName))
	}
	logFile := filepath.Join(logPath, "log-%Y%m%d.log")
	writer, err := rotatelogs.New(logFile, opts...)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(writer), nil
}
