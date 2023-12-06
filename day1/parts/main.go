package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day1/trebuchet"
	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func main() {
	start := time.Now()
	checkStringPatterns, filePath := trebuchet.ReadFlags()
	input := kit.ReadInputFile(filePath)
	result := trebuchet.SumCalibrationValues(input, checkStringPatterns)
	duration := time.Since(start).Milliseconds()
	log.Printf("Processed %s in %d ms with result %d.", filePath, duration, result)
}
