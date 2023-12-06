package trebuchet

import (
	"flag"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func ReadFlags() (bool, string) {
	checkStringPatterns := flag.Bool("s", false, "Check string patterns for digit values.")
	filePath := kit.ReadFileFlag()
	flag.Parse()

	return *checkStringPatterns, filePath
}
