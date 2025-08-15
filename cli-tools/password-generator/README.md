# Password Generator

A secure password generator with customizable options.

## Features

- Generate passwords of custom length
- Include/exclude uppercase letters, digits, and symbols
- Generate multiple passwords at once

## Usage

```bash
go run . [options]
```

### Options

- `-length int`: Length of the password (default 12)
- `-upper bool`: Include uppercase letters (default true)
- `-digits bool`: Include digits (default true)
- `-symbols bool`: Include symbols (default true)
- `-count int`: Number of passwords to generate (default 1)

### Examples

Generate a default password:
```bash
go run .
```

Generate a 16-character password:
```bash
go run . -length=16
```

Generate a password without symbols:
```bash
go run . -symbols=false
```

Generate 5 passwords:
```bash
go run . -count=5
```

Generate a password with only lowercase letters and digits:
```bash
go run . -upper=false -symbols=false
```

## Building

To build the executable:

```bash
go build -o password-generator
```

Then run it:

```bash
./password-generator [options]
