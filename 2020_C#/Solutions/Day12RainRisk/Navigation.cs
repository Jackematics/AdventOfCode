using System;
using System.Collections.Generic;
using System.Drawing;
using System.Text;

using Day12RainRisk.Actions;

namespace Day12RainRisk
{
    public class Navigation
    {
        private List<IInstruction> Instructions { get; } 

        public Navigation(List<IInstruction> instructions)
        {
            Instructions = instructions;
        }

        public int GetPart1ManhattanDistance()
        {
            var shipCoordinates = new Point(0, 0);
            var currentDirection = ShipDirection.Direction.East;

            foreach (IInstruction instruction in Instructions)
            {
                if (instruction is IDirection)
                {
                    IDirection direction = instruction as IDirection;
                    Point newCoordinate = direction.Execute();
                    shipCoordinates.X += newCoordinate.X;
                    shipCoordinates.Y += newCoordinate.Y;
                }

                if (instruction is IMovement)
                {
                    IMovement movement = instruction as IMovement;
                    Point newCoordinate = movement.Part1Execute(currentDirection);
                    shipCoordinates.X += newCoordinate.X;
                    shipCoordinates.Y += newCoordinate.Y;
                }

                if (instruction is IRotation)
                {
                    IRotation rotation = instruction as IRotation;
                    currentDirection = rotation.Part1Execute(currentDirection);
                }
            }

            return Math.Abs(shipCoordinates.X) + Math.Abs(shipCoordinates.Y);
        }

        public int GetPart2ManhattanDistance()
        {
            var shipCoordinates = new Point(0, 0);
            var waypoint = new Point(10, 1);

            foreach(IInstruction instruction in Instructions)
            {
                if (instruction is IDirection)
                {
                    IDirection direction = instruction as IDirection;
                    Point waypointMovement = direction.Execute();

                    waypoint.X += waypointMovement.X;
                    waypoint.Y += waypointMovement.Y;
                }

                if (instruction is IMovement)
                {
                    IMovement movement = instruction as IMovement;
                    Point shipMovement = movement.Part2Execute(waypoint);

                    shipCoordinates.X += shipMovement.X;
                    shipCoordinates.Y += shipMovement.Y;
                }

                if (instruction is IRotation)
                {
                    IRotation rotation = instruction as IRotation;
                    waypoint = rotation.Part2Execute(waypoint);
                }
            }

            return Math.Abs(shipCoordinates.X) + Math.Abs(shipCoordinates.Y);
        }
    }
}
