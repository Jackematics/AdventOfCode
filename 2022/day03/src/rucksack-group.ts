import PriorityCalculator from './priority-calculator';
import Rucksack from './rucksack';

export default class RucksackGroup {
  private readonly rucksackGroup: Rucksack[];
  private readonly expectedGroupSize = 3;

  constructor(rucksacks: Rucksack[] = []) {
    this.rucksackGroup = rucksacks;
  }

  public appendRucksack(rucksack: Rucksack): void {
    this.rucksackGroup.push(rucksack);
  }

  public get length() {
    return this.rucksackGroup.length;
  }

  public get groupContents() {
    return this.rucksackGroup.map((rucksack) => rucksack.contents);
  }

  public calculateSharedPriority(): number {
    if (this.rucksackGroup.length !== this.expectedGroupSize)
      throw new Error(`group must have ${this.expectedGroupSize} rucksacks`);

    const sharedType = this.rucksackGroup[0].contentsArray.find(
      (itemType: string) =>
        this.rucksackGroup[1].contentsArray.includes(itemType) &&
        this.rucksackGroup[2].contentsArray.includes(itemType)
    );

    if (!sharedType) throw new Error('group must share an item type');

    return PriorityCalculator.calculatePriority(sharedType);
  }
}
