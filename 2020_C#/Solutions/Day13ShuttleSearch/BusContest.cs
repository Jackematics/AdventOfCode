using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Day13ShuttleSearch
{
    public class BusContest
    {
        private List<Bus> Buses { get; }

        public BusContest(List<Bus> buses)
        {
            Buses = buses;
        }

        public double GetPart2ResultEfficient()
        {
            bool busTimesNotInSequence = true;

            double largestBusId = Buses.Max(bus => bus.BusId);
            double largestBusOffset = Buses.Where(bus => bus.BusId == largestBusId).FirstOrDefault().TimestampOffset;

            double currentTimestamp = 0;
            double offsetTimestamp = 0;

            while (busTimesNotInSequence)
            {
                offsetTimestamp += largestBusId;

                if (TimestampValid(offsetTimestamp, largestBusOffset))
                {
                    currentTimestamp = offsetTimestamp - largestBusOffset;

                    for (int i = 1; i < Buses.Count; i++)
                    {
                        if (!IsValid(currentTimestamp, Buses[i]))
                        {
                            break;
                        }

                        if (i == Buses.Count - 1)
                        {
                            busTimesNotInSequence = false;
                        }
                    }
                }                
            }

            return currentTimestamp;
        }

        private bool TimestampValid(double offsetTimestamp, double largestBusOffset) => (offsetTimestamp - largestBusOffset) % Buses[0].BusId == 0; 
        private bool IsValid(double currentTimestamp, Bus bus) => (currentTimestamp + bus.TimestampOffset) % bus.BusId == 0;

    }
}
