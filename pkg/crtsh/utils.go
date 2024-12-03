package crtsh

import (
	"os"
	"strings"
)

// WriteToFile writes a slice of strings to a file, one per line.
func WriteToFile(filePath string, lines []string) error {
	data := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(data), 0644)
}
