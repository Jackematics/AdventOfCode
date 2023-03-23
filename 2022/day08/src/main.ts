import FileProcessor from './file-processor';
import TreePatchAnalyser from './tree-patch-analyser';

const treePatch = FileProcessor.processInput('src/input.txt');
const part1Result: number = TreePatchAnalyser.visibleTreesCount(treePatch);
const part2Result: number =
  TreePatchAnalyser.calculateMaxScenicScore(treePatch);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
