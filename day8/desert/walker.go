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
	nodes := parseNodes(input[2:])
	steps := countStepsTaken(instructions, nodes)
	return steps
}

func CountStepCycles(input []string) int {
	instructions := []rune(input[0])
	nodes := parseNodes(input[2:])
	pathLengths := findPathLengths(instructions, nodes)
	steps := countStepsNeeded(pathLengths)
	return steps
}

func parseNodes(lines []string) map[string]Node {
	nodes := map[string]Node{}
	for _, line := range lines {
		cleanLeft := strings.ReplaceAll(line, "(", "")
		cleanLine := strings.ReplaceAll(cleanLeft, ")", "")
		parts := strings.Split(cleanLine, " = ")
		id := parts[0]
		directions := parts[1]
		parts = strings.Split(directions, ", ")
		left := parts[0]
		right := parts[1]
		node := Node{id, left, right}
		nodes[id] = node
	}

	return nodes
}

func countStepsTaken(instructions []rune, nodes map[string]Node) int {
	instruction := 0
	stepsTaken := 0
	currentNode := nodes["AAA"]
	for currentNode.Id != "ZZZ" {
		switch instructions[instruction] {
		case 'L':
			currentNode = nodes[currentNode.Left]
		case 'R':
			currentNode = nodes[currentNode.Right]
		}

		stepsTaken++
		instruction++
		if instruction == len(instructions) {
			instruction = 0
		}
	}

	return stepsTaken
}

func startNodes(nodes map[string]Node) []Node {
	startNodes := []Node{}
	for _, node := range nodes {
		if node.Id[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}

	return startNodes
}

func findPathLengths(instructions []rune, nodes map[string]Node) []int {
	currentNodes := startNodes(nodes)
	instruction := 0
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
			switch instructions[instruction] {
			case 'L':
				nextNodes[i] = nodes[nextNodes[i].Left]
			case 'R':
				nextNodes[i] = nodes[nextNodes[i].Right]
			}
		}

		stepsTaken++
		instruction++
		if instruction == len(instructions) {
			instruction = 0
		}

		currentNodes = nextNodes
	}

	return pathLengths
}

func countStepsNeeded(pathLengths []int) int {
	stepsNeeded := pathLengths[0]
	for i := 1; i < len(pathLengths); i++ {
		stepsNeeded = kit.LowestCommonMultiple(stepsNeeded, pathLengths[i])
	}

	return stepsNeeded
}
