package internal

import "testing"

type validateCase struct {
	filePath string
	red      int
	green    int
	blue     int
	expected int
}

var validateCases = []validateCase{
	{"../input/sample.txt", 12, 13, 14, 8},
	{"../input/data.txt", 12, 13, 14, 2685},
}

func TestSumValidGameIds(t *testing.T) {
	for i, testCase := range validateCases {
		input := ReadInputFile(testCase.filePath)
		result := SumValidGameIds(input, testCase.red, testCase.green, testCase.blue)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}
