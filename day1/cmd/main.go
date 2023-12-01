package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/calibrator"
	"github.com/skaisanlahti/advent-of-code-2023/day1/internal/reader"
)

func main() {
	start := time.Now()
	checkStrings, filePath := reader.ReadFlags()
	input := reader.ReadInputFile(filePath)
	result := calibrator.Calibrate(input, checkStrings)
	duration := time.Since(start).Nanoseconds()
	log.Printf("Calibrated %s in %d ns with result %d.", filePath, duration, result)
}
