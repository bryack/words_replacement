# Words

A Go CLI application for intelligent word form replacement implementing hexagonal architecture with clean separation of concerns.

## Features

- **Intelligent Word Replacement**: Context-aware singular/plural form transformations
- **Clean Architecture**: Hexagonal architecture with dependency injection
- **MediaWiki Integration**: API client for linguistic data retrieval
- **Comprehensive Testing**: Unit, integration, and acceptance test coverage
- **CLI Interface**: Command-line tool with proper error handling

## Installation

```bash
go build -o words ./cmd/cli
```

## Usage

```bash
./words [options]
```

## Architecture

The project demonstrates hexagonal architecture principles:

- `cmd/cli/` - Application entry point and acceptance tests
- `internal/` - Core business logic (replacer, cli)
- `adapters/` - External interface implementations
- `contracts/` - Interface definitions and ports
- `specifications/` - Behavior-driven specifications
- `wiki/` - MediaWiki API client integration

## Development

### Run tests
```bash
go test ./...
```

### Run with coverage
```bash
go test -cover ./...
```

### Format code
```bash
gofmt -w .
```

## Dependencies

- Go 1.25.5+
- `github.com/alecthomas/assert/v2` (testing)

## License

See project license file.
