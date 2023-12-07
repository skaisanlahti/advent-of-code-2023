package camel

import (
	"slices"
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

var FaceCardValues = map[rune]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

const (
	FiveOfAKind  string = "FiveOfAKind"
	FourOfAKind  string = "FourOfAKind"
	FullHouse    string = "FullHouse"
	ThreeOfAKind string = "ThreeOfAKind"
	TwoPairs     string = "TwoPairs"
	OnePair      string = "OnePair"
	HighCard     string = "HighCard"
	Nothing      string = "Nothing"
)

var Evaluations = map[string]int{
	FiveOfAKind:  7,
	FourOfAKind:  6,
	FullHouse:    5,
	ThreeOfAKind: 4,
	TwoPairs:     3,
	OnePair:      2,
	HighCard:     1,
	Nothing:      0,
}

type Hand struct {
	Cards      string
	Bid        int
	Evaluation int
}

func CountWinnings(input []string, useJokers bool) int {
	evalute := evaluateHand
	if useJokers {
		FaceCardValues['J'] = 1
		evalute = evaluateHandWithJokers
	}

	hands := []Hand{}
	for _, line := range input {
		splits := strings.Split(line, " ")
		cards := splits[0]
		bid, _ := strconv.Atoi(splits[1])
		eval := evalute(cards)
		hands = append(hands, Hand{cards, bid, eval})
	}

	slices.SortFunc(hands, compareHands)

	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.Bid
	}
	return total
}

func countOccurances(cards string) map[rune]int {
	counts := map[rune]int{}
	for _, card := range cards {
		count, ok := counts[card]
		if !ok {
			counts[card] = 1
		} else {
			counts[card] = count + 1
		}
	}

	return counts
}

func evaluateHand(cards string) int {
	counts := countOccurances(cards)
	kinds, pairs := 0, 0
	for _, count := range counts {
		if kinds < count {
			kinds = count
		}

		if count == 2 {
			pairs++
		}
	}

	switch {
	case kinds == 5:
		return Evaluations[FiveOfAKind]
	case kinds == 4:
		return Evaluations[FourOfAKind]
	case kinds == 3 && pairs == 1:
		return Evaluations[FullHouse]
	case kinds == 3:
		return Evaluations[ThreeOfAKind]
	case pairs == 2:
		return Evaluations[TwoPairs]
	case pairs == 1:
		return Evaluations[OnePair]
	default:
		return Evaluations[HighCard]
	}
}

func evaluateHandWithJokers(cards string) int {
	counts := countOccurances(cards)
	kinds, pairs, jokers := 0, 0, 0
	for key, count := range counts {
		if key == 'J' {
			jokers = count
			continue
		}

		if kinds < count {
			kinds = count
		}

		if count == 2 {
			pairs++
		}
	}

	switch {
	case kinds+jokers == 5:
		return Evaluations[FiveOfAKind]
	case kinds+jokers == 4:
		return Evaluations[FourOfAKind]
	case kinds == 3 && pairs == 1, kinds == 3 && jokers == 1, pairs == 2 && jokers == 1, pairs == 1 && jokers >= 2:
		return Evaluations[FullHouse]
	case kinds+jokers == 3, kinds == 2 && jokers >= 1:
		return Evaluations[ThreeOfAKind]
	case pairs == 2, pairs == 1 && jokers == 1:
		return Evaluations[TwoPairs]
	case pairs == 1, pairs == 0 && jokers == 1:
		return Evaluations[OnePair]
	default:
		return Evaluations[HighCard]
	}
}

func evaluateCard(card rune) int {
	value, ok := kit.RuneToInt(card)
	if !ok {
		value = FaceCardValues[card]
	}

	return value
}

func compareHands(a, b Hand) int {
	if a.Evaluation != b.Evaluation {
		return a.Evaluation - b.Evaluation
	}

	acards := []rune(a.Cards)
	bcards := []rune(b.Cards)
	for i := 0; i < len(a.Cards); i++ {
		avalue := evaluateCard(acards[i])
		bvalue := evaluateCard(bcards[i])
		if avalue != bvalue {
			return avalue - bvalue
		}
	}

	return 0
}
