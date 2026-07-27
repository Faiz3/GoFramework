package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	level     LogLevel
	writer    io.Writer
	appLogger *log.Logger
}

var instance *Logger

func New() *Logger {
	l := &Logger{
		level:     INFO,
		writer:    os.Stdout,
		appLogger: log.New(os.Stdout, "", 0),
	}
	instance = l
	return l
}

func GetInstance() *Logger {
	if instance == nil {
		instance = New()
	}
	return instance
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
	l.appLogger.SetOutput(w)
}

func (l *Logger) log(level LogLevel, msg string, args ...interface{}) {
	if level < l.level {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(msg, args...)
	l.appLogger.Printf("[%s] %s: %s", now, level.String(), message)

	if level == FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.log(DEBUG, msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.log(INFO, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.log(WARNING, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.log(ERROR, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log(FATAL, msg, args...)
}

func Debug(msg string, args ...interface{}) {
	GetInstance().Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	GetInstance().Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	GetInstance().Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	GetInstance().Error(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	GetInstance().Fatal(msg, args...)
}
