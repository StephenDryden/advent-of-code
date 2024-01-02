package main

import (
	"advent-of-code/2023/15/helpers"
	"fmt"
	"strings"
)

const inputLocation = "input.txt"

type input struct {
	rawLines    map[int]string
	part1Answer int
	sequences   []sequence
}

type sequence struct {
	string string
	value  int
}

func NewInput(inputLocation string) input {
	return input{
		rawLines: helpers.ReadFile(inputLocation),
	}
}

func NewSequence(string string, value int) sequence {
	return sequence{
		string: string,
		value:  value,
	}
}

func main() {

	input := NewInput(inputLocation)
	input.parseSequences()
	input.runHASHAlgorithm()
	input.calculatePart1Answer()

	fmt.Printf("The answer to day 15 part 1 is: %v", input.part1Answer)
}

func (i *input) parseSequences() {
	initialSequences := strings.Split(i.rawLines[1], ",")

	for _, sequence := range initialSequences {
		i.sequences = append(i.sequences, NewSequence(sequence, 0))
	}

}

func (i *input) runHASHAlgorithm() {

	var newSequences []sequence

	for _, sequence := range i.sequences {
		result := 0
		for _, rune := range sequence.string {
			result = ((result + int(rune)) * 17) % 256
		}
		newSequences = append(newSequences, NewSequence(sequence.string, result))
	}

	i.sequences = newSequences
}

func (i *input) calculatePart1Answer() {

	for _, v := range i.sequences {
		i.part1Answer = i.part1Answer + v.value
	}
}
