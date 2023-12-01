package reader

import (
	"bufio"
	"log"
	"os"
)

func ReadInputFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}

	log.Println("Input:")
	log.Println("---")
	for _, line := range lines {
		log.Println(line)
	}
	log.Println("---")

	return lines
}
