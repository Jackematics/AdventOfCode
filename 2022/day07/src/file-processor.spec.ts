import IDirectory from './directory/idirectory.interface';
import FileProcessor from './file-processor';

describe('FileProcessor()', () => {
  it('should convert a text file of commands into a file system', () => {
    const data: IDirectory = FileProcessor.processInput(
      'src/example-input.txt'
    );
    const expected = {
      '/': {
        a: {
          e: {
            i: 584,
          },
          f: 29116,
          g: 2557,
          'h.lst': 62596,
        },
        d: {
          j: 4060174,
          'd.log': 8033020,
          'd.ext': 5626152,
          k: 7214296,
        },
        'b.txt': 14848514,
        'c.dat': 8504156,
      },
    };

    expect(data.toJson()).toStrictEqual(expected);
  });
});
