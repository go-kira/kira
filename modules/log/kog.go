// Package log implements a simple logging package.
package log

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"sync"
	"time"
)

var (
	// DefaultFormatter - default log formatter.
	DefaultFormatter = new(DefaultLogFormatter)
	// StdLog - default log
	StdLog = New(os.Stderr, DefaultFormatter, Fields{})
)

// Logger - kira logger.
type Logger struct {
	formatter Formatter
	level     Level
	fields    Fields

	Writer io.Writer
	lock   sync.Mutex
}

const defaultLevel = DebugLevel

// New creates a new Logger.
func New(w io.Writer, f Formatter, fields Fields) *Logger {
	return &Logger{
		Writer:    w,
		formatter: f,
		level:     defaultLevel,
		fields:    fields,
	}
}

// log writes the output for a logging event.
func (l *Logger) log(level Level, msg interface{}) {
	if l.level <= level {
		if err := l.formatter.Format(l, level, msg, time.Now()); err != nil {
			stdlog.Printf("error logging: %s", err)
		}
	}
}

// SetFormatter sets the Formatter for the logger.
func (l *Logger) SetFormatter(f Formatter) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.formatter = f
}

// SetLevel sets the level severity for the logger.
func (l *Logger) SetLevel(level Level) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.level = level
}

// SetWriter sets the level severity for the logger.
func (l *Logger) SetWriter(w io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.Writer = w
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	fields := Fields{}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields[key] = value

	return New(l.Writer, l.formatter, fields)
}

// Debug calls log to print to the logger.
func (l *Logger) Debug(v ...interface{}) {
	l.log(DebugLevel, fmt.Sprint(v...))
}

// Debugf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(f string, v ...interface{}) {
	l.Debug(fmt.Sprintf(f, v...))
}

// Info calls log to print to the logger.
func (l *Logger) Info(v ...interface{}) {
	l.log(InfoLevel, fmt.Sprint(v...))
}

// Infof calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(f string, v ...interface{}) {
	l.Info(fmt.Sprintf(f, v...))
}

// Warn calls log to print to the logger.
func (l *Logger) Warn(v ...interface{}) {
	l.log(WarnLevel, fmt.Sprint(v...))
}

// Warnf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warnf(f string, v ...interface{}) {
	l.Warn(fmt.Sprintf(f, v...))
}

// Error calls log to print to the logger.
func (l *Logger) Error(v ...interface{}) {
	l.log(ErrorLevel, fmt.Sprint(v...))
}

// Errorf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(f string, v ...interface{}) {
	l.Error(fmt.Sprintf(f, v...))
}

// Fatal calls log to print to the logger.
func (l *Logger) Fatal(v ...interface{}) {
	l.log(FatalLevel, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(f string, v ...interface{}) {
	l.Fatal(fmt.Sprintf(f, v...))
}

// Panic calls log to print to the logger.
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log(PanicLevel, s)

	panic(s)
}

// Panicf calls l.log to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Panicf(f string, v ...interface{}) {
	l.Panic(fmt.Sprintf(f, v...))
}

// Debug calls log to print to the standard logger.
func Debug(v ...interface{}) {
	StdLog.Debug(v...)
}

// Debugf calls log to print to the standard logger.
func Debugf(f string, v ...interface{}) {
	StdLog.Debug(fmt.Sprintf(f, v...))
}

// Info calls log to print to the standard logger.
func Info(v ...interface{}) {
	StdLog.Info(v...)
}

// Infof calls log to print to the standard logger.
func Infof(f string, v ...interface{}) {
	StdLog.Info(fmt.Sprintf(f, v...))
}

// Warn calls log to print to the standard logger.
func Warn(v ...interface{}) {
	StdLog.Warn(v...)
}

// Warnf calls log to print to the standard logger.
func Warnf(f string, v ...interface{}) {
	StdLog.Warn(fmt.Sprintf(f, v...))
}

// Error calls log to print to the standard logger.
func Error(v ...interface{}) {
	StdLog.Error(v...)
}

// Errorf calls log to print to the standard logger.
func Errorf(f string, v ...interface{}) {
	StdLog.Error(fmt.Sprintf(f, v...))
}

// Fatal calls log to print to the standard logger.
func Fatal(v ...interface{}) {
	StdLog.Fatal(v...)
}

// Fatalf calls log to print to the standard logger.
func Fatalf(f string, v ...interface{}) {
	StdLog.Fatal(fmt.Sprintf(f, v...))
}

// Panic calls log to print to the standard logger.
func Panic(v ...interface{}) {
	StdLog.Panic(v...)
}

// Panicf calls log to print to the standard logger.
func Panicf(f string, v ...interface{}) {
	StdLog.Panic(fmt.Sprintf(f, v...))
}
