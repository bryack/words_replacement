# AGENTS.md

This file provides context and instructions for AI coding agents working on the Words project.

## Project Overview

Words is a Go CLI application implementing hexagonal architecture for intelligent word form replacement. It demonstrates clean architecture principles with dependency injection, interface-driven design, comprehensive testing strategies, and production-ready data integration with SQLite and Kaikki.org linguistic datasets.

**Current Implementation Status:**
- ✅ SQLite database integration with persistent file-based storage
- ✅ Kaikki.org JSONL data loading and parsing with smart loading detection
- ✅ Multi-case form extraction (nominative, accusative, genitive, dative, instrumental, prepositional)
- ✅ Unicode-aware word replacement with Cyrillic boundary detection
- ✅ Stress mark removal for consistent Russian text matching
- ✅ DataLoader pattern for flexible data sources with persistent database
- ✅ ATDD methodology with acceptance tests as north-star
- ✅ Comprehensive benchmarking for performance optimization
- ✅ Fast unit tests with test data, integration tests with real JSONL
- ⚠️ MediaWiki API integration (external service may return 403)

## Build and Test Commands

### Build
```bash
go build -o words ./cmd/cli
```

### Test
```bash
go test ./...
```

### Run with coverage
```bash
go test -cover ./...
```

### Run acceptance tests
```bash
go test ./cmd/cli/...
```

### Format code
```bash
gofmt -w .
```

## Code Style Guidelines

- Follow standard Go conventions and `gofmt` formatting
- Use PascalCase for exported functions (`NewReplacer`, `Replace`)
- Use camelCase for local variables, PascalCase for exported variables
- Use UPPER_SNAKE_CASE for package-level constants
- Test files must end with `_test.go` and be placed alongside source files
- Interface names should end with -er (`FormsProvider`, `WordReplacer`)

## Architecture Guidelines

- **Hexagonal Architecture**: Maintain strict separation between core business logic (`internal/`) and external adapters (`adapters/`)
- **Dependency Injection**: All external dependencies must be injected through interfaces defined in `contracts/`
- **Interface-First Design**: Define contracts before implementations
- **Provider Pattern**: Use provider interfaces for external services (wiki, database, file system)

## Testing Instructions

- **Unit Tests**: Place in `internal/` packages alongside source code
- **Integration Tests**: Place in `adapters/` packages for external interface testing
- **Acceptance Tests**: Place in `cmd/` packages for end-to-end scenarios
- Use `github.com/alecthomas/assert/v2` for enhanced assertions
- Test error conditions and edge cases
- Aim for high test coverage of public functions

## Dependencies

- Go version 1.25.5 required
- Module: `github.com/bryack/words`
- Dependencies: `github.com/alecthomas/assert/v2` for testing, `github.com/mattn/go-sqlite3` for database operations
- Target platform: Linux (primary), cross-platform compatible

## Project Structure

- `cmd/cli/` - CLI application entry point and acceptance tests
- `internal/` - Core business logic (private packages)
- `adapters/` - External interface implementations
  - `adapters/sqlite/` - SQLite database provider with JSONL data loading
    - `models.go` - JSONL data structures (KaikkiEntry, WordForm)
    - `models_test.go` - Models unit tests
    - `extractor.go` - Multi-case form extraction with stress removal
    - `extractor_test.go` - Extractor unit tests
    - `loader.go` - JSONL file parsing
    - `loader_test.go` - Loader unit tests
    - `provider.go` - SQLite forms provider with DataLoader pattern
  - `adapters/wiktionary/` - MediaWiki API client
  - `adapters/cli/` - CLI driver for file operations
- `contracts/` - Interface definitions and ports
- `specifications/` - Behavior-driven specifications
- `wiki/` - MediaWiki API client package
- `.kiro/steering/` - AI agent steering documents

## Security Considerations

- Never include API keys, passwords, or sensitive data in code
- Use proper error handling for all file operations
- Validate all input data for text processing functions
- Follow principle of least privilege for external service access

## Commit Guidelines

- Run `go test ./...` before committing
- Run `gofmt -w .` to ensure proper formatting
- Write clear, descriptive commit messages
- Test both success and error paths for new functionality
- Ensure all interfaces are properly implemented
- Maintain separation of concerns between architecture layers
