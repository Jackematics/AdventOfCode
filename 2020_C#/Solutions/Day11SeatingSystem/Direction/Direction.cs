using System;
using System.Collections.Generic;
using System.Text;

namespace Day11SeatingSystem.Direction
{
    public class Direction
    {
        protected bool SeatExistsInGrid(Seat[,] seatGrid, int column, int row)
        {
            return  column >= 0 &&
                    column < seatGrid.GetLongLength(0) &&
                    row >= 0 &&
                    row < seatGrid.GetLongLength(1);
        }
    }
}
