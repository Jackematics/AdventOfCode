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

func TotalWinnings(raw_input []string) int {
	hands := extractHands(raw_input)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].hand_type == hands[j].hand_type {
			return handWeaker(hands[i], hands[j])
		}

		return hands[i].hand_type < hands[j].hand_type
	})

	sum := 0
	for i := 0; i < len(hands); i++ {
		sum += (i + 1) * hands[i].bid
	}

	return sum
}

func extractHands(raw_input []string) []Hand {
	hands := []Hand{}
	for _, raw_hand := range raw_input {
		split := strings.Split(raw_hand, " ")
		bid, _ := strconv.Atoi(split[1])

		hands = append(hands, Hand{
			hand:      split[0],
			bid:       bid,
			hand_type: determineHandType(split[0]),
		})
	}

	return hands
}

func determineHandType(hand string) HandType {
	frequencies := calculateFrequencies(hand)

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

func calculateFrequencies(hand string) []int {
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

var cards = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func handWeaker(hand_a, hand_b Hand) bool {
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
