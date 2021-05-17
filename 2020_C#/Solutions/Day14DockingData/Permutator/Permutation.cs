using System;
using System.Collections.Generic;
using System.Text;

namespace Day14DockingData
{
    public class Permutation
    {
        public int Number { get; set; }
        public string BinaryAddress { get; set; }

        public Permutation(int number, string binaryAddress)
        {
            Number = number;
            BinaryAddress = binaryAddress;
        }
    }
}
