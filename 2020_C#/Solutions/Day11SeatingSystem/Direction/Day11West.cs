using System;
using System.Collections.Generic;
using System.Text;

namespace Day11SeatingSystem.Direction
{
    public class Day11West : Direction, IDirection
    {
        public bool AdjacentOccupiedSeatExists(Seat seat, Seat[,] seatGrid)
        {
            for (int column = seat.Position.X - 1; column >= 0; column--)
            {
                if (SeatExistsInGrid(seatGrid, seat.Position.Y, column))
                {
                    if (seatGrid[seat.Position.Y, column].CurrentSeatType == 'L')
                    {
                        return false;
                    }

                    if (seatGrid[seat.Position.Y, column].CurrentSeatType == '#')
                    {
                        return true;
                    }
                }
            }

            return false;
        }
    }
}
