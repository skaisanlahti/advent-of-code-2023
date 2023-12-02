package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day2/internal"
)

func main() {
	start := time.Now()
	red, green, blue, filePath := internal.ReadValidatorFlags()
	input := internal.ReadInputFile(filePath)
	result := internal.SumValidGameIds(input, red, green, blue)
	duration := time.Since(start).Milliseconds()
	log.Printf("Validated %s games in %d ms with result %d.", filePath, duration, result)
}
