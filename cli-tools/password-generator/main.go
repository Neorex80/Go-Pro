package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func main() {
	// Define command-line flags
	length := flag.Int("length", 12, "Length of the password")
	includeUpper := flag.Bool("upper", true, "Include uppercase letters")
	includeDigits := flag.Bool("digits", true, "Include digits")
	includeSymbols := flag.Bool("symbols", true, "Include symbols")
	count := flag.Int("count", 1, "Number of passwords to generate")

	flag.Parse()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate passwords
	for i := 0; i < *count; i++ {
		password := generatePassword(*length, *includeUpper, *includeDigits, *includeSymbols)
		fmt.Println(password)
	}
}

func generatePassword(length int, includeUpper, includeDigits, includeSymbols bool) string {
	// Build character set
	charset := lowercase
	if includeUpper {
		charset += uppercase
	}
	if includeDigits {
		charset += digits
	}
	if includeSymbols {
		charset += symbols
	}

	// Generate password
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
