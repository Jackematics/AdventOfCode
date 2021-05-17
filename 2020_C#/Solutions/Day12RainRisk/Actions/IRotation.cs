using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

namespace Day12RainRisk.Actions
{
    public interface IRotation
    {
        public ShipDirection.Direction Part1Execute(ShipDirection.Direction currentDirection);
        public Point Part2Execute(Point currentDirection);
    }
}
