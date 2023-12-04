package internal

import (
	"flag"
)

func ReadColorFlags() (int, int, int) {
	r := flag.Int("r", 0, "Number of red cubes.")
	g := flag.Int("g", 0, "Number of green cubes.")
	b := flag.Int("b", 0, "Number of blue cubes.")
	flag.Parse()

	return *r, *g, *b
}
