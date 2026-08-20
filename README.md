# System Info Collector

![CI](https://github.com/Qyroxen/System-Info-Collector/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/System-Info-Collector?style=social)

> A powerful CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/System-Info-Collector?style=social)](https://github.com/Qyroxen/System-Info-Collector/stargazers)

## What is it?

System Info Collector is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/System-Info-Collector.git
cd System-Info-Collector
go build -o systeminfocollector .

# Run
./systeminfocollector --help
```

## CLI Usage

```bash
# Basic usage
./systeminfocollector

# With flags
./systeminfocollector --verbose --output json

# Get help
./systeminfocollector --help
```

## Examples

```bash
# Example 1
./systeminfocollector example1

# Example 2
./systeminfocollector example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o systeminfocollector .

# Lint
go vet ./...
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/System-Info-Collector/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/System-Info-Collector?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/System-Info-Collector/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/System-Info-Collector?style=social" alt="Fork this repo">
  </a>
</p>
