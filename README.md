# 🌐 System Info Collector

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Infrastructure tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`infrastructure` `devops` `cli` `golang`

---

## What is System-Info-Collector?

**System-Info-Collector** is an infrastructure tool for monitoring, inspecting, and managing systems and services.

## Features

- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/System-Info-Collector.git
cd System-Info-Collector

# Build
go build -o system-info-collector .

# Run
./system-info-collector [file]
```

### Or directly with `go run`:
```bash
go run main.go [file]
```

## Usage

```bash
# Basic usage
./system-info-collector [file]
```

### Example Output

```
$ ./system-info-collector [file]
System Info Collector
====================
OS:         %s\n
```

## Project Structure

```
System-Info-Collector/
  main.go          # Entry point (29 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
