import * as fs from 'fs';

export default class FileProcessor {
  public static processInput(inputPath: string): string[] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n\n')
      .map((rawMonkey: string) => {});
  }
}
