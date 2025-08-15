# System Monitor

A script that monitors system resources (CPU, memory) and logs them.

## Features

- Monitors CPU usage
- Monitors memory usage
- Monitors disk usage
- Displays real-time system information
- Cross-platform support (Windows, Linux, macOS)
- Configurable monitoring interval

## Usage

```bash
go run . [interval]
```

The interval is specified in seconds. If no interval is provided, it defaults to 5 seconds.

### Examples

Monitor every 5 seconds (default):
```bash
go run .
```

Monitor every 10 seconds:
```bash
go run . 10
```

## Building

To build the executable:

```bash
go build -o system-monitor
```

Then run it:

```bash
./system-monitor [interval]
```

## Output Format

The script will continuously display system information in the following format:

```
[2023-05-15 14:30:22]
CPU Usage:    25%
Memory Usage: 60%
Disk Usage:   45%
----------------------------------------
```

Press `Ctrl+C` to stop the monitoring.

## How It Works

The script uses system-specific commands to gather resource information:

### Windows
- CPU usage: `wmic cpu get loadpercentage`
- Memory usage: `wmic OS get FreePhysicalMemory,TotalVisibleMemorySize`
- Disk usage: `wmic logicaldisk where caption='C:' get Size,FreeSpace`

### Linux/macOS
- CPU usage: `top -bn1` or `sar -u 1 1`
- Memory usage: `free -m`
- Disk usage: `df -h /`

## Notes

- The accuracy of the information may vary depending on the system
- Some commands may require appropriate permissions
- The script is designed to be lightweight and not significantly impact system performance
