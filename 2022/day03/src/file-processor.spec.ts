import FileProcessor from './file-processor';
import Rucksack from './rucksack';
import RucksackGroup from './rucksack-group';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const result: Rucksack[] = FileProcessor.processInputIntoRucksacks(
      'src/example-input.txt'
    );

    const expected = [
      'vJrwpWtwJgWrhcsFMMfFFhFp',
      'jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL',
      'PmmdzqPrVvPwwTWBwg',
      'wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn',
      'ttgJtRGJQctTZtZT',
      'CrZsJsPPZsGzwwsLwLmpwMDw',
    ];

    result.map((rucksack, index) =>
      expect(rucksack.contents).toBe(expected[index])
    );
  });

  it('should convert a text file of lines into an array of Rucksack groups', () => {
    const result: RucksackGroup[] =
      FileProcessor.processInputIntoRucksackGroups('src/example-input.txt');

    const expected = [
      [
        'vJrwpWtwJgWrhcsFMMfFFhFp',
        'jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL',
        'PmmdzqPrVvPwwTWBwg',
      ],
      [
        'wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn',
        'ttgJtRGJQctTZtZT',
        'CrZsJsPPZsGzwwsLwLmpwMDw',
      ],
    ];

    result.map((rucksackGroup, index) =>
      expect(rucksackGroup.groupContents).toStrictEqual(expected[index])
    );
  });
});
