package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day2/internal"
)

func main() {
	start1 := time.Now()
	red, green, blue, filePath := internal.ReadFlags()
	input := internal.ReadInputFile(filePath)
	result1 := internal.SumValidGameIds(input, red, green, blue)
	duration1 := time.Since(start1).Nanoseconds()
	log.Printf("Validated %s games in %d ns with result %d.", filePath, duration1, result1)

	start2 := time.Now()
	result2 := internal.SumGamePowers(input)
	duration2 := time.Since(start2).Nanoseconds()
	log.Printf("Calculated %s powers in %d ns with result %d.", filePath, duration2, result2)
}
