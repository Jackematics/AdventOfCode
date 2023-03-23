import { Part } from './enum';
import FileProcessor from './file-processor';
import MonkeyBusinessAnalyser from './monkey-analyser';
import RouteCharter from './route-charter';

const monkeys = FileProcessor.processInput('src/input.txt');
// MonkeyModeller.playRounds(monkeys, 20, Part.One);
RouteCharter.playRounds(monkeys, 10000, Part.Two);

const result = MonkeyBusinessAnalyser.calculateMonkeyBusinessLevel(monkeys);
// const part2Result =

console.log('Part 1 result: ', result);
// console.log('Part 2 result: ', part2Result);
