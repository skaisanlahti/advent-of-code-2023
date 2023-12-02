package internal

import "testing"

type validateCase struct {
	input    []string
	red      int
	green    int
	blue     int
	expected int
}

var validateCases = []validateCase{
	{ReadInputFile("../input/sample.txt"), 12, 13, 14, 8},
	{ReadInputFile("../input/data.txt"), 12, 13, 14, 2685},
}

func TestSumValidGameIds(t *testing.T) {
	for i, testCase := range validateCases {
		result := SumValidGameIds(testCase.input, testCase.red, testCase.green, testCase.blue)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumValidGameIds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumValidGameIds(validateCases[1].input, validateCases[1].red, validateCases[1].green, validateCases[1].blue)
	}
}
