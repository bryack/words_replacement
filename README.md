# Words

A Go CLI application for intelligent word form replacement implementing hexagonal architecture with clean separation of concerns and production-ready linguistic data integration.

## Features

- **Intelligent Word Replacement**: Context-aware nominative form transformations for Russian nouns
- **Clean Architecture**: Hexagonal architecture with dependency injection and interface-driven design
- **SQLite Database**: Persistent file-based database for fast word form lookups
- **Kaikki.org Integration**: Production linguistic data from Kaikki.org JSONL dictionary
- **Multiple Data Sources**: Support for SQLite, MediaWiki API, and custom providers
- **Unicode-Aware Processing**: Handles Cyrillic text boundaries and stress mark removal
- **Comprehensive Testing**: Unit, integration, and acceptance test coverage with ATDD methodology
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
  - `adapters/sqlite/` - SQLite database provider with JSONL data loading
  - `adapters/wiktionary/` - MediaWiki API client
  - `adapters/cli/` - CLI driver for file operations
- `contracts/` - Interface definitions and ports
- `specifications/` - Behavior-driven specifications
- `wiki/` - MediaWiki API client integration

## Data Sources

The application uses Kaikki.org JSONL dictionary data:
- Parses Russian noun entries from JSONL format
- Extracts multiple grammatical cases (nominative, accusative, genitive, dative, instrumental, prepositional)
- Removes stress marks for consistent matching
- Loads data into persistent SQLite database with smart loading detection

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
- `github.com/mattn/go-sqlite3` (database)

## License

See project license file.
