import FileProcessor from './file-processor';
import CathodRayTube from './cathod-ray-tube';

const instructions =
  FileProcessor.processInputIntoInstructions('src/input.txt');
const part1Result: number = CathodRayTube.calculateAggregatedOutput(
  instructions,
  [20, 60, 100, 140, 180, 220],
  { width: 40, height: 6 }
).signalStrength;
const part2Result: string[] = CathodRayTube.calculateAggregatedOutput(
  instructions,
  [],
  { width: 40, height: 6 }
).crtOutput;

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
