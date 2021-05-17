using System;
using System.Collections.Generic;
using System.Text;

namespace Day14DockingData
{
    public class BitMask : ICommand
    {
        public string BinaryStringValue { get; set; }

        public BitMask(string bitValue)
        {
            BinaryStringValue = bitValue;
        }
    }
}
