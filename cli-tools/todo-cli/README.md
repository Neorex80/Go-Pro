# Todo CLI

A simple command-line todo list manager.

## Features

- Add new todos
- List all todos
- Mark todos as complete
- Delete todos
- Persistent storage (todos are saved to a file)

## Usage

```bash
go run .
```

The application will present a menu with the following options:

1. List todos
2. Add todo
3. Complete todo
4. Delete todo
5. Exit

## Building

To build the executable:

```bash
go build -o todo-cli
```

Then run it:

```bash
./todo-cli
```

## Data Storage

Todos are stored in a file called `todos.txt` in the same directory as the executable. The file format is:

```
ID|STATUS|TASK
```

Where:
- ID is the todo's unique identifier
- STATUS is 1 if completed, 0 if not
- TASK is the todo description
