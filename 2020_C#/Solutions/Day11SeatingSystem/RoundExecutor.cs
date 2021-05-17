using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

using Day11SeatingSystem.Direction;

namespace Day11SeatingSystem
{
    public class RoundExecutor
    {
        private Seat[,] SeatGrid { get; }
        private List<IDirection> Directions { get; } = new List<IDirection>
        {
            new Day11North(),
            new Day11NorthEast(),
            new Day11East(),
            new Day11SouthEast(),
            new Day11South(),
            new Day11SouthWest(),
            new Day11West(),
            new Day11NorthWest(),
        };

        private const int _SeatTolerance = 5;

        public RoundExecutor(Seat[,] seatGrid)
        {
            SeatGrid = seatGrid;

            ExecuteRounds();

            Program.DisplayGridToConsole(SeatGrid, "Final Round: ");
        }

        public int GetEquilibriumOccupiedSeatCount()
        {
            int occupiedSeatCount = 0;

            for (int row = 0; row < SeatGrid.GetLongLength(0); row++)
            {
                for (int column = 0; column < SeatGrid.GetLongLength(1); column++)
                {
                    if (SeatGrid[row, column].CurrentSeatType == '#')
                    {
                        occupiedSeatCount++;
                    }
                }
            }

            return occupiedSeatCount;
        }

        private void ExecuteRounds()
        {
            int count = 1;
            while(true)
            {
                Program.DisplayGridToConsole(SeatGrid, "Round" + count + ": ");

                foreach (Seat seat in SeatGrid)
                {
                    seat.NextSeatType = GetNextSeatType(seat);
                }

                if (CurrentAndNextSeatsAreAllEqual())
                {
                    return;
                }

                foreach (Seat seat in SeatGrid)
                {
                    seat.CurrentSeatType = seat.NextSeatType;
                }

                count++;
            }            
        }

        private char GetNextSeatType(Seat seat)
        {
            char nextSeatType = '\0';

            //PART 1 Method
            //int adjacentSeatCount = GetPart1OccupiedAdjacentSeatCount(seat);

            //PART 2 Method
            int adjacentSeatCount = GetPart2OccupiedAdjacentSeatCount(seat);

            switch (seat.CurrentSeatType)
            {
                case 'L':
                    nextSeatType = adjacentSeatCount == 0 ? '#' : 'L';
                    break;

                case '#':
                    nextSeatType = adjacentSeatCount >= _SeatTolerance ? 'L' : '#';
                    break;

                case '.':
                    nextSeatType = '.';
                    break;

                default: throw new ArgumentException("Must be a valid seat type");
            }

            return nextSeatType;
        }

        private bool CurrentAndNextSeatsAreAllEqual()
        {
            foreach (var seat in SeatGrid)
            {
                if (seat.CurrentSeatType != seat.NextSeatType)
                {
                    return false;
                }
            }

            return true;
        }

        private int GetPart1OccupiedAdjacentSeatCount(Seat seat)
        {
            int count = 0;

            for (int row = seat.Position.Y - 1; row <= seat.Position.Y + 1; row++)
            {
                for (int column = seat.Position.X - 1; column <= seat.Position.X + 1; column++)
                {
                    if (SeatExistsInGrid(row, column) && IsNotCentre(seat, row, column))
                    {
                        Seat seatToCheck = SeatGrid[row, column]; 

                        if (seatToCheck.CurrentSeatType == '#')
                        {
                            count++;
                        }
                    }
                }
            }

            return count;
        }

        private bool SeatExistsInGrid(int column, int row)
        {
            return  column >= 0 &&
                    column < SeatGrid.GetLongLength(0) &&
                    row >= 0 &&
                    row < SeatGrid.GetLongLength(1);
        }

        private bool IsNotCentre(Seat seat, int row, int column)
        {
            return row != seat.Position.Y ||
                   column != seat.Position.X;
        }

        private int GetPart2OccupiedAdjacentSeatCount(Seat seat)
        {
            int count = 0;
            foreach(IDirection direction in Directions)
            {
                count = direction.AdjacentOccupiedSeatExists(seat, SeatGrid) ? count + 1 : count;
            }

            return count;
        }
    }
}
