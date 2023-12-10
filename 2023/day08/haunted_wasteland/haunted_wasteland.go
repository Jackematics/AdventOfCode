package haunted_wasteland

import "strings"

type Node struct {
	left  string
	right string
}

func TotalSteps(raw_input []string) int {
	instructions := raw_input[0]
	nodes := extractNodesMappings(raw_input[2:])

	current_node := "AAA"
	end_node := "ZZZ"

	steps := 0
	for current_node != end_node {
		for _, instruction := range instructions {
			if current_node == end_node {
				break
			}

			if instruction == 'L' {
				current_node = nodes[current_node].left
			}

			if instruction == 'R' {
				current_node = nodes[current_node].right
			}

			steps++
		}
	}

	return steps
}

func extractNodesMappings(raw_input []string) map[string]Node {
	node_mappings := map[string]Node{}

	for _, line := range raw_input {
		whitespace_split := strings.Split(line, " ")

		element := whitespace_split[0]
		left := whitespace_split[2][1:4]
		right := whitespace_split[3][:3]

		node_mappings[element] = Node{left: left, right: right}
	}

	return node_mappings
}
