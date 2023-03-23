import * as fs from 'fs';

export default class FileProcessor {
  public static processInputIntoInventories(inputPath: string): string[] {
    return fs.readFileSync(inputPath).toString().split('\n');
  }
}
