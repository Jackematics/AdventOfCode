using System;
using System.Collections.Generic;
using System.Linq;

using PuzzleInputRetriever;

namespace Day10AdapterArray
{
    class Program
    {
        private static List<int> _OrderedJoltageAdaptors = GetOrderedJoltageAdaptors();

        private const int JoltUpperThreshold = 3;


        static void Main(string[] args)
        {
            var example1OrderedAdaptors = new List<int>
            {
                1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19
            };
            int[] test1Distribution = GetJoltDistribution(example1OrderedAdaptors);

            var example2OrderedAdaptors = new List<int>
            {
                1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24,
                25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49
            };
            int[] test2Distribution = GetJoltDistribution(example2OrderedAdaptors);

            int[] joltDistribution = GetJoltDistribution(_OrderedJoltageAdaptors);

            // PART 1 Result

            Console.WriteLine(joltDistribution[0] * joltDistribution[2]);
            Console.WriteLine();

            // Part 2 Result

            var permutationCounterTest1 = new PermutationCounter(example1OrderedAdaptors);
            Console.WriteLine(permutationCounterTest1.GetArrangementCount());

            var permutationCounterTest2 = new PermutationCounter(example2OrderedAdaptors);
            Console.WriteLine(permutationCounterTest2.GetArrangementCount());

            var permutationCounter = new PermutationCounter(_OrderedJoltageAdaptors);
            Console.WriteLine(permutationCounter.GetArrangementCount());
        }

        //PART 2 DONE IN PermutationCounter.cs

        //PART 1

        private static int[] GetJoltDistribution(List<int> orderedAdaptors)
        {
            int currentAdaptor = 0;
            var distributions = new int[JoltUpperThreshold];

            while (currentAdaptor != orderedAdaptors[orderedAdaptors.Count - 1])
            {
                List<int> adaptorsWithinMaxJolts = GetNeighboursWithinMaxJolts(currentAdaptor, orderedAdaptors);
                int closestAdaptor = adaptorsWithinMaxJolts.Min();

                distributions = UpdateDistributions(distributions, closestAdaptor - currentAdaptor);
                currentAdaptor = closestAdaptor;
            }

            distributions[2]++;

            return distributions;
        }

        private static List<int> GetNeighboursWithinMaxJolts(int currentAdaptor, List<int> adaptors)
        {
            var neighboursWithinMaxJolts = adaptors
                    .Where(adaptor =>
                        adaptor - currentAdaptor > 0 &&
                        adaptor - currentAdaptor <= 3);

            return neighboursWithinMaxJolts.ToList();
        }

        private static int[] UpdateDistributions(int[] distributions, int jolt)
        {
            switch(jolt)
            {
                case 1:
                    distributions[0]++;
                    break;

                case 2:
                    distributions[1]++;
                    break;

                case 3:
                    distributions[2]++;
                    break;

                default:
                    throw new ArgumentException("Jolt must be between 1 and 3");
            }

            return distributions;
        }

        private static List<int> GetOrderedJoltageAdaptors()
        {
            List<string> puzzleInput = PuzzleInputRetriever.Retriever.Retrieve(@"..//..//..//Day10PuzzleInput.txt");
            var outputJoltages = new List<int>();

            foreach (string input in puzzleInput)
            {
                outputJoltages.Add(int.Parse(input));
            }

            outputJoltages.Sort();

            return outputJoltages;
        }        
    }
}
