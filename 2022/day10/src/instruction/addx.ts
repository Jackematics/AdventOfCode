import { InstructionOutput, ScreenDims } from '../types';
import IInstruction from './iinstruction.interface';

export default class Addx implements IInstruction {
  _x: number;

  constructor(x: number) {
    this._x = x;
  }

  public execute(aggregateOutput: InstructionOutput): InstructionOutput {
    return {
      X: aggregateOutput.X + this._x,
      cycle: aggregateOutput.cycle + 2,
      screenDims: aggregateOutput.screenDims,
      crtOutput: this.updateCrt(aggregateOutput),
      ...this.updateSignalStrengthAndIntervalsIfInterval(aggregateOutput),
    };
  }

  private updateCrt(aggregateOutput: InstructionOutput): string[] {
    const { X, cycle, screenDims, crtOutput } = aggregateOutput;

    const firstPixel = this.determinePixelLitOrDark(X, cycle);
    const firstPixelCrtRow = this.calculateCrtRow(cycle, screenDims);

    const secondPixel = this.determinePixelLitOrDark(X, cycle + 1);
    const secondPixelCrtRow = this.calculateCrtRow(cycle + 1, screenDims);

    const updatedCrt = [...crtOutput];
    updatedCrt[firstPixelCrtRow] += firstPixel;
    updatedCrt[secondPixelCrtRow] += secondPixel;

    return updatedCrt;
  }

  private determinePixelLitOrDark(X: number, cycle: number) {
    return [X - 1, X, X + 1].includes(cycle % 40) ? '#' : '.';
  }

  private calculateCrtRow(cycle: number, screenDims: ScreenDims) {
    return Math.floor(cycle / screenDims.width) % screenDims.height;
  }

  private updateSignalStrengthAndIntervalsIfInterval(
    aggregateOutput: InstructionOutput
  ): Pick<InstructionOutput, 'cycleIntervals' | 'signalStrength'> {
    const { X, cycle, signalStrength, cycleIntervals } = aggregateOutput;
    const interval = cycleIntervals[0];

    return this.cycleCrossesInterval(cycle, interval)
      ? {
          signalStrength: signalStrength + X * interval,
          cycleIntervals: cycleIntervals.slice(1),
        }
      : {
          signalStrength,
          cycleIntervals,
        };
  }

  private cycleCrossesInterval(cycle: number, interval: number): boolean {
    return cycle + 1 === interval || cycle + 2 === interval;
  }
}
