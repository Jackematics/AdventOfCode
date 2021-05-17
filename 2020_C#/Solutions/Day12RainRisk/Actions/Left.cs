using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

namespace Day12RainRisk.Actions
{
    public class Left : IInstruction, IRotation
    {
        public int Value { get; }

        public Left(int value)
        {
            Value = value;
        }

        public ShipDirection.Direction Part1Execute(ShipDirection.Direction currentDirection)
        {
            int currentEnumValue = (int)currentDirection;
            int rotationValue = Value / 90;
            int newEnumValue = (currentEnumValue - rotationValue + 4) % 4;

            return (ShipDirection.Direction)newEnumValue;
        }

        public Point Part2Execute(Point currentWaypointPosition)
        {
            var newWaypointPosition = new Point();

            switch (Value)
            {
                case 90:
                    newWaypointPosition.X = -currentWaypointPosition.Y;
                    newWaypointPosition.Y = currentWaypointPosition.X;
                    break;

                case 180:
                    newWaypointPosition.X = -currentWaypointPosition.X;
                    newWaypointPosition.Y = -currentWaypointPosition.Y;
                    break;

                case 270:
                    newWaypointPosition.X = currentWaypointPosition.Y;
                    newWaypointPosition.Y = -currentWaypointPosition.X;
                    break;
            }

            return newWaypointPosition;
        }
    }
}
