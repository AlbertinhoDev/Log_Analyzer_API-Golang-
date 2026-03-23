package service

import (
	"loganalyzerapi/internal/model"
	"sort"
	"strings"
)

type Analyzer struct {
	parser *Parser
}

func NewAnalyzer(parser *Parser) *Analyzer {
	return &Analyzer{
		parser: parser,
	}
}

func (a *Analyzer) Analyze(logText string) model.AnalyzeResponse {
	response := model.AnalyzeResponse{
		TopMessages: []model.MessageCount{},
	}

	lines := splitLines(logText)

	response.TotalLines = len(lines)

	errorMessages := make(map[string]int)

	for _, line := range lines {
		entry, err := a.parser.ParseLine(line)
		if err != nil {
			response.InvalidLines++
			continue
		}

		response.ParsedLines++

		switch entry.Level {
		case "ERROR":
			response.ErrorCount++
			errorMessages[entry.Message]++
		case "WARN":
			response.WarnCount++
		case "INFO":
			response.InfoCount++
		}
	}

	response.TopMessages = buildTopMessages(errorMessages)

	return response
}

func splitLines(logText string) []string {
	normalized := strings.ReplaceAll(logText, "\r\n", "\n")
	normalized = strings.TrimSpace(normalized)

	if normalized == "" {
		return []string{}
	}

	return strings.Split(normalized, "\n")
}

func buildTopMessages(errorMessages map[string]int) []model.MessageCount {
	topMessages := make([]model.MessageCount, 0, len(errorMessages))

	for message, count := range errorMessages {
		topMessages = append(topMessages, model.MessageCount{
			Message: message,
			Count:   count,
		})
	}

	sort.Slice(topMessages, func(i, j int) bool {
		if topMessages[i].Count == topMessages[j].Count {
			return topMessages[i].Message < topMessages[j].Message
		}
		return topMessages[i].Count > topMessages[j].Count
	})

	return topMessages
}
