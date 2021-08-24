package log

func (_ Logger) Trace(args ...interface{}) {
	Trace(args...)
}
func (_ Logger) Tracef(format string, args ...interface{}) {
	Tracef(format, args...)
}

func (_ Logger) Debug(args ...interface{}) {
	Debug(args...)
}

func (_ Logger) Debugf(format string, args ...interface{}) {
	Debugf(format, args...)
}

func (_ Logger) Info(args ...interface{}) {
	Info(args...)
}

func (_ Logger) Infof(format string, args ...interface{}) {
	Infof(format, args...)
}

func (_ Logger) Notice(args ...interface{}) {
	Notice(args...)
}

func (_ Logger) Noticef(format string, args ...interface{}) {
	Noticef(format, args...)
}

func (_ Logger) Warn(args ...interface{}) {
	Warn(args...)
}

func (_ Logger) Warnf(format string, args ...interface{}) {
	Warnf(format, args...)
}

func (_ Logger) Error(args ...interface{}) {
	Error(args...)
}

func (_ Logger) Errorf(format string, args ...interface{}) {
	Errorf(format, args...)
}

func (_ Logger) Panic(args ...interface{}) {
	Panic(args...)
}

func (_ Logger) Panicf(format string, args ...interface{}) {
	Panicf(format, args...)
}

func (_ Logger) Fatal(args ...interface{}) {
	Fatal(args...)
}

func (_ Logger) Fatalf(format string, args ...interface{}) {
	Fatalf(format, args...)
}
