package internal

import (
	"flag"
)

func ReadStringFlag() bool {
	checkStringPatterns := flag.Bool("s", false, "Check string patterns for digit values.")
	flag.Parse()

	return *checkStringPatterns
}
