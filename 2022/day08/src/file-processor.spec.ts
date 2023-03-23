import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const data: number[][] = FileProcessor.processInput(
      'src/example-input.txt'
    );

    expect(data).toStrictEqual([
      [3, 0, 3, 7, 3],
      [2, 5, 5, 1, 2],
      [6, 5, 3, 3, 2],
      [3, 3, 5, 4, 9],
      [3, 5, 3, 9, 0],
    ]);
  });
});
