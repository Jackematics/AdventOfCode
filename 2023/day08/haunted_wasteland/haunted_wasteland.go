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

func TotalSimultaneousSteps(raw_input []string) int {
	instructions := raw_input[0]
	node_mappings := extractNodesMappings(raw_input[2:])

	current_nodes := extractStartingNodes(node_mappings)

	lcdSteps := []int{}
	for _, node := range current_nodes {
		steps := 0
		end_char_not_z := true

		for end_char_not_z {
			for _, instruction := range instructions {
				if node[2] == 'Z' {
					lcdSteps = append(lcdSteps, steps)
					end_char_not_z = false
					break
				}

				if instruction == 'L' {
					node = node_mappings[node].left
				}

				if instruction == 'R' {
					node = node_mappings[node].right
				}

				steps++
			}
		}
	}

	return LCM(lcdSteps[0], lcdSteps[1], lcdSteps[2:]...)
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

func extractStartingNodes(nodes map[string]Node) []string {
	starting_nodes := []string{}
	for key := range nodes {
		if key[2] == 'A' {
			starting_nodes = append(starting_nodes, key)
		}
	}

	return starting_nodes
}

func nodesEndOnZ(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}

	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
