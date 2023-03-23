import FileProcessor from './file-processor';
import RopePhysicsModeller from './rope-physics-modeller';

const motions = FileProcessor.processInput('src/input.txt');
const part1Result: number = RopePhysicsModeller.calculateVisitedPositions(
  motions,
  2
);
const part2Result: number = RopePhysicsModeller.calculateVisitedPositions(
  motions,
  10
);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
