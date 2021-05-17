using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

namespace Day12RainRisk.Actions
{
    public class Forward : IInstruction, IMovement
    {
        public int Value { get; }

        public Forward(int value)
        {
            Value = value;
        }

        public Point Part1Execute(ShipDirection.Direction currentDirection)
        {
            var movement = new Point();

            switch (currentDirection)
            {
                case ShipDirection.Direction.North:
                    movement.X = 0;
                    movement.Y = Value;
                    break;

                case ShipDirection.Direction.East:
                    movement.X = Value;
                    movement.Y = 0;
                    break;

                case ShipDirection.Direction.South:
                    movement.X = 0;
                    movement.Y = -Value;
                    break;

                case ShipDirection.Direction.West:
                    movement.X = -Value;
                    movement.Y = 0;
                    break;
            }

            return movement;
        }

        public Point Part2Execute(Point currentWaypointPosition) => new Point(currentWaypointPosition.X * Value, currentWaypointPosition.Y * Value);
    }
}
