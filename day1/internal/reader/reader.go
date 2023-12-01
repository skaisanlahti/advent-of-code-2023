package reader

import (
	"bufio"
	"flag"
	"os"
)

func ReadFlags() (bool, string) {
	checkStrings := flag.Bool("s", false, "Check strings for digit values.")
	filePath := flag.String("f", "", "Input file path.")
	flag.Parse()

	if filePath == nil {
		panic("Usage: go run cmd/main.go -s -f=<filepath>")
	}

	return *checkStrings, *filePath
}

func ReadInputFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
