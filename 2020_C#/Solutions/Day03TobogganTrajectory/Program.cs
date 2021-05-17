using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;

using PuzzleInputRetriever;

namespace Day_3_Toboggan_Trajectory
{
    class Program
    {
        private static char[,] _Slope = Gridify(GetPuzzleInput());
        static void Main(string[] args)
        {
            var movements = new List<Point>()
            {
                new Point(1, 1),
                new Point(3, 1),
                new Point(5, 1),
                new Point(7, 1),
                new Point(1, 2)
            };

            var treeEncounters = new List<int>();

            foreach (Point movement in movements)
            {
                treeEncounters.Add(NumberOfTreeEncounters(_Slope, movement));
            }

            int result = 1;
            foreach (int treeEncounterSet in treeEncounters)
            {
                result *= treeEncounterSet;
            }

            Console.WriteLine(result);
        }

        private static char[,] Gridify(List<string> input)
        {
            int gridRowLength = input[0].Length;
            int gridColumnLength = input.Count;            

            char[,] grid = new char[gridColumnLength, gridRowLength];

            // Rows
            for (int i = 0; i < gridColumnLength; i++)
            {        
                // Columns
                for (int j = 0; j < gridRowLength; j++)
                {
                    grid[i, j] = input[i][j];
                }
            }

            return grid;
        }

        private static List<string> GetPuzzleInput()
        {
            var rows = new List<string>();

            using (var reader = new StreamReader(@"../../../Day3PuzzleInput.txt"))
            {
                string line;
                
                while ((line = reader.ReadLine()) != null)
                {
                    rows.Add(line);
                }
            }

            return rows;
        }

        private static int NumberOfTreeEncounters(char[,] slope, Point movement)
        {
            const char tree = '#';

            Point tobogganPosition = new Point(0, 0);
            int treesEncountered = 0;

            while (tobogganPosition.Y < slope.GetLength(0))
            {
                if (slope[tobogganPosition.Y, tobogganPosition.X] == tree)
                {
                    treesEncountered++;
                }

                tobogganPosition.X = (tobogganPosition.X + movement.X) % (slope.GetLength(1));
                tobogganPosition.Y += movement.Y;
            }

            return treesEncountered;
        }
    }
}
