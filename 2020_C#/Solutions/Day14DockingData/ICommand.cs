using System;
using System.Collections.Generic;
using System.Text;

namespace Day14DockingData
{
    public interface ICommand
    {
        public string BinaryStringValue { get; set; }
    }
}
