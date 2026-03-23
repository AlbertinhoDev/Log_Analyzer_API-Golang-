package service

import (
	"errors"
	"loganalyzerapi/internal/model"
	"strings"
	"time"
)

const logTimeLayout = "2006-01-02 15:04:05"

var (
	ErrInvalidLogFormat = errors.New("invalid log format")
	ErrInvalidLogLevel  = errors.New("invalid log level")
	ErrEmptyMessage     = errors.New("log text cannot be empty")
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseLine(line string) (model.LogEntry, error) {
	line = strings.TrimSpace(line)

	if line == "" {
		return model.LogEntry{}, ErrInvalidLogFormat
	}

	parts := strings.SplitN(line, " ", 4)

	if len(parts) < 4 {
		return model.LogEntry{}, ErrInvalidLogFormat
	}

	timestampText := parts[0] + " " + parts[1]

	timestamp, err := time.Parse(logTimeLayout, timestampText)
	if err != nil {
		return model.LogEntry{}, ErrInvalidLogFormat
	}

	level := parts[2]

	if !isValidLevel(level) {
		return model.LogEntry{}, ErrInvalidLogLevel
	}

	message := strings.TrimSpace(parts[3])

	if message == "" {
		return model.LogEntry{}, ErrEmptyMessage
	}

	return model.LogEntry{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
	}, nil
}

func isValidLevel(level string) bool {
	switch level {
	case "ERROR", "WARN", "INFO":
		return true
	default:
		return false
	}
}
