# CCWC - Word Count Utility

A lightweight word count utility written in Go, inspired by the Unix `wc` command. Count lines, words, bytes, and characters in files or from standard input.

## Features

- Count lines, words, bytes, and characters
- Read from files or standard input (piped input)
- Command-line flags for specific count types
- Cross-platform support (Linux, macOS, Windows)

## Installation

### Prerequisites

- Go 1.23 or higher

### Build from Source

```bash
cd learning/BYO/go/ccwc
go build .
```

This will create an executable file named `ccwc` in the current directory.

## Usage

### Basic Usage

Count all statistics (lines, words, bytes, characters):

```bash
./ccwc <filename>
```

Example:
```bash
./ccwc test1.txt
```

Output:
```
10 25 156 150 test1.txt
```

### Flags

Use flags to count specific metrics:

- `-l` : Count lines only
- `-w` : Count words only
- `-c` : Count bytes only
- `-m` : Count characters only

### Examples

Count lines in a file:
```bash
./ccwc -l test1.txt
```

Count words in a file:
```bash
./ccwc -w test1.txt
```

Count bytes in a file:
```bash
./ccwc -c test1.txt
```

Count characters in a file:
```bash
./ccwc -m test1.txt
```

### Reading from Standard Input

You can pipe input directly to `ccwc`:

```bash
echo "My name is John" | ./ccwc
```

Count words from piped input:
```bash
echo "My name is John" | ./ccwc -w
```

Using `cat` to pipe a file:
```bash
cat test1.txt | ./ccwc -l
```

## Project Structure

```
ccwc/
├── go.mod              # Go module definition
├── go.sum              # Go dependencies
├── main.go             # Command-line interface
└── scanner/
    └── scanner.go      # Core scanning logic
```

## Output Format

By default, the output includes all statistics followed by the filename:

```
<lines> <words> <bytes> <characters> <filename>
```

When reading from stdin, the filename field is empty:

```
<lines> <words> <bytes> <characters>
```

With flags, only the requested metric is displayed:

```
<count> <filename>
```

## Module Information

- **Module Path**: `github.com/P-H-Pancholi/BYO/go/ccwc`
- **Go Version**: 1.23.4+

## License

This project is part of the Build Your Own series and is available for educational purposes.

## Related

This project is inspired by the Unix `wc` command and the Build Your Own Wc Tool challenge at [codingchallenges](https://codingchallenges.fyi/).