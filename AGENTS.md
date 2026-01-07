# AGENTS.md

This file provides context and instructions for AI coding agents working on the Words project.

## Project Overview

Words is a Go utility for text processing and wiki content retrieval. It performs regex-based text replacement operations and provides a MediaWiki API client for content extraction from wiki sources.

## Build and Test Commands

### Build
```bash
go build -o words
```

### Test
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

## Code Style Guidelines

- Follow standard Go conventions and `gofmt` formatting
- Use PascalCase for exported functions (`Replace`, `ReadAndReplace`)
- Use camelCase for local variables, PascalCase for exported variables
- Use UPPER_SNAKE_CASE for package-level constants
- Test files must end with `_test.go` and be placed alongside source files

## Testing Instructions

- Use Go's built-in `testing` package
- Write table-driven tests for multiple test cases
- Test error conditions and edge cases
- Aim for high test coverage of public functions
- Use `io/fs` interfaces for file system mocking to maintain zero dependencies

## Development Environment

- Go version 1.25.5 required
- Module: `github.com/bryack/words`
- Zero external dependencies policy - use only standard library
- Target platform: Linux (primary), cross-platform compatible

## Project Structure

- `words.go` - Main text processing logic and CLI entry point
- `wiki/` - MediaWiki API client package
- `*_test.go` - Test files alongside source files
- `.kiro/steering/` - AI agent steering documents

## Security Considerations

- Never include API keys, passwords, or sensitive data in code
- Use proper error handling for all file operations
- Validate all input data for text processing functions

## Commit Guidelines

- Run `go test ./...` before committing
- Run `gofmt -w .` to ensure proper formatting
- Write clear, descriptive commit messages
- Test both success and error paths for new functionality
