package internal

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type Race struct {
	time   int
	record int
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
	return margin(race)
}

func margin(race Race) int {
	start := 0
	for hold := 0; hold < race.time; hold++ {
		distance := (race.time - hold) * hold
		if distance > race.record {
			start = hold
			break
		}
	}

	end := 0
	for hold := race.time; hold > 0; hold-- {
		distance := (race.time - hold) * hold
		if distance > race.record {
			end = hold
			break
		}
	}

	return 1 + end - start
}

func marginsProduct(races []Race) int {
	margins := []int{}
	for _, race := range races {
		margin := margin(race)
		margins = append(margins, margin)
	}

	product := 1
	for _, value := range margins {
		product *= value
	}

	return product
}
