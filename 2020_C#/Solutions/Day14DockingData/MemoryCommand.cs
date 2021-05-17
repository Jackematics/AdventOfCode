using System;
using System.Collections.Generic;
using System.Text;

namespace Day14DockingData
{
    public class MemoryCommand : ICommand
    {
        public int Address { get; set; }
        public string BinaryStringValue { get; set; }
        
        public MemoryCommand(int address, int value)
        {
            Address = address;
            BinaryStringValue = Convert.ToString(value, 2).PadLeft(36, '0');
        }
    }
}
