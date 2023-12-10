package camel_cards

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	High HandType = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

type Hand struct {
	hand      string
	hand_type HandType
	bid       int
}

func TotalWinnings(raw_input []string, part string) int {
	hands := extractHands(raw_input, part)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].hand_type == hands[j].hand_type {
			return handWeaker(hands[i], hands[j], part)
		}

		return hands[i].hand_type < hands[j].hand_type
	})

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += (i + 1) * hands[i].bid
	}

	return sum
}

func extractHands(raw_input []string, part string) []Hand {
	hands := []Hand{}
	for _, raw_hand := range raw_input {
		split := strings.Split(raw_hand, " ")
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, Hand{
			hand:      split[0],
			bid:       bid,
			hand_type: determineHandType(split[0], part),
		})
	}

	return hands
}

func determineHandType(hand string, part string) HandType {
	var frequencies []int

	if part == "part_one" {
		frequencies = calculateFrequenciesPartOne(hand)
	} else {
		frequencies = calculateFrequenciesPartTwo(hand)
	}

	if handHasFrequency(frequencies, 5) {
		return FiveKind
	}

	if handHasFrequency(frequencies, 4) {
		return FourKind
	}

	if handHasFrequency(frequencies, 3) && handHasFrequency(frequencies, 2) {
		return FullHouse
	}

	if handHasFrequency(frequencies, 3) {
		return ThreeKind
	}

	if handIsTwoPair(frequencies) {
		return TwoPair
	}

	if handHasFrequency(frequencies, 2) {
		return OnePair
	}

	return High
}

func handHasFrequency(frequencies []int, frequency int) bool {
	return slices.Index(frequencies, frequency) != -1
}

func handIsTwoPair(frequencies []int) bool {
	two_sum := 0
	for _, frequency := range frequencies {
		if frequency == 2 {
			two_sum++
		}
	}

	return two_sum == 2
}

func calculateFrequenciesPartOne(hand string) []int {
	frequencies := []int{}

	forbidden := []rune{}
	for _, current_card := range hand {
		if slices.Contains(forbidden, current_card) {
			continue
		}

		sum := 0
		for _, comparison_card := range hand {
			if current_card == comparison_card {
				sum++
			}
		}

		frequencies = append(frequencies, sum)
		forbidden = append(forbidden, current_card)
	}

	return frequencies
}

var cardsPartOne = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var cardsPartTwo = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func handWeaker(hand_a, hand_b Hand, part string) bool {
	var cards []byte
	if part == "part_one" {
		cards = cardsPartOne
	} else {
		cards = cardsPartTwo
	}

	for i := 0; i < len(hand_a.hand); i++ {
		a_value := slices.Index(cards, hand_a.hand[i])
		b_value := slices.Index(cards, hand_b.hand[i])

		if a_value == b_value {
			continue
		}

		return a_value <= b_value
	}

	return true
}

func calculateFrequenciesPartTwo(hand string) []int {
	frequencies := []int{}

	forbidden := []rune{'J'}
	for _, current_card := range hand {
		if slices.Contains(forbidden, current_card) {
			continue
		}

		sum := 0
		for _, comparison_card := range hand {
			if current_card == comparison_card {
				sum++
			}
		}

		frequencies = append(frequencies, sum)
		forbidden = append(forbidden, current_card)
	}

	println("hand: ", hand)
	jokerCount := getJokerCount(hand)

	if len(frequencies) == 0 {
		return []int{5}
	}

	max := slices.Max(frequencies)
	max_index := slices.Index(frequencies, max)
	frequencies[max_index] += jokerCount

	return frequencies
}

func getJokerCount(hand string) int {
	joker_count := 0
	for _, card := range hand {
		if card == 'J' {
			joker_count++
		}
	}

	return joker_count
}
