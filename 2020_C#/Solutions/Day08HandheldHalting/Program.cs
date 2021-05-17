using System;
using System.Collections.Generic;
using System.IO;

using PuzzleInputRetriever;

namespace Day8HandheldHalting
{
    class Program
    {
        private static readonly List<string> _PuzzleInput = GetPuzzleInput();

        static void Main(string[] args)
        {
            //Test Example
            List<string> testInput = new List<string>()
            {
                "nop +0",
                "acc +1",
                "jmp +4",
                "acc +3",
                "jmp -3",
                "acc -99",
                "acc +1",
                "jmp -4",
                "acc +6"
            };

            //Part 1 Test
            List<IInstruction> testInstructions = ConvertToInstructions(testInput);

            var part1TestExecutor = new Executor(testInstructions);
            part1TestExecutor.Execute();
            Console.WriteLine(part1TestExecutor.Accumulator);

            //Part 1
            List<IInstruction> instructions = ConvertToInstructions(_PuzzleInput);   
            var part1Executor = new Executor(instructions);
            part1Executor.Execute();

            Console.WriteLine(part1Executor.Accumulator);

            //Part 2 Test
            var part2TestExecutor = new Executor(testInstructions);
            part2TestExecutor.FullyExecute();
            Console.WriteLine(part2TestExecutor.Accumulator);

            //Part 2
            var part2Executor = new Executor(instructions);
            part2Executor.FullyExecute();
            Console.WriteLine(part2Executor.Accumulator);
        }

        private static List<IInstruction> ConvertToInstructions(List<string> input)
        {
            var instructions = new List<IInstruction>();

            foreach (string line in input)
            {
                string[] instruction = line.Split(" ");
                string operation = instruction[0];
                string argument = instruction[1];

                switch (operation)
                {
                    case "acc":
                        var accumulate = new Accumulate(int.Parse(argument));
                        instructions.Add(accumulate);
                        break;

                    case "jmp":
                        var jump = new Jump(int.Parse(argument));
                        instructions.Add(jump);
                        break;

                    case "nop":
                        var noOperation = new NoOperation(int.Parse(argument));
                        instructions.Add(noOperation);
                        break;
                }
            }

            return instructions;
        }

        private static List<string> GetPuzzleInput()
        {
            var puzzleInput = new List<string>();

            using (var reader = new StreamReader(@"..//..//..//Day8PuzzleInput.txt"))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    puzzleInput.Add(line);
                }
            }

            return puzzleInput;
        }
    }
}
