package internal

import "testing"

type testCase struct {
	filePath            string
	checkStringPatterns bool
	expected            int
}

var testCases = []testCase{
	{"../input/sample1.txt", false, 142},
	{"../input/sample2.txt", true, 281},
	{"../input/sample3.txt", true, 292},
	{"../input/data.txt", false, 56049},
	{"../input/data.txt", true, 54530},
}

func TestCalibrate(t *testing.T) {
	for i, testCase := range testCases {
		input := ReadInputFile(testCase.filePath)
		result := SumCalibrationValues(input, testCase.checkStringPatterns)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}
