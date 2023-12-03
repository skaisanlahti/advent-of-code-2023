package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day3/internal"
)

func main() {
	start := time.Now()
	filePath := internal.ReadFileFlag()
	input := internal.ReadInputFile(filePath)
	result := internal.SumGearRatios(input)
	duration := time.Since(start).Milliseconds()
	log.Printf("Calculated %s gear in %d ms with result %d.", filePath, duration, result)
}
