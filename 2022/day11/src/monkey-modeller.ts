import { Part } from './enum';
import Monkey from './monkey';

export default class MonkeyModeller {
  public static playRounds(monkeys: Monkey[], rounds: number, part: Part) {
    const modulus = monkeys.reduce(
      (acc: number, monkey: Monkey) => acc * monkey.divisor,
      1
    );

    for (let i = 0; i < rounds; i++) {
      monkeys.forEach((monkey) => {
        part === Part.One
          ? monkey.part1TakeTurn()
          : monkey.part2TakeTurn(modulus);

        const thrownItems = monkey.throwItems();

        thrownItems.forEach((monkeyItem) => {
          monkeys[monkeyItem.monkey].receiveItem(monkeyItem.item);
        });
      });
    }
  }
}
