package main

import (
	"advent-of-code/2023/1/helpers"
	"fmt"
)

const inputLocation = "input.txt"

type input struct {
	lines       map[int]string
	instruction string
	nodes       map[string]node
	steps       int
	ghostSteps  int
}

type node struct {
	left     string
	right    string
	position string
}

func main() {

	var input input
	input.lines = helpers.ReadFile(inputLocation)

	input.instruction = input.lines[1]

	nodes := make(map[string]node)
	for i := 3; i < len(input.lines)+1; i++ {
		var node node
		node.left = input.lines[i][7:10]
		node.right = input.lines[i][12:15]
		node.position = input.lines[i][0:3]

		if node.left == "" || node.right == "" || node.position == "" {
			fmt.Print("oh no")

		}

		nodes[input.lines[i][0:3]] = node
	}
	input.nodes = nodes
	//input.followInstructions()
	input.CalculateGhostSteps()
	fmt.Printf("The answer to day 8 part 1 is %v\n", input.steps)
	fmt.Printf("The answer to day 8 part 2 is %v", input.ghostSteps)

}

func (input *input) followInstructions() {

	endReached := false
	nextNode := "AAA"

	for !endReached {
		for _, instruction := range input.instruction {
			nextnode := input.nodes[nextNode]

			if instruction == 76 {
				nextNode = nextnode.left
			}
			if instruction == 82 {
				nextNode = nextnode.right
			}
			input.steps++
			if nextNode == "ZZZ" {
				endReached = true
				break
			}
		}
	}

}

// does not work
func (input *input) CalculateGhostSteps() {

	count := 0
	for !allNodesOnZ(input.nodes) {
		for _, instruction := range input.instruction {
			for key, node := range input.nodes {

				if key[2:3] == "A" {

					if instruction == 76 {
						node.position = input.nodes[node.position].left
					}
					if instruction == 82 {
						node.position = input.nodes[node.position].right
					}
					input.nodes[key] = node
				}
			}
			count++

		}
	}
	input.ghostSteps = count
}

func allNodesOnZ(nodes map[string]node) bool {

	totalA := 0
	totalZ := 0

	for key, node := range nodes {
		if key[2:3] == "A" {
			totalA++
			if node.position[2:3] == "Z" {
				totalZ++
			}

		}
	}

	if totalZ > 3 {
		fmt.Printf("Total A's on Z : %v : %v \n", totalA, totalZ)
	}

	return totalA == totalZ
}
