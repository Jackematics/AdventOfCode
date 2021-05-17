using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

using PuzzleInputRetriever;

namespace Day5BinaryBoarding
{
    class Program
    {
        private static readonly List<string> _PuzzleInput = GetPuzzleInput();
        private const int _PlaneRowCount = 128;
        private const int _PlaneColumnCount = 8;
        private const int _RowInputRange = 7;

        static void Main(string[] args)
        {
            double part1Result = GetHighestSeatID();
            Console.WriteLine(part1Result);

            Console.WriteLine(GetMissingSeatId());
        }

        private static List<string> GetPuzzleInput()
        {
            var puzzleInput = new List<string>();

            using (var reader = new StreamReader(@"..//..//..//Day5PuzzleInput.txt"))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    puzzleInput.Add(line);
                }
            }

            return puzzleInput;
        }

        private static int GetMissingSeatId()
        {
            List<int> seatIDs = GetSeatIDs();
            seatIDs.Sort();

            int lowerBound = seatIDs[0];
            int upperBound = seatIDs[seatIDs.Count - 1];

            for (int i = lowerBound; i <= upperBound; i++)
            {
                if (!seatIDs.Contains(i) &&
                    seatIDs.Contains(i - 1) &&
                    seatIDs.Contains(i + 1))
                {
                    return i;
                }
            }

            throw new ArgumentException("Your missing seat exists, you just haven't coded this properly");
        }

        private static int GetHighestSeatID()
        {
            return GetSeatIDs().Max();
        }

        private static List<int> GetSeatIDs()
        {
            var seatIDs = new List<int>();

            foreach (string input in _PuzzleInput)
            {
                seatIDs.Add(GetSeatID(input));
            }

            return seatIDs;
        }

        private static int GetSeatID(string input)
        {
            string rowInput = input.Substring(0, _RowInputRange);
            string columnInput = input.Substring(_RowInputRange);

            int row = (int)GetDimension(Dimension.Row, rowInput);
            int column = (int)GetDimension(Dimension.Column, columnInput);

            return row * _PlaneColumnCount + column;
        }

        private enum Dimension
        {
            Row,
            Column
        }

        private static double GetDimension(Dimension dimension, string rowInput)
        {
            double lowerBound = 0;
            double upperBound = dimension == Dimension.Row ? _PlaneRowCount : _PlaneColumnCount;

            foreach (char character in rowInput)
            {
                if (character == 'F' || character == 'L')
                {
                    upperBound = (upperBound + lowerBound) / 2;
                }

                if (character == 'B' || character == 'R')
                {
                    lowerBound += ((upperBound - lowerBound) / 2);
                }
            }

            return lowerBound;
        }
    }
}
