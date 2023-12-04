package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputLocation = "input.txt"

type card struct {
	id              int
	winningNumbers  []int
	chosenNumbers   []int
	numberOfMatches int
	isWinner        bool
	numberOfCopies  int
}

var cards []card

func main() {

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		cards = append(cards, parseCard(line))
		for i, card := range cards {

			cards[i].isWinner, cards[i].numberOfMatches = cardIsWinner(card.winningNumbers, card.chosenNumbers)
			fmt.Printf("Card %v chosen numbers are %v and the winning numbers are %v. Number of matches: %v\n", card.id, card.chosenNumbers, card.winningNumbers, card.numberOfMatches)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("The answer to part 1 is: %v\n", calculatePart1Answer(cards))
	fmt.Printf("The answer to part 2 is: %v\n", calculatePart2Answer(cards))
}

func parseCard(line string) card {

	card := card{}

	line = strings.ReplaceAll(line, "Card ", "")

	splitLine := strings.Split(line, ":")
	splitLine[0] = strings.ReplaceAll(splitLine[0], " ", "")
	id, err := strconv.Atoi(splitLine[0])

	if err != nil {
		panic(err)
	}

	card.id = id

	numbers := strings.Split(splitLine[1], "|")
	card.winningNumbers = extractNumbers(numbers[0])
	card.chosenNumbers = extractNumbers(numbers[1])

	return card
}

func extractNumbers(numbers string) []int {

	var parsedNumbers []int
	numbersSplit := strings.Split(numbers, " ")
	for _, v := range numbersSplit {
		if v != "" {
			v = strings.ReplaceAll(v, " ", "")
			parsedNumber, _ := strconv.Atoi(v)
			parsedNumbers = append(parsedNumbers, parsedNumber)
		}
	}

	return parsedNumbers

}

func cardIsWinner(winningNumbers []int, chosenNumbers []int) (isWinner bool, numberOfMatches int) {

	for _, chosenNumber := range chosenNumbers {
		for _, winningNumber := range winningNumbers {
			if chosenNumber == winningNumber {
				numberOfMatches++
			}
		}
	}

	if numberOfMatches == 0 {
		return false, 0
	}
	return true, numberOfMatches
}

func calculatePart1Answer(cards []card) (result int) {
	for _, card := range cards {
		sum := 0
		if card.isWinner {
			sum = 1
			for i := 1; i < card.numberOfMatches; i++ {
				sum = sum * 2
			}
		}
		result = result + sum
	}

	return result
}

func calculatePart2Answer(cards []card) (result int) {

	//for each card
	// if winner
	// for each matching number
	// cardid + i -> increment number of cards
	//

	for _, card := range cards {
		if card.isWinner {
			for i := 1; i < card.numberOfMatches; i++ {
				// card id 1 has 5 matches
				// card id 2,3,4,5,6 get + 1 copies
			}
		}

	}
	return 0
}
