import Monkey from './monkey';

export default class MonkeyBusinessAnalyser {
  public static calculateMonkeyBusinessLevel(monkeys: Monkey[]): number {
    return monkeys
      .sort((a, b) => b.itemInspectionCount - a.itemInspectionCount)
      .slice(0, 2)
      .reduce(
        (acc: number, monkey: Monkey) => acc * monkey.itemInspectionCount,
        1
      );
  }
}
