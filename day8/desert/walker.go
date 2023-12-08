package desert

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type Node struct {
	Id    string
	Left  string
	Right string
}

func CountSteps(input []string) int {
	instructions := []rune(input[0])
	nodes := map[string]Node{}
	for _, line := range input[2:] {
		cleanLine := strings.ReplaceAll(strings.ReplaceAll(line, "(", ""), ")", "")
		parts := strings.Split(cleanLine, " = ")
		id := parts[0]
		directions := parts[1]
		parts = strings.Split(directions, ", ")
		left := parts[0]
		right := parts[1]
		node := Node{id, left, right}
		nodes[id] = node
	}

	instructionsLength := len(instructions)
	instructionIndex := 0
	stepsTaken := 0
	currentNode := nodes["AAA"]
	for currentNode.Id != "ZZZ" {
		switch instructions[instructionIndex] {
		case 'L':
			currentNode = nodes[currentNode.Left]
		case 'R':
			currentNode = nodes[currentNode.Right]
		}

		stepsTaken++
		instructionIndex++
		if instructionIndex == instructionsLength {
			instructionIndex = 0
		}
	}

	return stepsTaken
}

func CountStepCycles(input []string) int {
	instructions := []rune(input[0])
	nodes := map[string]Node{}
	currentNodes := []Node{}
	for _, line := range input[2:] {
		cleanLine := strings.ReplaceAll(strings.ReplaceAll(line, "(", ""), ")", "")
		parts := strings.Split(cleanLine, " = ")
		id := parts[0]
		directions := parts[1]
		parts = strings.Split(directions, ", ")
		left := parts[0]
		right := parts[1]
		node := Node{id, left, right}
		nodes[id] = node
		if node.Id[2] == 'A' {
			currentNodes = append(currentNodes, node)
		}
	}

	instructionsLength := len(instructions)
	instructionIndex := 0
	stepsTaken := 0
	pathLengths := []int{}
	for len(currentNodes) > 0 {
		nextNodes := []Node{}
		for _, current := range currentNodes {
			if current.Id[2] == 'Z' {
				pathLengths = append(pathLengths, stepsTaken)
				continue
			}

			nextNodes = append(nextNodes, current)
		}

		for i := 0; i < len(nextNodes); i++ {
			switch instructions[instructionIndex] {
			case 'L':
				nextNodes[i] = nodes[nextNodes[i].Left]
			case 'R':
				nextNodes[i] = nodes[nextNodes[i].Right]
			}
		}

		stepsTaken++
		instructionIndex++
		if instructionIndex == instructionsLength {
			instructionIndex = 0
		}

		currentNodes = nextNodes
	}

	stepsNeeded := pathLengths[0]
	for i := 1; i < len(pathLengths); i++ {
		stepsNeeded = kit.LowestCommonMultiple(stepsNeeded, pathLengths[i])
	}

	return stepsNeeded
}
