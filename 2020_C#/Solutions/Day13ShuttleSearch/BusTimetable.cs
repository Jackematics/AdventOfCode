using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Day13ShuttleSearch
{
    public class BusTimetable
    {
        private int EarliestTimestamp { get; }
        private List<int> BusIds { get; }

        public BusTimetable(int earliestTimestamp, List<int> busIds)
        {
            EarliestTimestamp = earliestTimestamp;
            BusIds = busIds;
        }

        public int GetPart1Result()
        {
            List<int> nextBusTimes = GetNextBusTimes();

            int closestTime = nextBusTimes.Min();
            int closestTimeIndex = nextBusTimes.IndexOf(closestTime);

            int nextBusId = BusIds[closestTimeIndex];

            return nextBusId * (closestTime - EarliestTimestamp);
        }

        private List<int> GetNextBusTimes()
        {
            var timesBeyondTimestamp = new List<int>();

            foreach (int busId in BusIds)
            {
                int difference = EarliestTimestamp % busId;
                int nextBusTime = EarliestTimestamp - difference + busId;

                timesBeyondTimestamp.Add(nextBusTime);
            }

            return timesBeyondTimestamp;
        }
    }
}
