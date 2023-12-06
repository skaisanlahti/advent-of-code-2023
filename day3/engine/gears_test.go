package engine

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestSumGearRatiosData struct {
	Input    []string
	Expected int
}

var testSumGearRatiosData = []TestSumGearRatiosData{
	{kit.ReadInputFile("../input/sample.txt"), 467835},
	{kit.ReadInputFile("../input/data.txt"), 76504829},
}

func TestSumGearRatios(t *testing.T) {
	for i, data := range testSumGearRatiosData {
		result := SumGearRatios(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkSumGearRatios(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumGearRatios(testSumGearRatiosData[1].Input)
	}
}
