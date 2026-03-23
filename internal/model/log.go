package model

import "time"

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

type AnalyzeRequest struct {
	LogText string `json:"log_text"`
}

type MessageCount struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type AnalyzeResponse struct {
	TotalLines   int            `json:"total_lines"`
	ParsedLines  int            `json:"parsed_lines"`
	ErrorCount   int            `json:"error_count"`
	WarnCount    int            `json:"warn_count"`
	InfoCount    int            `json:"info_count"`
	InvalidLines int            `json:"invalid_lines"`
	TopMessages  []MessageCount `json:"top_messages"`
}
