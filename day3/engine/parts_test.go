package engine

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestSumEnginePartNumbersData struct {
	Input    []string
	Expected int
}

var testSumEnginePartNumbersData = []TestSumEnginePartNumbersData{
	{kit.ReadInputFile("../input/sample.txt"), 4361},
	{kit.ReadInputFile("../input/data.txt"), 525119},
}

func TestSumEnginePartNumbers(t *testing.T) {
	for i, data := range testSumEnginePartNumbersData {
		result := SumPartNumbers(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkSumEnginePartNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumPartNumbers(testSumEnginePartNumbersData[1].Input)
	}
}
