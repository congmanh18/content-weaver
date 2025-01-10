package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func ProcessFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chapterRegex := regexp.MustCompile(`(?i)^Chapter \d+:`) // Ví dụ: Chapter 1:
	for scanner.Scan() {
		line := scanner.Text()
		if chapterRegex.MatchString(line) {
			fmt.Printf("Found chapter: %s\n", line)
		}
	}
}
