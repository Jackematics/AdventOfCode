using System;
using System.Collections.Generic;
using System.Text;

namespace Day13ShuttleSearch
{
    public class Bus
    {
        public double BusId { get; }
        public double TimestampOffset { get; set; }

        public Bus(double busId, double timestampOffset)
        {
            BusId = busId;
            TimestampOffset = timestampOffset;
        }
    }
}
