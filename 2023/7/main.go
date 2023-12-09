package main

import (
	"advent-of-code/2023/1/helpers"
	"strconv"
	"strings"
)

const inputLocation = "sample.txt"

type strength int
type cardValue int

const (
	FiveOfAKind  strength = 1 // Five of a kind, where all five cards have the same label: AAAAA
	FourOfAKind  strength = 2 // Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	FullHouse    strength = 3 // Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	ThreeOfAKind strength = 4 // Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	TwoPair      strength = 5 // Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	OnePair      strength = 6 // One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	HighCard     strength = 7 // High card, where all cards' labels are distinct: 23456

	ace   cardValue = 14
	king  cardValue = 13
	queen cardValue = 12
	jack  cardValue = 11
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
	rank     int
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

	//game.rankHands()
	//fmt.Printf("The winnings for Day 1 are: %v", calculateWinnings(game))
}

func createHand(line string) hand {

	var hand hand
	split := strings.Split(line, " ")
	bid, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	hand.bid = bid

	for i, value := range split[0] {
		var card card
		card.value = determineValue(string(value - 0))
		card.positionInHand = i
		hand.cards = append(hand.cards, card)
	}

	hand.calculateStrength()
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

	//TODO work out strength

	strength := FiveOfAKind
	for _, card := range hand.cards {
		if card.appearancesInHand == 5 {
			strength = FiveOfAKind
			break
		}
		if card.appearancesInHand == 4 {
			strength = FourOfAKind
			break
		}
		if card.appearancesInHand == 3 {
			strength = ThreeOfAKind
			break
		}
		if card.appearancesInHand == 2 {
			for _, v := range v {

			}
		}
	}

	hand.strength = strength

}

// func (game *game) rankHands() {
// 	rank := 0
// 	for _, hand := range game.hands {

// 	}
// }

// func calculateWinnings(game game) int {
// 	winnings := 0
// 	for _, hand := range game.hands {
// 		winnings = winnings + (hand.rank * hand.bid)
// 	}
// 	return winnings
// }
