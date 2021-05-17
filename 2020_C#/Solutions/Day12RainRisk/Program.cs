using System;
using System.Collections.Generic;
using PuzzleInputRetriever;

using Day12RainRisk.Actions;

namespace Day12RainRisk
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> puzzleInput = Retriever.Retrieve(@"..//..//..//Day12PuzzleInput.txt");
            List<IInstruction> instructions = ConvertToInstructions(puzzleInput);

            Navigation navigation = new Navigation(instructions);

            Console.WriteLine(navigation.GetPart1ManhattanDistance());
            Console.WriteLine(navigation.GetPart2ManhattanDistance());
        }

        private static List<IInstruction> ConvertToInstructions(List<string> input)
        {
            var instructions = new List<IInstruction>();

            foreach (string line in input)
            {
                char action = line[0];
                int value = int.Parse(line.Substring(1, line.Length - 1));

                switch (action)
                {
                    case 'N':
                        instructions.Add(new North(value));
                        break;

                    case 'E':
                        instructions.Add(new East(value));
                        break;

                    case 'S':
                        instructions.Add(new South(value));
                        break;

                    case 'W':
                        instructions.Add(new West(value));
                        break;

                    case 'F':
                        instructions.Add(new Forward(value));
                        break;

                    case 'R':
                        instructions.Add(new Right(value));
                        break;

                    case 'L':
                        instructions.Add(new Left(value));
                        break;
                }
            }

            return instructions;
        }
    }
}
