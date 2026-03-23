package service

import "testing"

func TestAnalyzerAnalyzeSuccess(t *testing.T) {
	parser := NewParser()
	analyzer := NewAnalyzer(parser)

	logText := "2026-03-20 10:15:22 ERROR Connection timeout\n" +
		"2026-03-20 10:15:25 INFO Retry request\n" +
		"2026-03-20 10:15:30 WARN Disk space low"

	result := analyzer.Analyze(logText)

	if result.TotalLines != 3 {
		t.Fatalf("expected TotalLines = 3, got %d", result.TotalLines)
	}

	if result.ParsedLines != 3 {
		t.Fatalf("expected ParsedLines = 3, got %d", result.ParsedLines)
	}

	if result.InvalidLines != 0 {
		t.Fatalf("expected InvalidLines = 0, got %d", result.InvalidLines)
	}

	if result.ErrorCount != 1 {
		t.Fatalf("expected ErrorCount = 1, got %d", result.ErrorCount)
	}

	if result.WarnCount != 1 {
		t.Fatalf("expected WarnCount = 1, got %d", result.WarnCount)
	}

	if result.InfoCount != 1 {
		t.Fatalf("expected InfoCount = 1, got %d", result.InfoCount)
	}

	if len(result.TopMessages) != 1 {
		t.Fatalf("expected 1 top message, got %d", len(result.TopMessages))
	}

	if result.TopMessages[0].Message != "Connection timeout" {
		t.Fatalf("expected top message 'Connection timeout', got %q", result.TopMessages[0].Message)
	}

	if result.TopMessages[0].Count != 1 {
		t.Fatalf("expected top message count = 1, got %d", result.TopMessages[0].Count)
	}
}

func TestAnalyzerAnalyzeWithInvalidLine(t *testing.T) {
	parser := NewParser()
	analyzer := NewAnalyzer(parser)

	logText := "2026-03-20 10:15:22 ERROR Connection timeout\n" +
		"INVALID LOG LINE\n" +
		"2026-03-20 10:15:25 WARN Disk space low"

	result := analyzer.Analyze(logText)

	if result.TotalLines != 3 {
		t.Fatalf("expected TotalLines = 3, got %d", result.TotalLines)
	}

	if result.ParsedLines != 2 {
		t.Fatalf("expected ParsedLines = 2, got %d", result.ParsedLines)
	}

	if result.InvalidLines != 1 {
		t.Fatalf("expected InvalidLines = 1, got %d", result.InvalidLines)
	}

	if result.ErrorCount != 1 {
		t.Fatalf("expected ErrorCount = 1, got %d", result.ErrorCount)
	}

	if result.WarnCount != 1 {
		t.Fatalf("expected WarnCount = 1, got %d", result.WarnCount)
	}

	if result.InfoCount != 0 {
		t.Fatalf("expected InfoCount = 0, got %d", result.InfoCount)
	}
}

func TestAnalyzerAnalyzeTopMessagesSorted(t *testing.T) {
	parser := NewParser()
	analyzer := NewAnalyzer(parser)

	logText := "2026-03-20 10:15:22 ERROR Connection timeout\n" +
		"2026-03-20 10:15:23 ERROR Connection timeout\n" +
		"2026-03-20 10:15:24 ERROR Database unavailable"

	result := analyzer.Analyze(logText)

	if len(result.TopMessages) != 2 {
		t.Fatalf("expected 2 top messages, got %d", len(result.TopMessages))
	}

	if result.TopMessages[0].Message != "Connection timeout" {
		t.Fatalf("expected first top message 'Connection timeout', got %q", result.TopMessages[0].Message)
	}

	if result.TopMessages[0].Count != 2 {
		t.Fatalf("expected first top message count = 2, got %d", result.TopMessages[0].Count)
	}

	if result.TopMessages[1].Message != "Database unavailable" {
		t.Fatalf("expected second top message 'Database unavailable', got %q", result.TopMessages[1].Message)
	}

	if result.TopMessages[1].Count != 1 {
		t.Fatalf("expected second top message count = 1, got %d", result.TopMessages[1].Count)
	}
}
