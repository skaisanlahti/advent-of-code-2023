package internal

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func ReadFlags() (int, int, int, string) {
	r := flag.Int("r", 0, "Number of red cubes.")
	g := flag.Int("g", 0, "Number of green cubes.")
	b := flag.Int("b", 0, "Number of blue cubes.")
	filePath := flag.String("f", "", "Input file path.")
	flag.Parse()

	if *filePath == "" {
		log.Fatalln("File path flag is required. Usage: -f=<filepath>")
	}

	return *r, *g, *b, *filePath
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
