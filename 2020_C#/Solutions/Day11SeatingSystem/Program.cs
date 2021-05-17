using System;
using System.Collections.Generic;
using System.Drawing;

using PuzzleInputRetriever;

namespace Day11SeatingSystem
{
    class Program
    {
        private static Seat[,] SeatGrid { get; set; }

        static void Main(string[] args)
        {
            var testCase = new List<string>
            {
                "L.LL.LL.LL",
                "LLLLLLL.LL",
                "L.L.L..L..",
                "LLLL.LL.LL",
                "L.LL.LL.LL",
                "L.LLLLL.LL",
                "..L.L.....",
                "LLLLLLLLLL",
                "L.LLLLLL.L",
                "L.LLLLL.LL"
            };

            Seat[,] testGrid = GridifyInput(testCase);

            var roundExecutorTest = new RoundExecutor(testGrid);
            Console.WriteLine(roundExecutorTest.GetEquilibriumOccupiedSeatCount());

            SeatGrid = GridifyInput(Retriever.Retrieve(@"..//..//..//Day11PuzzleInput.txt"));

            Console.WriteLine();

            Console.WriteLine("Part 1 Solution: ");
            var roundExecutor = new RoundExecutor(SeatGrid);

            Console.WriteLine(roundExecutor.GetEquilibriumOccupiedSeatCount());
        }

        public static void DisplayGridToConsole(Seat[,] grid, string header)
        {
            Console.WriteLine(header);

            for (int row = 0; row < grid.GetLongLength(0); row++)
            {
                for (int column = 0; column < grid.GetLongLength(1); column++)
                {
                    Console.Write(grid[row, column].CurrentSeatType + " ");
                }

                Console.WriteLine();
            }

            Console.WriteLine();
        }

        private static Seat[,] GridifyInput(List<string> input)
        {
            int numberOfRows = input.Count;
            int numberOfColumns = input[0].Length;

            var seatGrid = new Seat[numberOfRows, numberOfColumns];

            for (int row = 0; row < numberOfRows; row++)
            {
                for (int column = 0; column < numberOfColumns; column++)
                {
                    var seat = new Seat
                    {
                        CurrentSeatType = input[row][column],
                        Position = new Point(column, row)
                    };

                    seatGrid[row, column] = seat;
                }
            }

            return seatGrid;
        }
    }
}
