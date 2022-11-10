package logger

// Logger .
type Logger interface {
	Info(v ...interface{})
	Infof(format string, args ...interface{})
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
}

// Info .
func Info(v ...interface{}) { lg.Info(v...) }

// Infof .
func Infof(format string, args ...interface{}) { lg.Infof(format, args...) }

// Debug .
func Debug(v ...interface{}) { lg.Debug(v...) }

// Debugf .
func Debugf(format string, args ...interface{}) { lg.Debugf(format, args...) }

// Error .
func Error(v ...interface{}) { lg.Error(v...) }

// Errorf .
func Errorf(format string, args ...interface{}) { lg.Errorf(format, args...) }

// Warn .
func Warn(v ...interface{}) { lg.Warn(v...) }

// Warnf .
func Warnf(format string, args ...interface{}) { lg.Warnf(format, args...) }
