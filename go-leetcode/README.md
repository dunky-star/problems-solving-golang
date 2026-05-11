# Go LeetCode

A Go problem-solving repository focused on LeetCode-style exercises, core algorithms, and practical Go fundamentals.

## Overview

This repository contains small, focused Go programs for practicing data structures, control flow, error handling, and language concepts commonly used in coding interviews.

## Structure

```text
├── array_max_num.go              # Find the maximum value in an integer slice
├── count_blue_red.go             # Count matching strings in a slice
├── count_capital_letters.go      # Count uppercase letters in a string
├── two_dim_array1.go             # Traverse a 2D array and find the maximum value
├── go-pointer.go                 # Pointers, structs, slices, maps, and sorting
├── return_values.go              # Errors, generics, interfaces, embedding, and goroutines
├── read_log_file.go              # Read and filter log files
├── main.go                       # Basic HTTP server example
├── home.html                     # Static HTML page for the web example
├── log.txt                       # Sample log data
└── numbers.txt                   # Sample numeric data
```

## Topics

- Arrays and slices
- Strings and counting
- Pointers and structs
- Maps and sorting
- Errors and multiple return values
- Interfaces and generics
- File reading
- Basic HTTP handling

## Requirements

- Go 1.21+

## Running

Run a standalone example:

```bash
go run go-pointer.go
```

Some files use renamed entry functions such as `mainMaxNum`, `mainRedBlue`, or `mainBkup`. To run one, call it from the active `main` function or temporarily rename it to `main`.

## Notes

This project is intentionally lightweight and currently has no `go.mod` file. It is designed as a hands-on workspace for practicing Go problem solving before expanding into structured packages and tests.

## License

No license specified.
