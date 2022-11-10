package logger

import "log"

// const .
const (
	INFO  = "[INFO]"
	DEBUG = "[DEBUG]"
	ERROR = "[ERROR]"
	WARN  = "[WARN]"
)

type defaultLogger struct{}

var lg Logger

func init() {
	lg = &defaultLogger{}
}

func (dl defaultLogger) info(level string, v ...interface{}) {
	// log.Printf(level, v...)
	log.Println(append([]interface{}{level}, v...)...)
}

func (dl defaultLogger) infof(level, format string, v ...interface{}) {
	log.Printf(level+" "+format, v...)
}

func (dl defaultLogger) Info(v ...interface{})                     { dl.info(INFO, v...) }
func (dl defaultLogger) Infof(format string, args ...interface{})  { dl.infof(INFO, format, args...) }
func (dl defaultLogger) Debug(v ...interface{})                    { dl.info(DEBUG, v...) }
func (dl defaultLogger) Debugf(format string, args ...interface{}) { dl.infof(DEBUG, format, args...) }
func (dl defaultLogger) Error(v ...interface{})                    { dl.info(ERROR, v...) }
func (dl defaultLogger) Errorf(format string, args ...interface{}) { dl.infof(ERROR, format, args...) }
func (dl defaultLogger) Warn(v ...interface{})                     { dl.info(WARN, v...) }
func (dl defaultLogger) Warnf(format string, args ...interface{})  { dl.infof(WARN, format, args...) }
