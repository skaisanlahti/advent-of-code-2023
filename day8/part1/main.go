package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day8/desert"
	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func main() {
	start := time.Now()
	filePath := kit.ReadFileFlag()
	input := kit.ReadInputFile(filePath)
	result := desert.CountSteps(input)
	duration := time.Since(start).Milliseconds()
	log.Printf("Processed %s in %d ms with result %d.", filePath, duration, result)
}
