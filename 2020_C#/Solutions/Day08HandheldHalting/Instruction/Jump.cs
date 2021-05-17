using System;
using System.Collections.Generic;
using System.Text;

namespace Day8HandheldHalting
{
    class Jump : IInstruction
    {
        public int Argument { get ; set ; }
        public int IndexMoveAmount { get; set; } 

        public Jump(int argument)
        {
            IndexMoveAmount = Argument = argument;
        }
    }
}
