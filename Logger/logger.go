package Logger

import "log"

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewFZLogger(info *log.Logger, warning *log.Logger, error *log.Logger) *Logger {
	return &Logger{infoLogger: info, warningLogger: warning, errorLogger: error}
}

func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warningLogger.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}
