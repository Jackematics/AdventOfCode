using System;
using System.Collections.Generic;
using System.Text;

using System.Drawing;

namespace Day12RainRisk.Actions
{
    public class North : IInstruction, IDirection
    {
        public int Value { get; }

        public North(int value)
        {
            Value = value;
        }

        public Point Execute()
        {
            return new Point(0, Value);
        }
    }
}
