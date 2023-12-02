package internal

import "testing"

type powerCase struct {
	filePath string
	expected int
}

var powerCases = []powerCase{
	{"../input/sample.txt", 2286},
	{"../input/data.txt", 83707},
}

func TestSumGamePowers(t *testing.T) {
	for i, testCase := range powerCases {
		input := ReadInputFile(testCase.filePath)
		result := SumGamePowers(input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}
