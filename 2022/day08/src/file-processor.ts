import * as fs from 'fs';

export default class FileProcessor {
  public static processInput(inputPath: string): number[][] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n')
      .map((rawTreeRow: string) =>
        rawTreeRow.split('').map((tree) => parseInt(tree))
      );
  }
}
