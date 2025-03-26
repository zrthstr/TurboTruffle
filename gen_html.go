package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// processEntry converts a JSON object to HTML representation
func processEntry(entry map[string]interface{}) string {
	// Simple representation - use json.MarshalIndent to get formatted JSON
	jsonBytes, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return fmt.Sprintf("<div class='error'>Error processing entry: %s</div>\n", err)
	}

	return fmt.Sprintf("<div class='entry'>\n  <pre>%s</pre>\n</div>\n", jsonBytes)
}

func main() {
	// Check arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.json [output.html]")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Determine output file - either from argument or derived from input filename
	var outputFile string
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	} else {
		baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
		outputFile = baseName + ".html"
	}

	// Open input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %s\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Write HTML header
	htmlHeader := `<!DOCTYPE html>
<html>
<head>
    <style>
        .entry {
            margin: 10px;
            padding: 10px;
            border: 1px solid #ccc;
        }
        .error {
            color: red;
            margin: 10px;
            padding: 10px;
            border: 1px solid red;
        }
    </style>
</head>
<body>
`
	outFile.WriteString(htmlHeader)

	// Process file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Parse JSON
		var entry map[string]interface{}
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			fmt.Printf("Skipping invalid JSON: %s...\n", line[:min(50, len(line))])
			continue
		}

		// Process entry and write to output file
		outFile.WriteString(processEntry(entry))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// Write HTML footer
	outFile.WriteString("</body>\n</html>")

	fmt.Printf("Created %s\n", outputFile)
}

// min returns the smaller of x or y
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
