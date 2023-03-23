import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of lines into an array of number arrays', () => {
    const data: string = FileProcessor.processInput('src/example-input.txt');

    expect(data).toBe('mjqjpqmgbljsphdztnvjfqwrcgsmlb');
  });
});
