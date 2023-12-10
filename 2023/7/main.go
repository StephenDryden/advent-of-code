package main

import (
	"advent-of-code/2023/1/helpers"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const inputLocation = "input.txt"

type strength int
type cardValue int

const (
	FiveOfAKind  strength = 7 // Five of a kind, where all five cards have the same label: AAAAA
	FourOfAKind  strength = 6 // Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	FullHouse    strength = 5 // Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	ThreeOfAKind strength = 4 // Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	TwoPair      strength = 3 // Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	OnePair      strength = 2 // One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	HighCard     strength = 1 // High card, where all cards' labels are distinct: 23456

	ace   cardValue = 14
	king  cardValue = 13
	queen cardValue = 12
	//jack  cardValue = 11
	jack  cardValue = 1
	ten   cardValue = 10
	nine  cardValue = 9
	eight cardValue = 8
	seven cardValue = 7
	six   cardValue = 6
	five  cardValue = 5
	four  cardValue = 4
	three cardValue = 3
	two   cardValue = 2
)

type card struct {
	positionInHand    int
	value             cardValue
	appearancesInHand int
}

type hand struct {
	raw      string
	bid      int
	cards    []card
	strength strength
}

type game struct {
	hands []hand
}

func main() {

	var game game
	lines := helpers.ReadFile(inputLocation)

	for _, line := range lines {
		game.hands = append(game.hands, createHand(line))
	}

	game.rankHands()
	fmt.Printf("The winnings for Day 7 part 1 are: %v", calculateWinnings(game))
}

func createHand(line string) hand {

	var hand hand
	split := strings.Split(line, " ")
	bid, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	hand.bid = bid
	hand.raw = split[0]

	for i, value := range split[0] {
		var card card
		card.value = determineValue(string(value - 0))
		card.positionInHand = i
		hand.cards = append(hand.cards, card)
	}

	//Uncomment and Part 1
	//hand.calculateStrength()
	hand.calculateStrengthWithJokers()
	return hand
}

func determineValue(input string) cardValue {
	switch input {
	case "A":
		return ace
	case "K":
		return king
	case "Q":
		return queen
	case "J":
		return jack
	case "T":
		return ten
	case "9":
		return nine
	case "8":
		return eight
	case "7":
		return seven
	case "6":
		return six
	case "5":
		return five
	case "4":
		return four
	case "3":
		return three
	case "2":
		return two
	}

	return 1
}

func (hand *hand) calculateStrength() {

	for i, currentCard := range hand.cards {
		count := 0
		for _, card := range hand.cards {
			if currentCard.value == card.value {
				count++
			}
		}
		hand.cards[i].appearancesInHand = count
		count = 0
	}

	appearsFiveTimes := 0
	appearsFourTimes := 0
	appearsThreeTimes := 0
	appearsTwoTimes := 0
	appearsOneTime := 0
	for _, card := range hand.cards {

		switch card.appearancesInHand {
		case 5:
			appearsFiveTimes++
		case 4:
			appearsFourTimes++
		case 3:
			appearsThreeTimes++
		case 2:
			appearsTwoTimes++
		case 1:
			appearsOneTime++

		}

	}

	if appearsFiveTimes == 5 {
		hand.strength = FiveOfAKind
	}
	if appearsFourTimes == 4 {
		hand.strength = FourOfAKind
	}
	if appearsThreeTimes == 3 && appearsTwoTimes == 2 {
		hand.strength = FullHouse
	}
	if appearsThreeTimes == 3 && appearsTwoTimes != 2 {
		hand.strength = ThreeOfAKind
	}
	if appearsTwoTimes == 4 {
		hand.strength = TwoPair
	}
	if appearsTwoTimes == 2 && appearsThreeTimes != 3 {
		hand.strength = OnePair
	}
	if appearsOneTime == 5 {
		hand.strength = HighCard
	}

}

func (hand *hand) calculateStrengthWithJokers() {

	hand.calculateStrength()
	numberOfJokers := countJokers(hand)
	if numberOfJokers > 0 {

		switch hand.strength {
		case FiveOfAKind:
			hand.strength = FiveOfAKind
		case FourOfAKind:
			hand.strength = FiveOfAKind
		case FullHouse:
			hand.strength = FiveOfAKind
		case ThreeOfAKind:
			hand.strength = FourOfAKind
		case TwoPair:
			if numberOfJokers == 2 {
				hand.strength = FourOfAKind
			}
			if numberOfJokers == 1 {
				hand.strength = ThreeOfAKind
			}

		case OnePair:
			hand.strength = ThreeOfAKind
		case HighCard:
			hand.strength = OnePair
		}

	}

}

func countJokers(hand *hand) (count int) {

	for _, card := range hand.cards {
		if card.value == jack {
			return card.appearancesInHand
		}
	}
	return 0
}

func (game *game) rankHands() {

	sort.Slice(game.hands, func(i, j int) bool {

		if game.hands[i].strength > game.hands[j].strength {
			return true
		}

		if game.hands[i].strength == game.hands[j].strength {
			for y := 0; y < 5; y++ {
				if game.hands[i].cards[y].value > game.hands[j].cards[y].value {
					return true
				}
				if game.hands[i].cards[y].value < game.hands[j].cards[y].value {
					return false
				}
				// if game.hands[i].strength == FiveOfAKind {
				// 	return game.hands[i].cards[0].value > game.hands[j].cards[0].value
				// }
			}
		}

		return false
	})

}

func calculateWinnings(game game) int {
	winnings := 0

	rank := len(game.hands)
	for i := range game.hands {

		fmt.Printf("Rank: %v Strength: %v Hand: %v\n", rank, game.hands[i].strength, game.hands[i].raw)

		winnings = (rank * game.hands[i].bid) + winnings
		//fmt.Printf("%v + (%v * %v)\n", winnings, rank, game.hands[i].bid)
		rank--
	}

	return winnings
}
