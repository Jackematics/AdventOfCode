using System;
using System.Collections;
using System.Collections.Generic;

using PuzzleInputRetriever;

namespace Day14DockingData
{
    class Program
    {
        static void Main(string[] args)
        {
            var testInput = new List<string>
            {
                "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
                "mem[8] = 11",
                "mem[7] = 101",
                "mem[8] = 0"
            };

            List<ICommand> testProgram = ConvertToInitProgram(testInput);
            var testInitialiser = new Initialiser(testProgram);
            testInitialiser.Part1Initialise();
            Console.WriteLine(testInitialiser.GetMemorySum());

            List<string> puzzleInput = Retriever.Retrieve(@"..//..//..//Day14PuzzleInput.txt");

            List<ICommand> initProgram = ConvertToInitProgram(puzzleInput);

            var initialiser = new Initialiser(initProgram);
            initialiser.Part1Initialise();
            Console.WriteLine(initialiser.GetMemorySum());
        }

        private static List<ICommand> ConvertToInitProgram(List<string> input)
        {
            var initProgram = new List<ICommand>();

            foreach (string line in input)
            {
                string[] equation = line.Split(" ");

                if (line.StartsWith("mask"))
                {
                    var bitMask = new BitMask(equation[2]);
                    initProgram.Add(bitMask);
                }

                if (line.StartsWith("mem"))
                {
                    int address = GetMemoryAddress(line);
                    var memory = new MemoryCommand(address, int.Parse(equation[2]));
                    initProgram.Add(memory);
                }
            }

            return initProgram;
        }

        private static int GetMemoryAddress(string memCommand)
        {
            int addressStart = memCommand.IndexOf('[') + 1;
            int countToAddressEnd = memCommand.IndexOf(']') - addressStart;

            return int.Parse(memCommand.Substring(addressStart, countToAddressEnd));
        }
    }
}
