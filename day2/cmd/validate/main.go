package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day2/internal"
	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func main() {
	start := time.Now()
	red, green, blue := internal.ReadColorFlags()
	filePath := kit.ReadFileFlag()
	input := kit.ReadInputFile(filePath)
	result := internal.SumValidGameIds(input, red, green, blue)
	duration := time.Since(start).Milliseconds()
	log.Printf("Processed %s in %d ms with result %d.", filePath, duration, result)
}
