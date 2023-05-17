package log

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

type (
	Logger interface {
		SetLevel(level string)
		Debug(msg string)
		Debugf(format string, v ...interface{})
		Info(msg string)
		Infof(format string, v ...interface{})
		Error(msg string)
		Errorf(format string, v ...interface{})
		Fatal(err error)
	}

	SimpleLogger struct {
		level string
		flag  int
		debug *log.Logger
		info  *log.Logger
		error *log.Logger
	}

	logLevel struct {
		Debug string
		Info  string
		Error string
	}
)

var (
	LogLevel = logLevel{
		Debug: "debug",
		Info:  "info",
		Error: "error",
	}
)

// NewLogger return a new logger.
func NewLogger(level string, isFlag bool) *SimpleLogger {
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lmicroseconds //| adr.Lshortfile
	}

	return newLogger(level, flag)
}

// SetLevel sets the logging level preference
func newLogger(level string, flag int) *SimpleLogger {
	switch level {
	case LogLevel.Debug:
		return &SimpleLogger{
			level: LogLevel.Debug,
			flag:  flag,
			debug: log.New(os.Stderr, "[DBG] ", flag),
			info:  log.New(os.Stderr, "[INF] ", flag),
			error: log.New(os.Stderr, "[ERR] ", flag),
		}

	case LogLevel.Info:
		return &SimpleLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(os.Stderr, "INFO: ", flag),
			error: log.New(os.Stderr, "ERROR: ", flag),
		}

	case LogLevel.Error:
		return &SimpleLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(ioutil.Discard, "INFO: ", flag),
			error: log.New(os.Stderr, "ERROR: ", flag),
		}

	default:
		return &SimpleLogger{
			level: LogLevel.Info,
			flag:  flag,
			debug: log.New(ioutil.Discard, "DEBUG: ", flag),
			info:  log.New(ioutil.Discard, "INFO: ", flag),
			error: log.New(ioutil.Discard, "ERROR: ", flag),
		}
	}
}

func (sl *SimpleLogger) SetLevel(level string) {
	if sl.level != level {
		*sl = *newLogger(level, sl.flag)
	}
}

// Debug calls l.Output to print to the logger.
func (sl *SimpleLogger) Debug(msg string) {
	sl.debug.Println(msg)
}

// Debugf calls l.Output to print to the logger.
func (sl *SimpleLogger) Debugf(format string, v ...interface{}) {
	sl.debug.Printf(format, v...)
}

// Info calls l.Output to print to the logger.
func (sl *SimpleLogger) Info(msg string) {
	sl.info.Println(msg)
}

// Infof calls l.Output to print to the logger.
func (sl *SimpleLogger) Infof(format string, v ...interface{}) {
	sl.info.Printf(format, v...)
}

// Error calls l.Output to print to the logger.
func (sl *SimpleLogger) Error(msg string) {
	sl.error.Println(msg)
}

// Errorf calls l.Output to print to the logger.
func (sl *SimpleLogger) Errorf(format string, v ...interface{}) {
	sl.error.Printf(format, v...)
}

// Dump calls l.Output to print error to the logger.
func (sl *SimpleLogger) Dump(error error) {
	sl.error.Println(error.Error())
}

// Fatal calls l.Output to print error to the logger and call os.Exit(1).
func (sl *SimpleLogger) Fatal(error error) {
	sl.error.Fatal(error.Error())
}

// SetDebugOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetDebugOutput(debug *bytes.Buffer) {
	sl.debug = log.New(debug, "", 0)
}

// SetInfoOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetInfoOutput(info *bytes.Buffer) {
	sl.info = log.New(info, "", 0)
}

// SetErrorOutput set the internal logger.
// Used for package testing.
func (sl *SimpleLogger) SetErrorOutput(error *bytes.Buffer) {
	sl.error = log.New(error, "", 0)
}
