import IDirectory from './directory/idirectory.interface';
import FileProcessor from './file-processor';

const dir: IDirectory = FileProcessor.processInput('src/input.txt');
const part1DirectoriesBelowSizeCap = dir.getDirectoriesBelowSizeCap(100000);

const part1Result: number = part1DirectoriesBelowSizeCap.reduce(
  (acc, size) => acc + size,
  0
);

const rootDirSize = dir.getDirectorySize();
const part2DirectoriesBelowSizeCap = dir
  .getDirectoriesBelowSizeCap(70000000)
  .sort((a, b) => b - a)
  .filter((size) => size >= 30000000 - (70000000 - rootDirSize));

const part2Result: number = Math.min(...part2DirectoriesBelowSizeCap);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
