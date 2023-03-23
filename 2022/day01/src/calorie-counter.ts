export default class CalorieCounter {
  public static calculateMaxInventoryValue(elfInventories: number[][]): number {
    return Math.max(...elfInventories.reduce(this.sumInventoriesReducer, []));
  }

  public static calculateSumOfTopThree(elfInventories: number[][]): number {
    const topThree = elfInventories
      .reduce(this.sumInventoriesReducer, [])
      .sort((a, b) => b - a)
      .slice(0, 3);

    return this.sumCalories(topThree);
  }

  private static sumInventoriesReducer(
    acc: number[],
    element: number[]
  ): number[] {
    return [...acc, CalorieCounter.sumCalories(element)];
  }

  public static sumCalories(inventory: number[]): number {
    return inventory.reduce((a: number, b: number) => a + b, 0);
  }
}
