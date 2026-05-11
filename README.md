# Problems Solving Go

A Go workspace for practicing problem solving, language fundamentals, concurrency, testing, debugging, database access, and web services.

## Architecture

```text
├── go-leetcode/           # LeetCode-style exercises and algorithm practice
├── go-reviews/            # Go syntax, functions, collections, errors, and interfaces
├── go-testing/            # Unit testing examples
├── go-debugging/          # Debugging, logs, calculations, and web server examples
├── go-concurrency/        # Goroutines, channels, and wait groups
├── concurrency-usecase/   # Practical concurrency workflow example
├── rdbms-go/              # SQL and relational database examples
├── web-services/          # Go web service project structure
└── certs/                 # Local certificate assets
```

## Focus

Go LeetCode — Core problem-solving exercises for arrays, strings, counting, traversal, pointers, and Go fundamentals.

Go Reviews — Quick reference examples for common Go language features.

Go Testing — Basic test-driven Go practice with `go test`.

Go Concurrency — Hands-on examples using goroutines, channels, and wait groups.

RDBMS Go — Database access and SQL practice.

Web Services — Structured Go web service examples.

## Requirements

- Go 1.21+

## Running

Run a module-based project:

```bash
cd go-testing
go test ./...
```

Run a standalone practice file:

```bash
cd go-leetcode
go run go-pointer.go
```

## Notes

This repository is a learning and practice workspace. Some directories are full Go modules with `go.mod`; others contain standalone examples intended for quick experimentation.

## License

No license specified.
