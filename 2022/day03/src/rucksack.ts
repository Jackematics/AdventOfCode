import PriorityCalculator from './priority-calculator';

export default class Rucksack {
  private readonly rucksackContents: string[];
  private readonly leftCompartment: string[];
  private readonly rightCompartment: string[];

  constructor(rucksackContents: string) {
    this.rucksackContents = rucksackContents.split('');

    const midIndex = Math.floor(rucksackContents.length / 2);
    this.leftCompartment = rucksackContents.slice(0, midIndex).split('');
    this.rightCompartment = rucksackContents.slice(midIndex).split('');
  }

  get contents(): string {
    return this.rucksackContents.join('');
  }

  get contentsArray(): string[] {
    return this.rucksackContents;
  }

  public calculateSharedPriority(): number {
    const sharedItemType = this.findSharedItemType();
    if (!sharedItemType)
      throw new Error('Compartments must share an item type');

    return PriorityCalculator.calculatePriority(sharedItemType);
  }

  private findSharedItemType(): string | undefined {
    return this.leftCompartment.find((itemType: string) =>
      this.rightCompartment.includes(itemType)
    );
  }
}
