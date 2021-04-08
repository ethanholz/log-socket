package lambo_log_socket

import (
	"testing"
)

// SetLogLevel set log level of logger
func TestSetLogLevel(t *testing.T) {
	logLevels := [...]Level{LTrace, LDebug, LInfo, LWarn, LError, LPanic, LFatal}
	var c Client
	for _, x := range logLevels {
		c.SetLogLevel(x)
		if c.GetLogLevel() != x {
			t.Errorf("Got logLevel %d, but expected %d", int(c.GetLogLevel()), int(x))
		}
	}
}

// Trace prints out logs on trace level
func TestTrace(t *testing.T) {
	var c Client
	c.SetLogLevel(LTrace)
	//	if logLevel >= LTrace {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Debug(args...)
	//	}
}

// Debug prints out logs on debug level
func TestDebug(t *testing.T) {
	//	if logLevel >= LDebug {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Debug(args...)
	//	}
}

// Info prints out logs on info level
func TestInfo(t *testing.T) {
	//	if logLevel >= LInfo {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Info(args...)
	//	}
}

// Warn prints out logs on warn level
func TestWarn(t *testing.T) {
	//	if logLevel >= LWarn {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Warn(args...)
	//	}
}

// Error prints out logs on error level
func TestError(t *testing.T) {
	//	if logLevel >= LError {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Error(args...)
	//	}
}

// Fatal prints out logs on fatal level
func TestFatal(t *testing.T) {
	//	if logLevel >= LFatal {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Fatal(args...)
	//	}
}

// Panic prints out logs on panic level
func TestPanic(t *testing.T) {
	//	if logLevel >= LPanic {
	//		entry := logger.WithFields(logrus.Fields{})
	//		entry.Data["file"] = fileInfo(2)
	//		entry.Panic(args...)
	//	}
}
