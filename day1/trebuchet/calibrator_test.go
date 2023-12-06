package trebuchet

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestSumCalibrationValuesData struct {
	Input               []string
	CheckStringPatterns bool
	Expected            int
}

var testSumCalibrationValuesData = []TestSumCalibrationValuesData{
	{kit.ReadInputFile("../input/sample1.txt"), false, 142},
	{kit.ReadInputFile("../input/sample2.txt"), true, 281},
	{kit.ReadInputFile("../input/sample3.txt"), true, 292},
	{kit.ReadInputFile("../input/data.txt"), false, 56049},
	{kit.ReadInputFile("../input/data.txt"), true, 54530},
}

func TestSumCalibrationValues(t *testing.T) {
	for i, data := range testSumCalibrationValuesData {
		result := SumCalibrationValues(data.Input, data.CheckStringPatterns)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkSumCalibrationValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCalibrationValues(testSumCalibrationValuesData[3].Input, false)
	}
}

func BenchmarkSumCalibrationValuesWithStringPatterns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCalibrationValues(testSumCalibrationValuesData[3].Input, true)
	}
}
