# System Info Collector

Collect system information for diagnostics and monitoring.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/system-info-collector/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/system-info-collector/actions/workflows/ci.yml)

> Collect system information for diagnostics and monitoring.

## What is it?

System Info Collector is a command-line tool built with Go that helps developers collect system information for diagnostics and monitoring. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs system info collector — but existing tools are either too complex, too slow, or require cloud dependencies. We built System Info Collector to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- CPU/Memory/Disk info
- Network configuration
- Process information
- Hardware details
- Export to JSON/CSV
- CLI interface

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/system-info-collector@latest

# Or build from source
git clone https://github.com/Qyroxen/system-info-collector.git
cd system-info-collector
go build -o system-info-collector .
```

### Usage

```bash
# Basic usage
.system-info-collector --help

# Example
./system-info-collector collect --output system-info.json
```

## Output

```
System Info Collector v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
system info collector [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.system-info-collector --path ./src
```

### Advanced Example

```bash
.system-info-collector --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run System Info Collector
  run: |
    go install github.com/Qyroxen/system-info-collector@latest
    system-info-collector --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
