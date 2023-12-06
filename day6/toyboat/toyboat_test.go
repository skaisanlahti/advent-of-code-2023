package toyboat

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestRaceSetMarginProductData struct {
	Input    []string
	Expected int
}

var testRaceSetMarginProductData = []TestRaceSetMarginProductData{
	{kit.ReadInputFile("../input/sample.txt"), 288},
	{kit.ReadInputFile("../input/data.txt"), 4403592},
}

func TestRaceSetMarginProduct(t *testing.T) {
	for i, data := range testRaceSetMarginProductData {
		result := RaceSetMarginProduct(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkRaceSetMarginProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RaceSetMarginProduct(testRaceSetMarginProductData[1].Input)
	}
}

type TestRaceMarginData struct {
	Input    []string
	Expected int
}

var testRaceMarginData = []TestRaceMarginData{
	{kit.ReadInputFile("../input/sample.txt"), 71503},
	{kit.ReadInputFile("../input/data.txt"), 38017587},
}

func TestRaceMargin(t *testing.T) {
	for i, data := range testRaceMarginData {
		result := RaceMargin(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkRaceMargin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RaceMargin(testRaceMarginData[1].Input)
	}
}
