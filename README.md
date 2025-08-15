# Go Projects Toolkit

![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)

A collection of small, practical, and beginner-friendly Go projects — perfect for learning and building a useful toolkit. Includes CLI tools, APIs, and automation scripts.

## 📁 Project Structure

- **cli-tools/** - Command-line interface tools
- **web-apis/** - Web API applications
- **automation-scripts/** - Automation and system scripts

## 🚀 Projects

### CLI Tools

1. **File Renamer** - Batch rename files with various patterns
   - Add prefixes, suffixes, or replace text in filenames
   - Process files in a specified directory

2. **Password Generator** - Secure password generator with customizable options
   - Generate passwords of custom length
   - Include/exclude uppercase letters, digits, and symbols
   - Generate multiple passwords at once

3. **Todo CLI** - Simple command-line todo list manager
   - Add new todos
   - List all todos
   - Mark todos as complete
   - Delete todos
   - Persistent storage (todos are saved to a file)

### Web APIs

1. **Simple REST API** - Basic RESTful API with CRUD operations
   - Create, Read, Update, and Delete items
   - JSON data format
   - RESTful endpoints

2. **Weather API Client** - Client that fetches and displays weather data
   - Fetches current weather data for a specified city
   - Displays temperature, weather condition, humidity, and wind speed
   - Simple command-line interface

### Automation Scripts

1. **File Organizer** - Organizes files based on their extensions
   - Automatically organizes files into folders based on their file extensions
   - Works on the current directory or a specified directory
   - Creates folders for each file extension type

2. **System Monitor** - Monitors system resources (CPU, memory)
   - Monitors CPU usage
   - Monitors memory usage
   - Monitors disk usage
   - Displays real-time system information
   - Cross-platform support (Windows, Linux, macOS)

## 🏁 Getting Started

Each project is self-contained with its own `go.mod` file. Navigate to any project directory and run:

```bash
go run .
```

Or build the project with:

```bash
go build
```

## 📋 Prerequisites

- Go 1.16 or higher installed on your system

## 🧪 Testing

To test any project, navigate to its directory and follow the instructions in its README file.

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on how to contribute to this project.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
