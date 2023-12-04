package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//const inputLocation = "input.txt"

const inputLocation = "sample.txt"

type input struct {
	coordinates []coordinate
}

type coordinate struct {
	x     int
	y     int
	value rune
}

type number struct {
	coordinates       []coordinate
	perimeter         string
	isPartOfSchematic bool
}

func main() {

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := parseSchematic(scanner)

	fmt.Printf("The first coordinate is: %v", input.coordinates[0])
}

func parseSchematic(scanner *bufio.Scanner) input {
	var input input
	var coordinate coordinate

	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		for columnNumber, value := range line {
			coordinate.x = rowNumber
			coordinate.y = columnNumber
			coordinate.value = value
			input.coordinates = append(input.coordinates, coordinate)
		}
		rowNumber++
	}
	return input
}
