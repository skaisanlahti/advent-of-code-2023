package camel

import (
	"flag"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func ReadFlags() (bool, string) {
	useJokers := flag.Bool("j", false, "Input file path.")
	filePath := kit.ReadFileFlag()
	return *useJokers, filePath
}
