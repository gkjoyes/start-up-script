package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

type bLogger struct {
	*log.Logger
}

// Log define new log struct.
type Log struct {
	message string
	level   string
}

// Color macro for bash.
var (
	errCode   = "\033[38;5;196m"
	infoCode  = "\033[38;5;266m"
	warnCode  = "\033[38;5;214m"
	flagsCode = "\033[38;5;8m"
	nameCode  = "\033[38;5;157m"
	msgCode   = "\033[38;5;243m"
	resetCode = "\033[0m"
	bl        = bLogger{log.New(os.Stdout, nameCode+"start-up-script"+flagsCode, 0)}
)

// Message sets logging message.
func (l Log) Message(message ...string) Log {
	l.message = strings.Join(message, " ")
	return l
}

// Error level.
func Error() Log {
	return Log{level: "error"}
}

// Info level.
func Info() Log {
	return Log{level: "info"}
}

// Warn level.
func Warn() Log {
	return Log{level: "warn"}
}

// Log construct log message and display.
func (l Log) Log() {
	out := "[" + time.Now().Format("15:04:05") + "]"
	switch l.level {
	case "error":
		out += errCode
	case "info":
		out += infoCode
	case "warn":
		out += warnCode
	}

	out += l.message + resetCode
	bl.Print(out)
}

// FormattedMessage format logging message.
func FormattedMessage(message ...string) string {
	return msgCode + "[" + strings.Join(message, " ") + "]"
}
