package cube

import (
	"flag"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func ReadFlags() (int, int, int, string) {
	r := flag.Int("r", 0, "Number of red cubes.")
	g := flag.Int("g", 0, "Number of green cubes.")
	b := flag.Int("b", 0, "Number of blue cubes.")
	filePath := kit.ReadFileFlag()

	return *r, *g, *b, filePath
}
