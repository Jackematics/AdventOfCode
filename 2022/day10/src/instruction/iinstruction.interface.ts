import { InstructionOutput } from '../types';

interface IInstruction {
  execute(aggregateOutput: InstructionOutput): InstructionOutput;
}

export default IInstruction;
