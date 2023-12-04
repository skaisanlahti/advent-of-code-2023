package internal

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type testCase struct {
	input    []string
	expected int
}

var testCases = []testCase{
	{kit.ReadInputFile("../input/sample.txt"), 13},
	{kit.ReadInputFile("../input/data.txt"), 15268},
	{kit.ReadInputFile("../input/sample.txt"), 30},
	{kit.ReadInputFile("../input/data.txt"), 6283755},
}

func TestSumCardPoints(t *testing.T) {
	for i, testCase := range testCases[0:1] {
		result := SumCardPoints(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumCardPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCardPoints(testCases[1].input)
	}
}

func TestCountPointsSinglePass(t *testing.T) {
	for i, testCase := range testCases[0:1] {
		result := CountPointsSinglePass(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkCountPointsSinglePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPointsSinglePass(testCases[1].input)
	}
}

func TestSumCardCount(t *testing.T) {
	for i, testCase := range testCases[2:3] {
		result := SumCardCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumCardCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumCardCount(testCases[3].input)
	}
}

func TestSumCardCountSinglePass(t *testing.T) {
	for i, testCase := range testCases[2:3] {
		result := CountCardsSinglePass(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumCardCountSinglePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountCardsSinglePass(testCases[3].input)
	}
}
