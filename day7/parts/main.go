package main

import (
	"log"
	"time"

	"github.com/skaisanlahti/advent-of-code-2023/day7/camel"
	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func main() {
	start := time.Now()
	useJokers, filePath := camel.ReadFlags()
	input := kit.ReadInputFile(filePath)
	result := camel.CountWinnings(input, useJokers)
	duration := time.Since(start).Milliseconds()
	log.Printf("Processed %s in %d ms with result %d.", filePath, duration, result)
}
