using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

namespace Day11SeatingSystem
{
    public class Seat
    {
        public char CurrentSeatType { get; set; }
        public char NextSeatType { get; set; }

        public Point Position { get; set; }
    }
}
