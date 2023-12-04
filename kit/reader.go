package kit

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func ReadFileFlag() string {
	filePath := flag.String("f", "", "Input file path.")
	flag.Parse()

	if *filePath == "" {
		log.Fatalln("File path flag is required. Usage: -f=<filepath>")
	}

	return *filePath
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
