package service

import "testing"

func TestParserParseLineSuccess(t *testing.T) {
	parser := NewParser()

	line := "2026-03-20 10:15:22 ERROR Connection timeout"

	entry, err := parser.ParseLine(line)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if entry.Level != "ERROR" {
		t.Fatalf("expected level ERROR, got %s", entry.Level)
	}

	if entry.Message != "Connection timeout" {
		t.Fatalf("expected message 'Connection timeout', got %q", entry.Message)
	}

	expectedTimestamp := "2026-03-20 10:15:22"
	if entry.Timestamp.Format(logTimeLayout) != expectedTimestamp {
		t.Fatalf("expected timestamp %s, got %s", expectedTimestamp, entry.Timestamp.Format(logTimeLayout))
	}
}

func TestParserParseLineInvalidLevel(t *testing.T) {
	parser := NewParser()

	line := "2026-03-20 10:15:22 DEBUG Connection timeout"

	_, err := parser.ParseLine(line)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrInvalidLogLevel {
		t.Fatalf("expected ErrInvalidLogLevel, got %v", err)
	}
}

func TestParserParseLineEmptyMessage(t *testing.T) {
	parser := NewParser()

	line := "2026-03-20 10:15:22 ERROR "

	_, err := parser.ParseLine(line)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrInvalidLogFormat {
		t.Fatalf("expected ErrInvalidLogFormat, got %v", err)
	}
}
