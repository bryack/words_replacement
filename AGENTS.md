# AGENTS.md

This file provides context and instructions for AI coding agents working on the Words project.

## Project Overview

Words is a Go CLI application implementing hexagonal architecture for intelligent word form replacement. It demonstrates clean architecture principles with dependency injection, interface-driven design, comprehensive testing strategies, and production-ready data integration with SQLite and Kaikki.org linguistic datasets.

**Current Implementation Status:**
- ✅ SQLite database integration with persistent file-based storage and DataLoader pattern
- ✅ Kaikki.org JSONL data loading (334MB+ Russian noun dataset) with smart loading detection
- ✅ Multi-case form extraction (nominative, accusative, genitive, dative, instrumental, prepositional)
- ✅ Unicode-aware word replacement with Cyrillic boundary detection and stress mark removal
- ✅ Cobra CLI framework with flag-based configuration (--input, --data, --old, --new)
- ✅ FormsProvider interface with SQLite and Wiktionary implementations
- ✅ ATDD methodology with acceptance tests using binary execution
- ✅ Comprehensive benchmarking for performance optimization
- ✅ Contract-based testing with WordReplacerCLI interface
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
- Dependencies: 
  - `github.com/alecthomas/assert/v2` for enhanced test assertions
  - `github.com/mattn/go-sqlite3` for SQLite database operations
  - `github.com/spf13/cobra` for CLI framework
- Target platform: Linux (primary), cross-platform compatible

## Project Structure

- `cmd/cli/` - CLI application entry point with main.go and acceptance tests
- `internal/` - Core business logic (private packages)
  - `internal/cli/` - CLI implementation with Cobra commands and CLI struct
  - `internal/replacer/` - Word replacement logic with FormsProvider interface
- `adapters/` - External interface implementations
  - `adapters/sqlite/` - SQLite database provider with JSONL data loading and DataLoader pattern
    - `models.go` - JSONL data structures (KaikkiEntry, WordForm)
    - `extractor.go` - Multi-case form extraction with stress removal
    - `loader.go` - JSONL file parsing with LoadFromJSONL function
    - `provider.go` - SQLiteFormsProvider with persistent database
  - `adapters/wiktionary/` - MediaWiki API client with provider and parser
  - `adapters/cli/` - CLI driver implementing WordReplacerCLI interface
  - `adapters/acceptance/` - Acceptance test driver implementing WordReplacer interface
- `contracts/` - Interface definitions (WordReplacerCLI) and contract tests
- `specifications/` - Behavior-driven specifications (WordReplacer, WiktionaryFormsProvider)
- `testhelpers/` - Test utility functions for file operations
- `wiki/` - MediaWiki API client package with WikiClient
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
