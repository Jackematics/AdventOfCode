import AssignmentPair from './assignment-pair';
import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const result: AssignmentPair[] =
      FileProcessor.processInputIntoAssignmentPairs('src/example-input.txt');
    const expected = [
      '2-4,6-8',
      '2-3,4-5',
      '5-7,7-9',
      '2-8,3-7',
      '6-6,4-6',
      '2-6,4-8',
    ];
    result.map((assignmentPair, index) =>
      expect(assignmentPair.printRaw).toBe(expected[index])
    );
  });
});
