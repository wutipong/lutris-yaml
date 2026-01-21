package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExtractTextOutsideBrackets extracts text from a string that is not enclosed within square brackets.
// It returns a string containing the extracted text and an error if the input string is invalid.
func ExtractTextOutsideBrackets(input string) (string, error) {
	// Check for empty string
	if input == "" {
		return "", fmt.Errorf("input string is empty")
	}

	inBracket := false

	builder := strings.Builder{}

	for _, r := range input {
		if r == '[' {
			inBracket = true
		} else if r == ']' && inBracket {
			inBracket = false
		} else if !inBracket {
			builder.WriteRune(r)
		}
	}

	return strings.TrimSpace(builder.String()), nil
}

func GetName(path string, config Config) (name string, err error) {
	gamePath, err := filepath.Abs(config.GamesLocation)
	if err != nil {
		return
	}

	relative, err := filepath.Rel(gamePath, path)
	if err != nil {
		return
	}

	parts := strings.Split(relative, string(os.PathSeparator))

	if len(parts) == 0 {
		err = fmt.Errorf("Invalid path")
		return
	}

	name, err = ExtractTextOutsideBrackets(parts[0])

	return
}
