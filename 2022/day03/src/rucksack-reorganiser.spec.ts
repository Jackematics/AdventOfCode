import Rucksack from './rucksack';
import RucksackGroup from './rucksack-group';
import RucksackReorganiser from './rucksack-reorganiser';

describe('RucksackReorganiser()', () => {
  it.each([
    { input: [new Rucksack('vJrwpWtwJgWrhcsFMMfFFhFp')], expected: 16 },
    { input: [new Rucksack('jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL')], expected: 38 },
    { input: [new Rucksack('PmmdzqPrVvPwwTWBwg')], expected: 42 },
    { input: [new Rucksack('wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn')], expected: 22 },
    { input: [new Rucksack('ttgJtRGJQctTZtZT')], expected: 20 },
    { input: [new Rucksack('CrZsJsPPZsGzwwsLwLmpwMDw')], expected: 19 },
  ])(
    `should return the priority $expected of the item type shared between two rucksack compartments`,
    ({ input, expected }) => {
      expect(RucksackReorganiser.calculateSharedPriorityTotal(input)).toBe(
        expected
      );
    }
  );

  it('should calculate the priority sum of multiple rucksack compartments', () => {
    const input: Rucksack[] = [
      new Rucksack('vJrwpWtwJgWrhcsFMMfFFhFp'),
      new Rucksack('jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL'),
      new Rucksack('PmmdzqPrVvPwwTWBwg'),
      new Rucksack('wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn'),
      new Rucksack('ttgJtRGJQctTZtZT'),
      new Rucksack('CrZsJsPPZsGzwwsLwLmpwMDw'),
    ];

    expect(RucksackReorganiser.calculateSharedPriorityTotal(input)).toBe(157);
  });

  it.each([
    {
      input: [
        new RucksackGroup([
          new Rucksack('vJrwpWtwJgWrhcsFMMfFFhFp'),
          new Rucksack('jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL'),
          new Rucksack('PmmdzqPrVvPwwTWBwg'),
        ]),
      ],
      expected: 18,
    },
    {
      input: [
        new RucksackGroup([
          new Rucksack('wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn'),
          new Rucksack('ttgJtRGJQctTZtZT'),
          new Rucksack('CrZsJsPPZsGzwwsLwLmpwMDw'),
        ]),
      ],
      expected: 52,
    },
  ])(
    `should return the priority $expected of the item type shared between a group of rucksacks`,
    ({ input, expected }) => {
      expect(RucksackReorganiser.calculateGroupSharedPriorityTotal(input)).toBe(
        expected
      );
    }
  );

  it('should calculate the priority sum of multiple rucksack compartments', () => {
    const input: RucksackGroup[] = [
      new RucksackGroup([
        new Rucksack('vJrwpWtwJgWrhcsFMMfFFhFp'),
        new Rucksack('jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL'),
        new Rucksack('PmmdzqPrVvPwwTWBwg'),
      ]),
      new RucksackGroup([
        new Rucksack('wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn'),
        new Rucksack('ttgJtRGJQctTZtZT'),
        new Rucksack('CrZsJsPPZsGzwwsLwLmpwMDw'),
      ]),
    ];

    expect(RucksackReorganiser.calculateGroupSharedPriorityTotal(input)).toBe(
      70
    );
  });
});
