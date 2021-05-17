using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Day8HandheldHalting
{
    public class Executor
    {
        public int Accumulator { get; set; }
        private List<IInstruction> Instructions { get; set; }
        private List<int> ExecutedInstructionIndexes { get; set; } = new List<int>();
        private bool ProgramTerminates => Instructions.Count == ExecutedInstructionIndexes.Count;
        public bool Terminated = false;

        public Executor(List<IInstruction> instructions)
        {
            Instructions = instructions;
        }

        public void Execute()
        {
            int currentIndex = 0;
            Terminated = false;

            do
            {
                if (currentIndex == Instructions.Count || currentIndex < 0)
                {
                    return;
                }

                ExecutedInstructionIndexes.Add(currentIndex);

                IInstruction currentInstruction = Instructions[currentIndex];

                if (currentInstruction is Accumulate)
                {
                    Accumulator += currentInstruction.Argument;
                }

                if (currentIndex == Instructions.Count - 1)
                {
                    Terminated = true;
                    return;
                }

                currentIndex += currentInstruction.IndexMoveAmount;
            }
            while (!ExecutedInstructionIndexes.Contains(currentIndex));
        }

        public void FullyExecute()
        {
            int startIndex = 0;
            var tempInstructions = new List<IInstruction>(Instructions);
            while (Terminated == false)
            {
                ExecutedInstructionIndexes.Clear();
                Accumulator = 0;                

                for (int i = startIndex; i < Instructions.Count; i++)
                {
                    if (Instructions[i] is NoOperation)
                    {
                        Instructions[i] = SwapToJump(Instructions[i]);
                        startIndex = i + 1;
                        break;
                    }

                    if (Instructions[i] is Jump)
                    {
                        Instructions[i] = SwapToNOP(Instructions[i]);
                        startIndex = i + 1;
                        break;
                    }
                }                

                Execute();
                Instructions = new List<IInstruction>(tempInstructions);
            }
        }

        private IInstruction SwapToNOP(IInstruction instruction)
        {
            return new NoOperation(instruction.Argument);
        }

        private IInstruction SwapToJump(IInstruction instruction)
        {
            return new Jump(instruction.Argument);
        }
    }
}
