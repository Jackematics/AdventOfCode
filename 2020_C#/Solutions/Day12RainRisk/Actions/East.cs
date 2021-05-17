using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

namespace Day12RainRisk.Actions
{
    public class East : IInstruction, IDirection
    {
        public int Value { get; }

        public East(int value)
        {
            Value = value;
        }

        public Point Execute()
        {
            return new Point(Value, 0);
        }
    }
}
