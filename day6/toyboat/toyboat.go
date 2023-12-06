package toyboat

import (
	"math"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type Race struct {
	Time   int
	Record int
}

func RaceSetMarginProduct(input []string) int {
	times := kit.NumbersFromString(input[0])
	records := kit.NumbersFromString(input[1])
	races := []Race{}
	for i := range times {
		races = append(races, Race{times[i], records[i]})
	}

	return marginsProduct(races)
}

func RaceMargin(input []string) int {
	timeString := strings.Replace(input[0], " ", "", -1)
	time := kit.NumbersFromString(timeString)[0]
	recordString := strings.Replace(input[1], " ", "", -1)
	record := kit.NumbersFromString(recordString)[0]
	race := Race{time, record}
	margin := calculateMargin(race)
	return margin
}

func marginsProduct(races []Race) int {
	margins := []int{}
	for _, race := range races {
		margin := calculateMargin(race)
		margins = append(margins, margin)
	}

	product := 1
	for _, value := range margins {
		product *= value
	}

	return product
}

func bruteForceMargin(race Race) int {
	margin := 0
	for hold := 0; hold < race.Time; hold++ {
		distance := (race.Time - hold) * hold
		if distance > race.Record {
			margin++
		}
	}

	return margin
}

func findMarginBounds(race Race) int {
	start := 0
	for hold := 0; hold < race.Time; hold++ {
		distance := (race.Time - hold) * hold
		if distance > race.Record {
			start = hold
			break
		}
	}

	end := 0
	for hold := race.Time; hold > 0; hold-- {
		distance := (race.Time - hold) * hold
		if distance > race.Record {
			end = hold
			break
		}
	}

	return 1 + end - start
}

func calculateMargin(race Race) int {
	root1, root2, err := kit.QuadraticRoots(-1, float64(race.Time), -float64(race.Record))
	if err != nil {
		return 0
	}

	if root1 == root2 {
		return 1
	}

	start := 0
	end := 0
	if root1 < root2 {
		start = int(math.Floor(root1 + 1))
		end = int(math.Ceil(root2 - 1))
	} else {
		start = int(math.Floor(root2 + 1))
		end = int(math.Ceil(root1 - 1))
	}

	margin := 1 + end - start
	return margin
}

func ternarySearchMargin(race Race) int {
	distance := func(hold float64) float64 {
		return (hold * (float64(race.Time) - hold))
	}

	middle := int(math.RoundToEven(kit.TernarySearch(distance, 0, float64(race.Time), 1e2)))
	s, e := middle, middle
	for int(distance(float64(s))) > race.Record {
		s--
	}
	for int(distance(float64(e))) > race.Record {
		e++
	}

	start := s + 1
	end := e - 1
	margin := 1 + end - start
	return margin
}
