using System;
using System.Collections.Generic;
using System.Text;

namespace Day11SeatingSystem.Direction
{
    public class Day11South : Direction, IDirection
    {
        public bool AdjacentOccupiedSeatExists(Seat seat, Seat[,] seatGrid)
        {
            for (int row = seat.Position.Y + 1; row < seatGrid.GetLongLength(0); row++)
            {
                if (SeatExistsInGrid(seatGrid, row, seat.Position.X))
                {
                    if (seatGrid[row, seat.Position.X].CurrentSeatType == 'L')
                    {
                        return false;
                    }

                    if (seatGrid[row, seat.Position.X].CurrentSeatType == '#')
                    {
                        return true;
                    }
                }
            }

            return false;
        }
    }
}
