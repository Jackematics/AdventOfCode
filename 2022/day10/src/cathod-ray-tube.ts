import { IInstruction } from './instruction';
import { InstructionOutput, ScreenDims } from './types';

export default class CathodRayTube {
  static calculateAggregatedOutput(
    instructions: IInstruction[],
    cycleIntervals: number[],
    screenDims: ScreenDims
  ): InstructionOutput {
    const aggregateOutput: InstructionOutput = {
      X: 1,
      cycle: 0,
      signalStrength: 0,
      cycleIntervals,
      screenDims,
      crtOutput: Array(screenDims.height).fill(''),
    };

    return instructions.reduce(
      CathodRayTube.instructionsReducer,
      aggregateOutput
    );
  }

  private static instructionsReducer(
    acc: InstructionOutput,
    instruction: IInstruction
  ) {
    return instruction.execute(acc);
  }
}
