package logger

import (
	"BeGo/constants"
	"errors"
	"io"

	"github.com/rifflock/lfshook"
	logrus "github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(path string, env string) (*logrus.Logger, error) {
	switch env {
	case constants.DEV_ENV_KEY: //DEV
		log = logrus.New()
		log.Formatter = &logrus.TextFormatter{TimestampFormat: constants.TIME_FORMAT_DEFAULT}
	case constants.DEBUG_ENV_KEY: //DEBUG
		fallthrough
	case constants.PROD_ENV_KEY: // PROD
		log = logrus.New()
		log.Out = io.Discard
		log.Hooks.Add((lfshook.NewHook(
			lfshook.PathMap{
				logrus.InfoLevel:  path + "/" + env + "_info.log",
				logrus.TraceLevel: path + "/" + env + "_trace.log",
				logrus.WarnLevel:  path + "/" + env + "_warn.log",
				logrus.DebugLevel: path + "/" + env + "_debug.log",
				logrus.ErrorLevel: path + "/" + env + "_error.log",
				logrus.FatalLevel: path + "/" + env + "_fatal.log",
				logrus.PanicLevel: path + "/" + env + "_panic.log",
			},
			&logrus.JSONFormatter{
				TimestampFormat: constants.TIME_FORMAT_DEFAULT,
			},
		)))
	default:
		errTitle := "cannot create logger"
		Error(errTitle)
		return nil, errors.New(errTitle)
	}

	return log, nil
}

func Debug(format string, v ...any) {
	log.Debugf(constants.LogDebugPrefix+format, v...)
}

func Info(format string, v ...any) {
	log.Infof(constants.LogInfoPrefix+format, v...)
}

func Warn(format string, v ...any) {
	log.Warnf(constants.LogWarnPrefix+format, v...)
}

func Error(format string, v ...any) {
	log.Errorf(constants.LogErrorPrefix+format, v...)
}

func Fatal(format string, v ...any) {
	log.Fatalf(constants.LogFatalPrefix+format, v...)
}

func Panic(format string, v ...any) {
	log.Panicf(constants.LogPanicPrefix+format, v...)
}
