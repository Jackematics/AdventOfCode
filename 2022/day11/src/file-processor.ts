import * as fs from 'fs';
import Monkey from './monkey';
import MonkeyFactory from './monkey-factory';

export default class FileProcessor {
  public static processInput(inputPath: string): Monkey[] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n\n')
      .map((rawMonkey: string) => {
        return MonkeyFactory.createMonkey(rawMonkey);
      });
  }
}
