import Addx from './addx';
import IInstruction from './iinstruction.interface';
import Noop from './noop';

export default class InstructionFactory {
  public static createInstruction(raw: string): IInstruction {
    const split = raw.split(' ');
    switch (split[0]) {
      case 'addx':
        return new Addx(parseInt(split[1]));
      case 'noop':
        return new Noop();
      default:
        throw new Error(`Invalid instruction ${split[0]}`);
    }
  }
}
