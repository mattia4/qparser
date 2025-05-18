package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	verbose bool
}

func NewLogger(verbose bool) *Logger {
	return &Logger{verbose: verbose}
}

func (l *Logger) Info(msg string, a ...any) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s: %s\n", ts, TagInfo, fmt.Sprintf(msg, a...))
}

func (l *Logger) Error(msg string, a ...any) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stderr, "%s %s: %s\n", ts, TagError, fmt.Sprintf(msg, a...))
}

func (l *Logger) Fatal(msg string, a ...any) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stderr, "%s %s: %s\n", ts, TagFatal, fmt.Sprintf(msg, a...))
	os.Exit(1)
}

func (l *Logger) Success(msg string, a ...any) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s: %s\n", ts, TagSuccess, fmt.Sprintf(msg, a...))
}

func (l *Logger) Debug(msg string, a ...any) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	if l.verbose {
		fmt.Printf("%s %s: %s\n", ts, TagDebug, fmt.Sprintf(msg, a...))
	}
}
