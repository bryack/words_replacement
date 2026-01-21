# Words

A Go CLI application for intelligent word form replacement implementing hexagonal architecture with clean separation of concerns and production-ready linguistic data integration.

## Features

- **Intelligent Word Replacement**: Context-aware Russian noun form transformations with multi-case support (nominative, accusative, genitive, dative, instrumental, prepositional)
- **Clean Architecture**: Hexagonal architecture with dependency injection and FormsProvider interface
- **SQLite Database**: Persistent file-based database with DataLoader pattern for fast word form lookups
- **Kaikki.org Integration**: Production linguistic data from 334MB+ Kaikki.org JSONL dictionary
- **Cobra CLI Framework**: Command-line interface with flag-based configuration (--input, --data, --old, --new)
- **Multiple Data Sources**: Support for SQLite, MediaWiki API, and custom providers through FormsProvider interface
- **Unicode-Aware Processing**: Handles Cyrillic text boundaries and automatic stress mark removal
- **Comprehensive Testing**: Unit, integration, and acceptance test coverage with ATDD methodology and contract-based testing

## Installation

```bash
go build -o words ./cmd/cli
```

## Usage

```bash
# Replace word forms in a text file
./words replace --input input.txt --old подделка --new fake

# With custom data file
./words replace --input input.txt --data custom.jsonl --old подделка --new fake
```

### Business Rule Examples

**Core Replacement Logic:**
```
Input: "У меня есть подделка" 
Command: ./words replace --input text.txt --old подделка --new fake
Output: "У меня есть fake" 
```

**Error Handling:**
```bash
# Missing word in dictionary
./words replace --input text.txt --old несуществующееслово --new fake
# Returns: Error: word "несуществующееслово" not found in dictionary
# No replacement performed - this is correct behavior
```

**Explicit Parameters (No Hardcoded Paths):**
```bash
# ✅ Correct: All inputs via flags
./words replace --input /path/to/input.txt --data /path/to/data.jsonl --old word1 --new word2

# ❌ Wrong: Hardcoded paths not supported
./words replace word1 word2  # This won't work
```

## Architecture

The project demonstrates hexagonal architecture principles:

- `cmd/cli/` - Application entry point with main.go and acceptance tests
- `internal/` - Core business logic (CLI struct, Replacer with FormsProvider interface)
- `adapters/` - External interface implementations
  - `adapters/sqlite/` - SQLiteFormsProvider with DataLoader pattern and JSONL parsing
  - `adapters/wiktionary/` - MediaWiki API client implementing FormsProvider
  - `adapters/cli/` - CLI driver implementing WordReplacerCLI interface
  - `adapters/acceptance/` - Acceptance test driver implementing WordReplacer interface
- `contracts/` - Interface definitions (WordReplacerCLI) and contract tests
- `specifications/` - Behavior-driven specifications (WordReplacer, WiktionaryFormsProvider)
- `testhelpers/` - Test utility functions for file operations
- `wiki/` - MediaWiki API client integration with WikiClient

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
- `github.com/spf13/cobra` (CLI framework)

## License

See project license file.
