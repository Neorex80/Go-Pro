# File Organizer

A script that organizes files in directories based on their extensions.

## Features

- Automatically organizes files into folders based on their file extensions
- Works on the current directory or a specified directory
- Creates folders for each file extension type
- Moves files to their respective extension folders

## Usage

```bash
go run . [directory]
```

If no directory is specified, it will organize files in the current directory.

### Examples

Organize files in the current directory:
```bash
go run .
```

Organize files in a specific directory:
```bash
go run . /path/to/directory
```

## Building

To build the executable:

```bash
go build -o file-organizer
```

Then run it:

```bash
./file-organizer [directory]
```

## How It Works

1. The script scans the specified directory for files
2. For each file with an extension, it creates a folder named after the extension (if it doesn't already exist)
3. It then moves the file to the appropriate extension folder

## Example

Before running the script:
```
documents/
├── report.pdf
├── image1.jpg
├── image2.jpg
├── data.csv
└── notes.txt
```

After running the script:
```
documents/
├── pdf/
│   └── report.pdf
├── jpg/
│   ├── image1.jpg
│   └── image2.jpg
├── csv/
│   └── data.csv
└── txt/
    └── notes.txt
