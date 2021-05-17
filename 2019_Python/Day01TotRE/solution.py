def get_masses():
    with open("./Day01TotRE/input.txt") as file:
        masses = file.read().splitlines()
    return masses

get_required_fuel = lambda mass: int(mass) // 3 - 2

def get_nested_fuel(mass):
    total_fuel = 0
    current_fuel = get_required_fuel(mass)

    while (current_fuel >= 0):
        total_fuel += current_fuel
        current_fuel = get_required_fuel(current_fuel)

    return total_fuel


def get_fuel_requirement_sum(part):
    count = 0
    for mass in get_masses():
        count += get_required_fuel(mass) if part == "part1" else get_nested_fuel(mass)

    return count

print(get_fuel_requirement_sum("part1"))
print(get_fuel_requirement_sum("part2"))
