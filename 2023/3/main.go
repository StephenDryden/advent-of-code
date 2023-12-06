package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const inputLocation = "input.txt"

type number struct {
	line              int
	startPosition     int
	value             int
	lengthOfNumber    int
	perimeter         []rune
	isPartOfSchematic bool
}

func main() {
	var input = make(map[int]string)

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		input[lineCount] = line

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		lineCount++
	}

	var numbers = parseNumbers(input)
	for i, number := range numbers {
		numbers[i].lengthOfNumber = lengthOfNumber(number.value)
		numbers[i].determinePerimeter(input)
		numbers[i].isPartOfSchematic = perimeterContainsSymbol(numbers[i].perimeter)
	}

	fmt.Printf("The answer to Day3 Part 1 is %v", calculateDay1Answer(numbers))

}

func parseNumbers(lines map[int]string) []number {

	var numbers []number
	for i := 0; i <= len(lines); i++ {
		var number number
		numberFound := false
		number.line = i

		line := lines[i]
		var numberString string
		for i, v := range line {

			// if a digit is found and it’s the first digit found, i.e a new number
			// set the start position and add the number to numberString
			if unicode.IsDigit(v) && !numberFound {
				numberFound = true
				number.startPosition = i
				numberString = numberString + string(v-0)
				continue
			}
			// if a number is found and it’s not the first number, i.e continuing multi digit number
			// add the number to parsed Numbers
			if unicode.IsDigit(v) && numberFound {

				numberString = numberString + string(v-0)

				if i+1 == len(line) {
					i, err := strconv.Atoi(numberString)
					if err != nil {
						// ... handle error
						panic(err)
					}
					number.value = i
					numbers = append(numbers, number)
					numberFound = false
					numberString = ""
				}

				continue
			}
			// if a digit is not found and numberFound is set to true, this number must have ended at the previous loop
			// set number found to false and append the completed number
			if !unicode.IsDigit(v) && numberFound {

				i, err := strconv.Atoi(numberString)
				if err != nil {
					// ... handle error
					panic(err)
				}
				number.value = i
				numbers = append(numbers, number)
				numberFound = false
				numberString = ""
			}
		}
	}
	return numbers
}

func lengthOfNumber(number int) int {
	count := 0
	for number > 0 {
		number = number / 10
		count++
	}
	return count
}

func perimeterContainsSymbol(perimeter []rune) bool {
	for _, value := range perimeter {
		if !unicode.IsDigit(value) && value != 46 {
			return true
		}
	}

	return false
}

func calculateDay1Answer(numbers []number) int {
	result := 0
	for _, number := range numbers {
		if number.isPartOfSchematic {
			result = result + number.value
		}
	}

	return result

}

func (number *number) determinePerimeter(input map[int]string) {

	var perimeter []rune

	previousLine := ""
	currentLine := input[number.line]
	nextLine := ""

	if number.line != 0 {
		previousLine = input[number.line-1]
	}
	if number.line != len(input) {
		nextLine = input[number.line+1]
	}

	for i, v := range previousLine {
		if !unicode.IsDigit(v) && v != 46 && i >= number.startPosition-1 && i <= number.startPosition+number.lengthOfNumber {
			perimeter = append(perimeter, v)
		}
	}

	for i, v := range currentLine {
		if !unicode.IsDigit(v) && v != 46 && (i == number.startPosition-1 || i == number.startPosition+number.lengthOfNumber) {
			perimeter = append(perimeter, v)
		}
	}

	for i, v := range nextLine {
		if !unicode.IsDigit(v) && v != 46 && i >= number.startPosition-1 && i <= number.startPosition+number.lengthOfNumber {
			perimeter = append(perimeter, v)
		}
	}

	number.perimeter = perimeter
}
