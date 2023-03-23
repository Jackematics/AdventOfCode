import { InstructionOutput } from '../types';
import IInstruction from './iinstruction.interface';

export default class Noop implements IInstruction {
  public execute(aggregateOutput: InstructionOutput): InstructionOutput {
    return {
      X: aggregateOutput.X,
      cycle: aggregateOutput.cycle + 1,
      screenDims: aggregateOutput.screenDims,
      crtOutput: this.updateCrt(aggregateOutput),
      ...this.updateSignalStrengthAndIntervalsIfInterval(aggregateOutput),
    };
  }

  private updateCrt(aggregateOutput: InstructionOutput): string[] {
    const { X, cycle, screenDims, crtOutput } = aggregateOutput;
    const crtRow = Math.floor(cycle / screenDims.width) % screenDims.height;
    const spriteIndexes = [X - 1, X, X + 1];

    const updatedCrt = [...crtOutput];
    updatedCrt[crtRow] += spriteIndexes.includes(cycle % 40) ? '#' : '.';

    return updatedCrt;
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
    return cycle + 1 === interval;
  }
}
