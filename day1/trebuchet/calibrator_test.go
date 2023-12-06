package trebuchet

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type testCase struct {
	input               []string
	checkStringPatterns bool
	expected            int
}

var testCases = []testCase{
	{kit.ReadInputFile("../input/sample1.txt"), false, 142},
	{kit.ReadInputFile("../input/sample2.txt"), true, 281},
	{kit.ReadInputFile("../input/sample3.txt"), true, 292},
	{kit.ReadInputFile("../input/data.txt"), false, 56049},
	{kit.ReadInputFile("../input/data.txt"), true, 54530},
}

func TestSumCalibrationValues(t *testing.T) {
	for i, testCase := range testCases {
		result := SumCalibrationValues(testCase.input, testCase.checkStringPatterns)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumCalibrationValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCalibrationValues(testCases[3].input, false)
	}
}

func BenchmarkSumCalibrationValuesWithStringPatterns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCalibrationValues(testCases[3].input, true)
	}
}