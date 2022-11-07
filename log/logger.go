package log

import (
	"os"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerSetup sync.Once

//SetupLogger sets up a centralized logger for the repo with json format and default info level logging

func SetupLogger() {
	loggerSetup.Do(func() {
		logger = logrus.New()
	})

	logger.SetFormatter(&logrus.JSONFormatter{})
	var err error
	logger.Level, err = logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.Level = logrus.InfoLevel
	}
}

func checkLoggerIntialized() {
	if logger == nil {
		SetupLogger()
	}
}

func Info(msg interface{}, extraFields ...map[string]interface{}) {
	checkLoggerIntialized()
	logger.WithFields(GetWithFields(false, extraFields)).Info(getErrorMessage(msg))
}

func Debug(msg interface{}, extraFields ...map[string]interface{}) {
	checkLoggerIntialized()
	logger.WithFields(GetWithFields(true, extraFields)).Debug(getErrorMessage(msg))
}

func Warning(msg interface{}, extraFields ...map[string]interface{}) {
	checkLoggerIntialized()
	logger.WithFields(GetWithFields(true, extraFields)).Warn(getErrorMessage(msg))
}

func Panic(msg interface{}, extraFields ...map[string]interface{}) {
	checkLoggerIntialized()
	logger.WithFields(GetWithFields(true, extraFields)).Panic(getErrorMessage(msg))
}

func Error(msg interface{}, extraFields ...map[string]interface{}) {
	checkLoggerIntialized()
	logger.WithFields(GetWithFields(true, extraFields)).Error(getErrorMessage(msg))
}

func getErrorMessage(codemsg interface{}) (msg interface{}) {
	switch code := codemsg.(type) {
	case int:
		msg = LogString(code)
	default:
		msg = codemsg
	}
	return
}

func GetWithFields(tracing bool, extraFields []map[string]interface{}) logrus.Fields {
	var fields = logrus.Fields{}
	/* function runtime.Caller is internally in used in logrus to get the file , function , line number details */
	if tracing { // tracing flag added to get the file details in case debug, warning, panic, error
		pc, file, line, ok := runtime.Caller(2)
		if ok {
			f := runtime.FuncForPC(pc)
			fields = logrus.Fields{
				"source": file,
				"method": f.Name(),
				"line":   line,
			}
		}
	}
	for _, extraf := range extraFields {
		for x, y := range extraf {
			fields[x] = y
		}
	}

	return fields
}
