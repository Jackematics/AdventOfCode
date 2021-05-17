using System;
using System.Collections.Generic;
using System.Text;

namespace Day8HandheldHalting
{
    class NoOperation : IInstruction
    {
        public int Argument { get ; set ; }
        public int IndexMoveAmount { get; set; } = 1;

        public NoOperation(int argument)
        {
            Argument = argument;
        }
    }
}
