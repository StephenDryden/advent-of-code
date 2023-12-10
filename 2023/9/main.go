package main

import (
	"advent-of-code/2023/1/helpers"
	"fmt"
	"strconv"
	"strings"
)

const inputLocation = "input.txt"

type input struct {
	rawLines   map[int]string
	day1Answer int
	day2Answer int
	lines      []line
}

type line struct {
	rows []row
}

type row struct {
	numbers    map[int]int
	difference map[int]int
	prediction int
	history    int
}

func main() {

	var input input
	var lines []line
	input.rawLines = helpers.ReadFile(inputLocation)

	for _, rawLine := range input.rawLines {
		var line line
		line.CalculateRows(rawLine)
		line.CalculatePredictions()
		lines = append(lines, line)
	}

	input.lines = lines

	input.CalculatePart1Answer()
	input.CalculatePart2Answer()

	fmt.Printf("The answer to day 9 part 1 is: %v\n", input.day1Answer)
	fmt.Printf("The answer to day 9 part 2 is: %v\n", input.day2Answer)

}

func (line *line) CalculateRows(rawLine string) {
	var rows []row
	var row row
	numbers := make(map[int]int)

	for i, stringValue := range strings.Split(rawLine, " ") {
		integerValue, err := strconv.Atoi(stringValue)
		if err != nil {
			panic(err)
		}
		numbers[i] = integerValue
	}
	row.numbers = numbers
	row.difference = calculateDifference(row.numbers)
	rows = append(rows, row)

	for !differenceEqualToZero(row.difference) {
		row.numbers = row.difference
		row.difference = calculateDifference(row.numbers)
		rows = append(rows, row)
	}
	line.rows = rows
}

func (line *line) CalculatePredictions() {

	for i := len(line.rows) - 1; i > -1; i-- {
		countNumbers := len(line.rows[i].numbers)

		lastNumberInRow := line.rows[i].numbers[countNumbers-1]
		firstNumberInRow := line.rows[i].numbers[0]

		if differenceEqualToZero(line.rows[i].difference) {
			line.rows[i].prediction = lastNumberInRow
			line.rows[i].history = lastNumberInRow

		} else {
			line.rows[i].prediction = lastNumberInRow + line.rows[i+1].prediction
			line.rows[i].history = firstNumberInRow - line.rows[i+1].history
		}
	}

}

func calculateDifference(row map[int]int) map[int]int {
	difference := make(map[int]int)

	for i := 0; i < len(row)-1; i++ {
		difference[i] = row[i+1] - row[i]
	}

	return difference
}

func differenceEqualToZero(difference map[int]int) bool {

	for i := 0; i < len(difference); i++ {
		if difference[i] != 0 {
			return false
		}
	}
	return true

}

func (input *input) CalculatePart1Answer() {

	answer := 0

	for _, line := range input.lines {
		answer = answer + line.rows[0].prediction

	}
	input.day1Answer = answer
}

func (input *input) CalculatePart2Answer() {

	answer := 0

	for _, line := range input.lines {
		answer = answer + line.rows[0].history

	}
	input.day2Answer = answer
}
