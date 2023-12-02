package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day2/internal"
)

func main() {
	start := time.Now()
	filePath := internal.ReadCalculatorFlags()
	input := internal.ReadInputFile(filePath)
	result := internal.SumGamePowers(input)
	duration := time.Since(start).Nanoseconds()
	log.Printf("Calculated %s powers in %d ns with result %d.", filePath, duration, result)
}
