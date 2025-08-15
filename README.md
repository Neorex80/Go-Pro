# Go Projects Toolkit

A collection of small, practical, and beginner-friendly Go projects — perfect for learning and building a useful toolkit. Includes CLI tools, APIs, and automation scripts.

## Project Structure

- **cli-tools/** - Command-line interface tools
- **web-apis/** - Web API applications
- **automation-scripts/** - Automation and system scripts

## Projects

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

## Getting Started

Each project is self-contained with its own `go.mod` file. Navigate to any project directory and run:

```bash
go run .
```

Or build the project with:

```bash
go build
```

## Prerequisites

- Go 1.16 or higher installed on your system

## Testing

To test any project, navigate to its directory and follow the instructions in its README file.

## GitHub Setup

To set up this repository on GitHub, you have several options:

### Option 1: Using GitHub CLI (Recommended)

If you have GitHub CLI installed, you can use the provided scripts:

**For Unix-like systems (Linux, macOS):**
```bash
./github-setup-gh.sh
```

**For Windows:**
```cmd
github-setup-gh.bat
```

These scripts will:
1. Initialize a Git repository
2. Add and commit all files
3. Create a new GitHub repository
4. Push the code to GitHub

### Option 2: Quick Setup with Predefined Remote

If you want to use the specific repository `https://github.com/Neorex80/Go-Pro.git`:

**For Unix-like systems (Linux, macOS):**
```bash
./setup-remote.sh
```

**For Windows:**
```cmd
setup-remote.bat
```

These scripts will initialize the Git repository, add and commit all files, and set the remote origin to `https://github.com/Neorex80/Go-Pro.git`.

After running the script, push the code to GitHub:
```bash
git push -u origin main
```

### Option 3: Manual Setup

1. Create a new repository on GitHub
2. Clone this project to your local machine
3. Initialize Git in the project directory:
   ```bash
   git init
   ```
4. Add all files to Git:
   ```bash
   git add .
   ```
5. Commit the files:
   ```bash
   git commit -m "Initial commit"
   ```
6. Add your GitHub repository as the remote origin:
   ```bash
   git remote add origin https://github.com/yourusername/your-repo-name.git
   ```
7. Push the code to GitHub:
   ```bash
   git push -u origin main
   ```

### Option 4: Platform-specific Scripts

If you don't have GitHub CLI installed, you can use the platform-specific scripts:

**For Unix-like systems (Linux, macOS):**
```bash
./github-setup.sh
```

**For Windows:**
```cmd
github-setup.bat
```

These scripts will initialize the Git repository and commit the files, but you'll need to manually create the GitHub repository and push the code.

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
