package camel

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type TestCountWinningsData struct {
	Input     []string
	UseJokers bool
	Expected  int
}

var testCountWinningsData = []TestCountWinningsData{
	{kit.ReadInputFile("../input/sample.txt"), false, 6440},
	{kit.ReadInputFile("../input/data.txt"), false, 250957639},
	{kit.ReadInputFile("../input/sample.txt"), true, 5905},
	{kit.ReadInputFile("../input/data.txt"), true, 251515496},
}

func TestCountWinnings(t *testing.T) {
	for i, data := range testCountWinningsData {
		result := CountWinnings(data.Input, data.UseJokers)
		if result != data.Expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, data.Expected, result)
		}
	}
}

func BenchmarkCountWinnings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWinnings(testCountWinningsData[1].Input, false)
	}
}

func BenchmarkCountWinningsWithJokers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWinnings(testCountWinningsData[1].Input, true)
	}
}
