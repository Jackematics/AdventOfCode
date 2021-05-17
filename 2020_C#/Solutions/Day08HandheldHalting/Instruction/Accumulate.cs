using System;
using System.Collections.Generic;
using System.Text;

namespace Day8HandheldHalting
{
    class Accumulate : IInstruction
    {        
        public int Argument { get ; set ; }
        public int IndexMoveAmount { get; set; } = 1;

        public Accumulate(int argument)
        {
            Argument = argument;
        }
    }
}
