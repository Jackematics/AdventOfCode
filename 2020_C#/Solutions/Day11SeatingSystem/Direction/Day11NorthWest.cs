using System;
using System.Collections.Generic;
using System.Text;

namespace Day11SeatingSystem.Direction
{
    public class Day11NorthWest : Direction, IDirection
    {
        public bool AdjacentOccupiedSeatExists(Seat seat, Seat[,] seatGrid)
        {
            for (int row = seat.Position.Y - 1, column = seat.Position.X - 1; row >= 0 || column >= 0; row--, column--)
            {
                if (SeatExistsInGrid(seatGrid, row, column))
                {
                    if (seatGrid[row, column].CurrentSeatType == 'L')
                    {
                        return false;
                    }

                    if (seatGrid[row, column].CurrentSeatType == '#')
                    {
                        return true;
                    }
                }
            }

            return false;
        }
    }
}
