import FileProcessor from './file-processor';
import RpsStrategyGuide from './rps-strategy-guide';
import ScoringSystems from './scoring-systems';

const games: string[] =
  FileProcessor.processInputIntoInventories('src/input.txt');

const part1Result = RpsStrategyGuide.calculateTotalScore(
  games,
  ScoringSystems.part1ScoringSystem
);

const part2Result = RpsStrategyGuide.calculateTotalScore(
  games,
  ScoringSystems.part2ScoringSystem
);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
