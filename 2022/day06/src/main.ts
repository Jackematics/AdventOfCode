import FileProcessor from './file-processor';
import SignalProcessor from './signal-processor';

const data = FileProcessor.processInput('src/input.txt');
const part1Result: number = SignalProcessor.findStartOfPacketMarkerIndex(data);
const part2Result: number = SignalProcessor.findStartOfMessageMarkerIndex(data);

console.log('Part 1 result: ', part1Result);
console.log('Part 2 result: ', part2Result);
