import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of strings', () => {
    const result = FileProcessor.processInputIntoInventories(
      'src/example-input.txt'
    );
    const expected = ['A Y', 'B X', 'C Z'];

    expect(result).toStrictEqual(expected);
  });
});
