package main

import (
	"log"
	"os"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/calibrator"
	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/reader"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		panic("Usage: go run cmd/part1/main.go <filepath>")
	}

	filePath := os.Args[1]
	input := reader.ReadInputFile(filePath)

	result := calibrator.Calibrate(input, false)
	duration := time.Since(start).Nanoseconds()

	log.Printf("Calibrated %s in %d ns with result %d.", filePath, duration, result)
}
