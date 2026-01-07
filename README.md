# Words

A Go utility for text processing and wiki content retrieval.

## Features

- Text replacement with regex pattern matching
- Wikipedia/MediaWiki API client for content retrieval
- Command-line interface for text processing

## Installation

```bash
go build -o words
```

## Usage

```bash
./words [input-file]
```

## Project Structure

- `words.go` - Main text processing logic
- `wiki/` - MediaWiki API client package
- `*_test.go` - Test files

## Development

Run tests:
```bash
go test ./...
```

## License

See project license file.
