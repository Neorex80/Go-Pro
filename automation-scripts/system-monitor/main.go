package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type SystemInfo struct {
	CPUUsage    string
	MemoryUsage string
	DiskUsage   string
	Timestamp   string
}

func main() {
	// Check if interval is provided
	interval := 5 // default interval in seconds
	if len(os.Args) > 1 {
		fmt.Sscanf(os.Args[1], "%d", &interval)
	}

	fmt.Printf("System Monitor started. Press Ctrl+C to stop.\n")
	fmt.Printf("Monitoring every %d seconds...\n\n", interval)

	// Monitor system resources
	for {
		info, err := getSystemInfo()
		if err != nil {
			log.Printf("Error getting system info: %v", err)
			continue
		}

		// Display system information
		fmt.Printf("[%s]\n", info.Timestamp)
		fmt.Printf("CPU Usage:    %s\n", info.CPUUsage)
		fmt.Printf("Memory Usage: %s\n", info.MemoryUsage)
		fmt.Printf("Disk Usage:   %s\n", info.DiskUsage)
		fmt.Println(strings.Repeat("-", 40))

		// Wait for the specified interval
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func getSystemInfo() (SystemInfo, error) {
	var info SystemInfo
	info.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	var err error

	// Get CPU usage
	info.CPUUsage, err = getCPUUsage()
	if err != nil {
		info.CPUUsage = "N/A"
	}

	// Get memory usage
	info.MemoryUsage, err = getMemoryUsage()
	if err != nil {
		info.MemoryUsage = "N/A"
	}

	// Get disk usage
	info.DiskUsage, err = getDiskUsage()
	if err != nil {
		info.DiskUsage = "N/A"
	}

	return info, nil
}

func getCPUUsage() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsCPUUsage()
	case "linux", "darwin":
		return getUnixCPUUsage()
	default:
		return "N/A", fmt.Errorf("unsupported OS")
	}
}

func getMemoryUsage() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsMemoryUsage()
	case "linux", "darwin":
		return getUnixMemoryUsage()
	default:
		return "N/A", fmt.Errorf("unsupported OS")
	}
}

func getDiskUsage() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsDiskUsage()
	case "linux", "darwin":
		return getUnixDiskUsage()
	default:
		return "N/A", fmt.Errorf("unsupported OS")
	}
}

func getWindowsCPUUsage() (string, error) {
	// This is a simplified approach for Windows
	// In a real application, you might want to use a more accurate method
	cmd := exec.Command("wmic", "cpu", "get", "loadpercentage")
	output, err := cmd.Output()
	if err != nil {
		return "N/A", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		// Find the line with the actual value
		for _, line := range lines[1:] {
			line = strings.TrimSpace(line)
			if line != "" && line != "LoadPercentage" {
				return line + "%", nil
			}
		}
	}

	return "N/A", nil
}

func getWindowsMemoryUsage() (string, error) {
	cmd := exec.Command("wmic", "OS", "get", "FreePhysicalMemory,TotalVisibleMemorySize", "/Value")
	output, err := cmd.Output()
	if err != nil {
		return "N/A", err
	}

	lines := strings.Split(string(output), "\n")
	var freeMemory, totalMemory int64

	for _, line := range lines {
		if strings.HasPrefix(line, "FreePhysicalMemory=") {
			fmt.Sscanf(line, "FreePhysicalMemory=%d", &freeMemory)
		} else if strings.HasPrefix(line, "TotalVisibleMemorySize=") {
			fmt.Sscanf(line, "TotalVisibleMemorySize=%d", &totalMemory)
		}
	}

	if totalMemory > 0 {
		usedMemory := totalMemory - freeMemory
		percentage := float64(usedMemory) / float64(totalMemory) * 100
		return fmt.Sprintf("%.1f%%", percentage), nil
	}

	return "N/A", nil
}

func getWindowsDiskUsage() (string, error) {
	cmd := exec.Command("wmic", "logicaldisk", "where", "caption='C:'", "get", "Size,FreeSpace", "/Value")
	output, err := cmd.Output()
	if err != nil {
		return "N/A", err
	}

	lines := strings.Split(string(output), "\n")
	var totalSize, freeSpace int64

	for _, line := range lines {
		if strings.HasPrefix(line, "Size=") {
			fmt.Sscanf(line, "Size=%d", &totalSize)
		} else if strings.HasPrefix(line, "FreeSpace=") {
			fmt.Sscanf(line, "FreeSpace=%d", &freeSpace)
		}
	}

	if totalSize > 0 {
		usedSpace := totalSize - freeSpace
		percentage := float64(usedSpace) / float64(totalSize) * 100
		return fmt.Sprintf("%.1f%%", percentage), nil
	}

	return "N/A", nil
}

func getUnixCPUUsage() (string, error) {
	// For Unix-like systems, we'll use a simple approach
	// A more accurate implementation would parse /proc/stat on Linux
	cmd := exec.Command("top", "-bn1")
	output, err := cmd.Output()
	if err != nil {
		// Try alternative command
		cmd = exec.Command("sar", "-u", "1", "1")
		output, err = cmd.Output()
		if err != nil {
			return "N/A", err
		}
	}

	// This is a simplified parsing
	// A real implementation would need more robust parsing
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Cpu(s)") {
			// Extract CPU usage from the line
			// This is highly dependent on the output format
			parts := strings.Fields(line)
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}

	return "N/A", nil
}

func getUnixMemoryUsage() (string, error) {
	cmd := exec.Command("free", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "N/A", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		// Parse the memory line
		fields := strings.Fields(lines[1])
		if len(fields) > 2 {
			total, used := fields[1], fields[2]
			return fmt.Sprintf("Used: %s MB / Total: %s MB", used, total), nil
		}
	}

	return "N/A", nil
}

func getUnixDiskUsage() (string, error) {
	cmd := exec.Command("df", "-h", "/")
	output, err := cmd.Output()
	if err != nil {
		return "N/A", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		if len(fields) > 4 {
			return fmt.Sprintf("Used: %s / Available: %s (%s)", fields[2], fields[3], fields[4]), nil
		}
	}

	return "N/A", nil
}
