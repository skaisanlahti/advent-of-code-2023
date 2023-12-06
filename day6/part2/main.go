package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day6/toyboat"
	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func main() {
	start := time.Now()
	filePath := kit.ReadFileFlag()
	input := kit.ReadInputFile(filePath)
	result := toyboat.RaceMargin(input)
	duration := time.Since(start).Milliseconds()
	log.Printf("Processed %s in %d ms with result %d.", filePath, duration, result)
}
