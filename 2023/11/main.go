package main

import (
	"advent-of-code/2023/11/helpers"
)

const inputLocation = "sample.txt"

type input struct {
	rawLines   map[int]string
	day1Answer int
	day2Answer int
	spaces     []space
}

type space struct {
	row    int
	column int
	galaxy int
}

func main() {

	var input input
	input.rawLines = helpers.ReadFile(inputLocation)

	galaxy := 0

	emptyRows := make(map[int]bool)
	emptyColumns := make(map[int]bool)

	for i := 0; i < len(input.rawLines[1]); i++ {
		emptyColumns[i] = true
	}

	for row := 0; row <= len(input.rawLines); row++ {

		var spaces []space
		emptyRows[row] = true
		for column, character := range input.rawLines[row+1] {
			var space space
			space.row = row
			space.column = column
			space.galaxy = 0

			if character == 35 {
				galaxy++
				space.galaxy = galaxy
				emptyRows[row] = false
				emptyColumns[column] = false
				spaces = append(spaces, space)
			}

		}

		input.spaces = append(input.spaces, spaces...)
	}

}
