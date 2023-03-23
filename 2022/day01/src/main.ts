import CalorieCounter from './calorie-counter';
import FileProcessor from './file-processor';

const data = FileProcessor.processInputIntoInventories('src/input.txt');
const part1Result = CalorieCounter.calculateMaxInventoryValue(data);
const part2Result = CalorieCounter.calculateSumOfTopThree(data);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
