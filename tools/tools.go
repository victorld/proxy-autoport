package tools

import (
	"bufio"
	"encoding/json"
	"os"
)

func ReadFileLine(lineNumber int, filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	return ""
}

func MarshalPrint(v any) string {
	sm, _ := json.Marshal(v)
	return string(sm)
}
