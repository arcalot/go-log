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
	SetColor = map[Level]string{
		"debug":   AnsiDim,
		"info":    AnsiReset,
		"warning": AnsiYellow,
		"error":   AnsiRed,
	}
	ResetColor = AnsiReset
)

func init() {
	if !term.IsTerminal(int(os.Stderr.Fd())) {
		SetColor = map[Level]string{}
		ResetColor = ""
	}
}

// Message is a single log message.
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	Level     Level     `json:"level"`
	Labels    Labels    `json:"labels"`
	Message   string    `json:"message"`
}

func (m Message) String() string {
	return fmt.Sprintf(
		"%s%s\t%s\t%s\t%s%s",
		SetColor[m.Level],
		m.Timestamp.Format(time.RFC3339),
		m.Level,
		m.Labels,
		m.Message,
		ResetColor,
	)
}
