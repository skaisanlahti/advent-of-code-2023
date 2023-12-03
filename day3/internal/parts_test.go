package internal

import "testing"

type partCase struct {
	input    []string
	expected int
}

var partCases = []partCase{
	{ReadInputFile("../input/sample.txt"), 4361},
	{ReadInputFile("../input/data.txt"), 525119},
}

func TestSumEnginePartNumbers(t *testing.T) {
	for i, testCase := range partCases {
		result := SumPartNumbers(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumEnginePartNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumPartNumbers(partCases[1].input)
	}
}