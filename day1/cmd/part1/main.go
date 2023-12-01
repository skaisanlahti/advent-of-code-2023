package main

import (
	"log"
	"os"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/calibrator"
	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/reader"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/part1/main.go <filepath>")
	}

	filePath := os.Args[1]
	input := reader.ReadInputFile(filePath)

	start := time.Now()
	sum := calibrator.Calibrate(input, false)
	duration := time.Since(start).Milliseconds()

	log.Printf("Calibrated %s in %d ms with result %d.", filePath, duration, sum)
}
