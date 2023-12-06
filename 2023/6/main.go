package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type attempt struct {
	heldTime int
	time     int
	distance int
}

type race struct {
	mainAttempt       attempt
	potentialAttempts []attempt
}

const inputLocation = "input.txt"

func main() {

	var races []race

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var times map[int]int
	var distances map[int]int

	for scanner.Scan() {
		line := scanner.Text()

		//part 2 only - this is ugly as I can't generate the answer for day 1 and day 2
		// at the same time but no time for rewrites!
		line = strings.ReplaceAll(line, " ", "")

		if strings.Contains(line, "Time:") {
			times = parseNumbers(line)
		}
		if strings.Contains(line, "Distance:") {
			distances = parseNumbers(line)
		}
	}

	for i := 1; i <= len(times); i++ {
		var race race
		race.mainAttempt.time = times[i]
		race.mainAttempt.distance = distances[i]
		races = append(races, race)
	}

	for i := range races {
		races[i].calculatePotentialAttempts()
	}

	fmt.Printf("The answer to day 6 is: %v", calculateAnswer(races))

}

func parseNumbers(line string) map[int]int {

	var numbers = make(map[int]int)

	numberFound := false

	var numberString string
	for i, v := range line {

		// if a digit is found and it’s the first digit found, i.e a new number
		// set the start position and add the number to numberString
		if unicode.IsDigit(v) && !numberFound {
			numberFound = true
			numberString = numberString + string(v-0)
			continue
		}
		// if a number is found and it’s not the first number, i.e continuing multi digit number
		// add the number to parsed Numbers
		if unicode.IsDigit(v) && numberFound {

			numberString = numberString + string(v-0)

			// if it's the last number on a line then it's the end of a number
			if i+1 == len(line) {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					// ... handle error
					panic(err)
				}
				numbers[len(numbers)+1] = number
				numberFound = false
				numberString = ""
			}

			continue
		}
		// if a digit is not found and numberFound is set to true, this number must have ended at the previous loop
		// set number found to false and append the completed number
		if !unicode.IsDigit(v) && numberFound {

			number, err := strconv.Atoi(numberString)
			if err != nil {
				// ... handle error
				panic(err)
			}
			numbers[len(numbers)+1] = number
			numberFound = false
			numberString = ""
		}
	}
	return numbers

}

func (race *race) calculatePotentialAttempts() {

	var potentialAttempts []attempt

	for i := 0; i <= race.mainAttempt.time; i++ {
		var attempt attempt
		attempt.time = race.mainAttempt.time
		attempt.heldTime = i
		attempt.distance = i * (attempt.time - attempt.heldTime)
		potentialAttempts = append(potentialAttempts, attempt)
	}
	race.potentialAttempts = potentialAttempts

}

func calculateAnswer(races []race) int {

	var winningAttempts []int
	count := 0
	result := 0

	for _, race := range races {
		for _, attempt := range race.potentialAttempts {
			if attempt.distance > race.mainAttempt.distance {
				count++
			}
		}
		winningAttempts = append(winningAttempts, count)
		count = 0
	}

	result = winningAttempts[0]

	for i := 1; i < len(winningAttempts); i++ {
		result = result * winningAttempts[i]
	}

	return result
}
