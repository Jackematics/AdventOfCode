import FileProcessor from './file-processor';
import RucksackGroup from './rucksack-group';
import RucksackReorganiser from './rucksack-reorganiser';

const rucksacks = FileProcessor.processInputIntoRucksacks('src/input.txt');
const part1Result = RucksackReorganiser.calculateSharedPriorityTotal(rucksacks);

const rucksackGroups =
  FileProcessor.processInputIntoRucksackGroups('src/input.txt');

const part2Result =
  RucksackReorganiser.calculateGroupSharedPriorityTotal(rucksackGroups);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
