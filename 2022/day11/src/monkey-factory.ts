import Monkey from './monkey';

export default class MonkeyFactory {
  public static createMonkey(rawMonkey: string): Monkey {
    const lines = rawMonkey.split('\n');

    return new Monkey({
      items: JSON.parse('[' + lines[1].split(': ')[1] + ']'),
      divisor: parseInt(lines[3].split(' ')[5]),
      operation: this.constructOperationFunction(lines),
      test: this.constructTestFunction(lines),
    });
  }

  private static constructOperationFunction(lines: string[]) {
    const rawOperation = lines[2].split(' ');

    return (oldWorry: number) =>
      this.calculateOperation(oldWorry, rawOperation[6], rawOperation[7]);
  }

  private static calculateOperation(
    oldWorry: number,
    rawOperand: string,
    rawArgument: string
  ): number {
    const validIntegerOperation = this.integerOperationIfValidInt(
      oldWorry,
      rawOperand,
      rawArgument
    );
    return validIntegerOperation
      ? validIntegerOperation
      : this.selfOperation(oldWorry, rawOperand);
  }

  private static integerOperationIfValidInt(
    oldWorry: number,
    rawOperand: string,
    rawArgument: string
  ): number | undefined {
    const argument = parseInt(rawArgument);
    if (argument) {
      return {
        '+': oldWorry + argument,
        '-': oldWorry - argument,
        '*': oldWorry * argument,
        '/': oldWorry / argument,
      }[rawOperand];
    }

    return undefined;
  }

  private static selfOperation(oldWorry: number, rawOperand: string): number {
    return (
      {
        '+': 2 * oldWorry,
        '-': 0,
        '*': oldWorry * oldWorry,
        '/': 1,
      }[rawOperand] ?? 0
    );
  }

  private static constructTestFunction(lines: string[]) {
    const testDivisor = parseInt(lines[3].split(' ')[5]);
    const trueMonkey = parseInt(lines[4].split(' ')[9]);
    const falseMonkey = parseInt(lines[5].split(' ')[9]);

    return (worry: number) =>
      this.calculateTest(worry, testDivisor, trueMonkey, falseMonkey);
  }

  private static calculateTest(
    worry: number,
    testDivisor: number,
    trueMonkey: number,
    falseMonkey: number
  ): number {
    return worry % testDivisor === 0 ? trueMonkey : falseMonkey;
  }
}
