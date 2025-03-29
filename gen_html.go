package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.json [output.html] [js_highlighter.js]")
		os.Exit(1)
	}

	// Get input and output file paths
	inputFile := os.Args[1]

	var outputFile string
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	} else {
		baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
		outputFile = baseName + ".html"
	}

	// Get optional JS highlighter file
	var jsHighlighterContent string
	if len(os.Args) >= 4 {
		jsFile := os.Args[3]
		jsBytes, err := os.ReadFile(jsFile)
		if err != nil {
			fmt.Printf("Warning: Could not read JS highlighter file: %s\n", err)
		} else {
			jsHighlighterContent = string(jsBytes)
		}
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
        body { font-family: Arial, sans-serif; margin: 20px; }
        .entry { margin: 10px 0; padding: 10px; border: 1px solid #ccc; background-color: #f8f8f8; }
        .commit-link { margin-bottom: 10px; }
        .commit-link a { font-family: monospace; text-decoration: none; }
        pre { background-color: #fff; padding: 10px; border: 1px solid #ddd; overflow: auto; }
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
		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(line), &jsonData); err != nil {
			fmt.Printf("Skipping invalid JSON: %s...\n", line[:min(50, len(line))])
			continue
		}

		// Start entry div
		outFile.WriteString("<div class='entry'>\n")

		// Add commit link if available
		commitHash, repoName := getCommitHash(jsonData)
		if commitHash != "" {
			var commitLink string
			if repoName != "" {
				commitLink = fmt.Sprintf("https://github.com/%s/commit/%s", repoName, commitHash)
			} else {
				// Fallback if we couldn't parse the repo properly
				commitLink = fmt.Sprintf("https://github.com/search?q=%s", commitHash)
			}
			outFile.WriteString(fmt.Sprintf("  <div class='commit-link'><a href='%s' target='_blank'>[commit]</a></div>\n", commitLink))
		}

		// Format and write the JSON
		jsonBytes, _ := json.MarshalIndent(jsonData, "", "  ")
		outFile.WriteString(fmt.Sprintf("  <pre>%s</pre>\n", jsonBytes))

		// End entry div
		outFile.WriteString("</div>\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	// Write HTML footer with optional JS syntax highlighter
	if jsHighlighterContent != "" {
		outFile.WriteString(fmt.Sprintf("<script>\n%s\n</script>\n", jsHighlighterContent))
	}
	outFile.WriteString("</body>\n</html>")

	fmt.Printf("Created %s\n", outputFile)
}

// getCommitHash extracts the git commit hash and repository from the JSON data
// Returns commit hash and GitHub URL format of the repo (owner/repo)
func getCommitHash(data map[string]interface{}) (string, string) {
	var commit, repo string

	// Get the nested maps, returning early if any are missing
	sourceMetadata, ok := data["SourceMetadata"].(map[string]interface{})
	if !ok {
		return "", ""
	}

	dataField, ok := sourceMetadata["Data"].(map[string]interface{})
	if !ok {
		return "", ""
	}

	git, ok := dataField["Git"].(map[string]interface{})
	if !ok {
		return "", ""
	}

	// Get commit hash
	if commitVal, ok := git["commit"].(string); ok {
		commit = commitVal
	}

	// Get repository if available
	repoVal, ok := git["repository"].(string)
	if !ok {
		return commit, ""
	}

	// Extract repository info
	if strings.Contains(repoVal, ":") {
		// Format like git@server:path/repo.git
		parts := strings.Split(repoVal, ":")
		if len(parts) <= 1 {
			return commit, ""
		}

		repoPath := parts[len(parts)-1]
		repoPath = strings.TrimPrefix(repoPath, "/")
		repoPath = strings.TrimSuffix(repoPath, ".git")

		// Split path to get owner/repo
		pathParts := strings.Split(repoPath, "/")
		if len(pathParts) >= 2 {
			// Extract the final two parts as owner/repo
			repo = pathParts[len(pathParts)-2] + "/" + pathParts[len(pathParts)-1]
		}
	} else if strings.Contains(repoVal, "/") {
		// Other format with slashes
		parts := strings.Split(repoVal, "/")
		if len(parts) >= 2 {
			repo = parts[len(parts)-2] + "/" + strings.TrimSuffix(parts[len(parts)-1], ".git")
		}
	}

	return commit, repo
}

// min returns the smaller of x or y
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
