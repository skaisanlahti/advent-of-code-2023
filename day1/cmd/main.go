package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day1/internal"
)

func main() {
	start := time.Now()
	checkStringPatterns, filePath := internal.ReadFlags()
	input := internal.ReadInputFile(filePath)
	result := internal.SumCalibrationValues(input, checkStringPatterns)
	duration := time.Since(start).Milliseconds()
	log.Printf("Calibrated %s in %d ms with result %d.", filePath, duration, result)
}
