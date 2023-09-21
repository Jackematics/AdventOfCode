package bitwise_operations

import (
	"math"
	"strconv"
	"strings"
)

func ExecuteBitwiseInstructions(input []string) map[string]interface{} {
	circuit := make(map[string]interface{})

	for _, raw_instruction := range input {
		raw_instruction_split := strings.Split(raw_instruction, " ")

		output_wires := raw_instruction_split[len(raw_instruction_split)-1]

		for i := range output_wires {
			if raw_instruction_split[1] == "->" {
				signal, _ := strconv.Atoi(raw_instruction_split[0])
				circuit[string(output_wires[i])] = signal
			} else if raw_instruction_split[0] == "NOT" {
				operand := raw_instruction_split[1]
				_, operand_in_circuit := circuit[operand]

				if operand_in_circuit {
					operand_int, _ := circuit[operand].(int)

					circuit[string(output_wires[i])] = math.MaxUint16 ^ operand_int
				}
			} else if raw_instruction_split[1] == "AND" || raw_instruction_split[1] == "OR" {
				result, valid := perform_bitwise_andor_operation(
					raw_instruction_split,
					circuit,
					raw_instruction_split[1],
				)

				if valid {
					circuit[string(output_wires[i])] = result
				}
			} else if raw_instruction_split[1] == "LSHIFT" || raw_instruction_split[1] == "RSHIFT" {
				result, valid := perform_bitwise_shift_operation(
					raw_instruction_split,
					circuit,
					raw_instruction_split[1],
				)

				if valid {
					circuit[output_wires] = result
				}
			}
		}
	}

	return circuit
}

func perform_bitwise_andor_operation(
	raw_instruction_split []string,
	circuit map[string]interface{},
	operator string,
) (int, bool) {
	wire_a := raw_instruction_split[0]
	wire_b := raw_instruction_split[2]

	_, wire_a_in_circuit := circuit[wire_a]
	_, wire_b_in_circuit := circuit[wire_b]

	if wire_a_in_circuit && wire_b_in_circuit {
		wire_a_int, _ := circuit[wire_a].(int)
		wire_b_int, _ := circuit[wire_b].(int)
		if operator == "AND" {
			return wire_a_int & wire_b_int, true
		} else if operator == "OR" {
			return wire_a_int | wire_b_int, true
		}
	}

	return 0, false
}

func perform_bitwise_shift_operation(
	raw_instruction_split []string,
	circuit map[string]interface{},
	operator string,
) (int, bool) {
	wire := raw_instruction_split[0]
	_, wire_in_circuit := circuit[wire]

	if wire_in_circuit {
		signal, _ := circuit[wire].(int)
		shift_amount, _ := strconv.Atoi(raw_instruction_split[2])

		if operator == "LSHIFT" {
			return signal << shift_amount, true
		} else if operator == "RSHIFT" {
			return signal >> shift_amount, true
		}

		return 0, false
	}

	return 0, false
}
