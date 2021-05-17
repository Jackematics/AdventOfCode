using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;


namespace Day12RainRisk.Actions
{
    public class South : IInstruction, IDirection
    {
        public int Value { get; }

        public South(int value)
        {
            Value = value;
        }

        public Point Execute()
        {
            return new Point(0, -Value);
        }
    }
}
