using System;
using System.Collections.Generic;
using System.Text;

using System.Drawing;

namespace Day12RainRisk.Actions
{
    public class West : IInstruction, IDirection
    {
        public int Value { get; }

        public West(int value)
        {
            Value = value;
        }

        public Point Execute()
        {
            return new Point(-Value, 0);
        }
    }
}
