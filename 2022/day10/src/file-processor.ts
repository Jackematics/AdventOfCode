import * as fs from 'fs';
import { IInstruction } from './instruction';
import InstructionFactory from './instruction/instruction-factory';

export default class FileProcessor {
  public static processInputIntoInstructions(
    inputPath: string
  ): IInstruction[] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n')
      .map((rawInstruction: string) => {
        return InstructionFactory.createInstruction(rawInstruction);
      });
  }
}
