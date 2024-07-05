package log

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	AnsiReset  = "\033[0m"
	AnsiDim    = "\033[0;2m"
	AnsiYellow = "\033[1;33m"
	AnsiRed    = "\033[1;31m"
)

var (
	DefaultColor = AnsiReset
	InfoColor    = AnsiReset
	DebugColor   = AnsiDim
	WarningColor = AnsiYellow
	ErrorColor   = AnsiRed
)

// Message is a single log message.
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	Level     Level     `json:"level"`
	Labels    Labels    `json:"labels"`
	Message   string    `json:"message"`
}

func (m Message) String() string {
	var Color string
	if term.IsTerminal(int(os.Stderr.Fd())) {
		switch m.Level {
		case "debug":
			Color = DebugColor
		case "warning":
			Color = WarningColor
		case "error":
			Color = ErrorColor
		case "info":
			Color = InfoColor
		default:
			Color = DefaultColor
		}
	} else {
		Color = ""
		DefaultColor = ""
	}
	return fmt.Sprintf(
		"%s%s\t%s\t%s\t%s%s",
		Color,
		m.Timestamp.Format(time.RFC3339),
		m.Level,
		m.Labels,
		m.Message,
		DefaultColor,
	)
}
