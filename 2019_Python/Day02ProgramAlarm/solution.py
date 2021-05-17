from instruction import Instruction

def get_file_data():
    with open("./Day02ProgramAlarm/input.txt") as file:
        data = [int(line) for line in file.read().split(',')]
    return data
     
def load_instructions():
    file_data = get_file_data()
    int_codes = []
    instruction = Instruction()

    for i in range(len(file_data)):        
        if (i % 4 == 0):
            instruction = Instruction()
            instruction.opcode = file_data[i]
        elif (i % 3 == 0 and i != 0):
            instruction.parameters.append(file_data[i])
            int_codes.append(instruction)
        else:
            instruction.parameters.append(file_data[i])

    return int_codes

def execute_opcode(opcode, current_position):
    if (opcode == 1):
        INT_CODES[INT_CODES[current_position + 3]] = INT_CODES[INT_CODES[current_position + 1]] + INT_CODES[INT_CODES[current_position + 2]]
    elif (opcode == 2):
        INT_CODES[current_position + 3] = INT_CODES[INT_CODES[current_position + 1]] * INT_CODES[INT_CODES[current_position + 2]]
    elif (opcode == 99):
        print("Program has now finished")
    else:
        raise ValueError("Opcode must be 1, 2, or 99")

def execute_program():
    current_position = 0
    INT_CODES[1] = 12
    INT_CODES[2] = 2

    while (INT_CODES[current_position] != 99):
        current_op_code = INT_CODES[current_position]
        execute_opcode(current_op_code, current_position)
        current_position += 4

    return INT_CODES[0]

print(execute_program())
