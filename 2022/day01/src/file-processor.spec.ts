import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const result = FileProcessor.processInputIntoInventories(
      'src/example-input.txt'
    );
    const expected = [
      [1000, 2000, 3000],
      [4000],
      [5000, 6000],
      [7000, 8000, 9000],
      [10000],
    ];

    expect(result).toStrictEqual(expected);
  });
});
