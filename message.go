package log

import (
	"fmt"
	"time"
)

const (
	ResetColor   = "\033[0m"
	DebugColor   = "\033[0;2m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
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
	if m.Level == "debug" {
		Color = DebugColor
	} else if m.Level == "warning" {
		Color = WarningColor
	} else if m.Level == "error" {
		Color = ErrorColor
	} else {
		Color = ResetColor
	}
	return fmt.Sprintf("%s%s\t%s\t%s\t%s%s", Color, m.Timestamp.Format(time.RFC3339), m.Level, m.Labels, m.Message, ResetColor)
}
