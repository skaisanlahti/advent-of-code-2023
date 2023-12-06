package garden

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestFindLowestLocationValueData struct {
	Input    []string
	Expected int
}

var testFindLowestLocationValueData = []TestFindLowestLocationValueData{
	{kit.ReadInputFile("../input/sample.txt"), 35},
	{kit.ReadInputFile("../input/data.txt"), 836040384},
}

func TestFindLowestLocationValue(t *testing.T) {
	for i, data := range testFindLowestLocationValueData {
		result := FindLowestLocationValue(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkFindLowestLocationValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindLowestLocationValue(testFindLowestLocationValueData[1].Input)
	}
}

type TestFindLowestLocationRangeData struct {
	Input    []string
	Expected int
}

var testFindLowestLocationRangeData = []TestFindLowestLocationRangeData{
	{kit.ReadInputFile("../input/sample.txt"), 46},
	{kit.ReadInputFile("../input/data.txt"), 10834440},
}

func TestFindLowestLocationRange(t *testing.T) {
	for i, data := range testFindLowestLocationRangeData {
		result := FindLowestLocationRange(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkFindLowestLocationRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindLowestLocationRange(testFindLowestLocationValueData[1].Input)
	}
}
