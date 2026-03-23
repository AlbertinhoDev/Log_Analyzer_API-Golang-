# Log Analyzer API

A small REST API service written in Go for parsing and analyzing raw log text.

The project is designed to demonstrate backend engineering basics:
- clean project structure
- separation of concerns
- HTTP API development with the Go standard library
- simple log parsing and analytics

## Features

- `GET /health` health check endpoint
- `POST /analyze` log analysis endpoint
- parsing log lines in format: `YYYY-MM-DD HH:MM:SS LEVEL message`
- counting:
  - total lines
  - parsed lines
  - invalid lines
  - error count
  - warn count
  - info count
- extracting top error messages sorted by frequency

## Tech Stack

- Go
- Standard library:
  - `net/http`
  - `encoding/json`
  - `strings`
  - `sort`
  - `time`

## Project Structure

```text
cmd/server/main.go
internal/handler/logs.go
internal/handler/logs_test.go
internal/service/parser.go
internal/service/parser_test.go
internal/service/analyzer.go
internal/service/analyzer_test.go
internal/model/log.go
```

## API Endpoints

### `GET /health`

Health check endpoint.

Response:

```json
{
  "status": "ok"
}
```

### `POST /analyze`

Analyzes raw log text.

Request body:

```json
{
  "log_text": "2026-03-20 10:15:22 ERROR Connection timeout\n2026-03-20 10:15:25 INFO Retry request"
}
```

Response:

```json
{
  "total_lines": 2,
  "parsed_lines": 2,
  "error_count": 1,
  "warn_count": 0,
  "info_count": 1,
  "invalid_lines": 0,
  "top_messages": [
    {
      "message": "Connection timeout",
      "count": 1
    }
  ]
}
```

## Supported Log Format

Each log line must follow this format:

```text
YYYY-MM-DD HH:MM:SS LEVEL message
```

Example:

```text
2026-03-20 10:15:22 ERROR Connection timeout
```

Supported levels:
- `ERROR`
- `WARN`
- `INFO`

If a line does not match the format, it is counted as `invalid_lines`.

## How to Run

1. Make sure Go is installed.
2. Run the server:

```bash
go run ./cmd/server
```

By default, the server starts on:

```text
http://localhost:8080
```

## Example Requests

Health check:

```bash
curl http://localhost:8080/health
```

Analyze logs:

```bash
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "log_text": "2026-03-20 10:15:22 ERROR Connection timeout\n2026-03-20 10:15:25 INFO Retry request"
  }'
```

Analyze logs with invalid line:

```bash
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "log_text": "2026-03-20 10:15:22 ERROR Connection timeout\nINVALID LOG LINE\n2026-03-20 10:15:25 WARN Disk space low"
  }'
```

## Error Handling

If request JSON is invalid, the API returns:

```json
{
  "error": "invalid JSON body"
}
```

## Testing

Run all tests with:

```bash
go test ./...
```

The project includes:
- unit tests for the parser
- unit tests for the analyzer
- HTTP handler tests using `httptest`

## Docker

Build the Docker image:

```bash
docker build -t log-analyzer-api .
```

Run the container:

```bash
docker run -p 8080:8080 log-analyzer-api
```

After the container starts, the API is available at:

```text
http://localhost:8080
```

## Notes

This project intentionally keeps the implementation simple while following a modular backend structure:
- handler layer for HTTP
- service layer for parsing and analysis
- model layer for data structures

It uses the Go standard library as much as possible and avoids heavy frameworks.
