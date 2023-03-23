import Rucksack from './rucksack';
import RucksackGroup from './rucksack-group';

export default class RucksackReorganiser {
  public static calculateSharedPriorityTotal(rucksacks: Rucksack[]): number {
    return rucksacks.reduce(
      (acc: number, rucksack: Rucksack) =>
        acc + rucksack.calculateSharedPriority(),
      0
    );
  }

  public static calculateGroupSharedPriorityTotal(
    rucksackGroups: RucksackGroup[]
  ): number {
    return rucksackGroups.reduce(
      (acc: number, rucksackGroup: RucksackGroup) =>
        acc + rucksackGroup.calculateSharedPriority(),
      0
    );
  }
}
