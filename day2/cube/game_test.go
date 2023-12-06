package cube

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestSumValidGameIdsData struct {
	Input    []string
	Red      int
	Green    int
	Blue     int
	Expected int
}

var testSumValidGameIdsData = []TestSumValidGameIdsData{
	{kit.ReadInputFile("../input/sample.txt"), 12, 13, 14, 8},
	{kit.ReadInputFile("../input/data.txt"), 12, 13, 14, 2685},
}

func TestSumValidGameIds(t *testing.T) {
	for i, data := range testSumValidGameIdsData {
		result := SumValidGameIds(data.Input, data.Red, data.Green, data.Blue)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkSumValidGameIds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumValidGameIds(testSumValidGameIdsData[1].Input, testSumValidGameIdsData[1].Red, testSumValidGameIdsData[1].Green, testSumValidGameIdsData[1].Blue)
	}
}

type TestSumGamePowersData struct {
	Input    []string
	Expected int
}

var testSumGamePowersData = []TestSumGamePowersData{
	{kit.ReadInputFile("../input/sample.txt"), 2286},
	{kit.ReadInputFile("../input/data.txt"), 83707},
}

func TestSumGamePowers(t *testing.T) {
	for i, data := range testSumGamePowersData {
		result := SumGamePowers(data.Input)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkSumGamePowers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumGamePowers(testSumGamePowersData[1].Input)
	}
}
