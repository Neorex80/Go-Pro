# File Renamer

A simple command-line tool to batch rename files with various patterns.

## Features

- Add prefixes to filenames
- Add suffixes to filenames
- Replace strings in filenames
- Process files in a specified directory

## Usage

```bash
go run . [options]
```

### Options

- `-prefix string`: Prefix to add to filenames
- `-suffix string`: Suffix to add to filenames
- `-replace string`: String to replace in filenames
- `-with string`: String to replace with in filenames
- `-dir string`: Directory to process (default ".")

### Examples

Add a prefix to all files in the current directory:
```bash
go run . -prefix="new_"
```

Add a suffix to all files in a specific directory:
```bash
go run . -suffix="_backup" -dir="/path/to/directory"
```

Replace spaces with underscores in all files:
```bash
go run . -replace=" " -with="_"
```

Combine multiple operations:
```bash
go run . -prefix="img_" -replace=" " -with="_" -suffix="_v2"
```

## Building

To build the executable:

```bash
go build -o file-renamer
```

Then run it:

```bash
./file-renamer [options]
