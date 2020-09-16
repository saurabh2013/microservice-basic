package logger

import (
	"fmt"
	"os"

	"github.com/saurabh2013/microservice-basic/internal/constant"

	lr "github.com/sirupsen/logrus"
)

//Log is logger instance
type Log struct {
	lr.Logger
}

//New ...
func New(level, output string) (*Log, error) {

	var log = new(Log)

	// set logger output
	switch output {
	case "json":
		log.SetFormatter(&lr.JSONFormatter{})
	default:
		log.SetFormatter(&lr.TextFormatter{
			FullTimestamp: true,
		})
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	switch level {
	case constant.PanicLevel:
		log.SetLevel(lr.PanicLevel)
	case constant.FatalLevel:
		log.SetLevel(lr.FatalLevel)
	case constant.ErrorLevel:
		log.SetLevel(lr.ErrorLevel)
	case constant.WarnLevel:
		log.SetLevel(lr.WarnLevel)
	case constant.InfoLevel:
		log.SetLevel(lr.InfoLevel)
	case constant.DebugLevel:
		log.SetLevel(lr.DebugLevel)
	case constant.TraceLevel:
		log.SetLevel(lr.TraceLevel)
	default:
		return nil, fmt.Errorf("log level not found '%s'", level)
	}
	return log, nil
}
