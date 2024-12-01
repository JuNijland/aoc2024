package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadInputFile reads a file and returns its contents as a string slice
func ReadInputFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadInputFileRaw reads a file and returns its contents as a single string
func ReadInputFileRaw(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(content))
}
