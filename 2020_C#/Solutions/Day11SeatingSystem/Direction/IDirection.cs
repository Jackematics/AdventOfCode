using System;
using System.Collections.Generic;
using System.Text;

namespace Day11SeatingSystem.Direction
{
    interface IDirection
    {
        public bool AdjacentOccupiedSeatExists(Seat seat, Seat[,] seatGrid);
    }
}
