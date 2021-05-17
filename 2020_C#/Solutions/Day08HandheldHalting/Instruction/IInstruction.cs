using System;
using System.Collections.Generic;
using System.Text;

namespace Day8HandheldHalting
{
    public interface IInstruction
    {
        public int Argument { get; set; }
        public int IndexMoveAmount { get; set; }
    }
}
