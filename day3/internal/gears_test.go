package internal

import "testing"

type gearCase struct {
	input    []string
	expected int
}

var gearCases = []gearCase{
	{ReadInputFile("../input/sample.txt"), 467835},
	{ReadInputFile("../input/data.txt"), 76504829},
}

func TestSumGearRatios(t *testing.T) {
	for i, testCase := range gearCases {
		result := SumGearRatios(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumGearRatios(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumGearRatios(gearCases[1].input)
	}
}
