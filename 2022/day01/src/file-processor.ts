import * as fs from 'fs';

export default class FileProcessor {
  public static processInputIntoInventories(inputPath: string): number[][] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n')
      .reduce(this.rawInventoryReducer, [[]]);
  }

  private static rawInventoryReducer(
    acc: number[][],
    element: string
  ): number[][] {
    if (element === '') {
      return [...acc, []];
    }

    acc[acc.length - 1].push(parseInt(element));
    return acc;
  }
}
