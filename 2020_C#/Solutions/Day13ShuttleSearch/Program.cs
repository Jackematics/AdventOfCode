using System;
using System.Collections.Generic;
using PuzzleInputRetriever;

namespace Day13ShuttleSearch
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> puzzleInput = Retriever.Retrieve(@"..//..//..//Day13PuzzleInput.txt");

            // Part 1

            int earliestTimestamp = int.Parse(puzzleInput[0]);
            List<int> busIds = GetPart1BusIds(puzzleInput[1]);

            var busTimetable = new BusTimetable(earliestTimestamp, busIds);

            Console.WriteLine(busTimetable.GetPart1Result());

            // Part 2

            List<Bus> testBuses1 = GetPart2Buses("17,x,13,19");
            var test1 = new BusContest(testBuses1);
            Console.WriteLine(test1.GetPart2ResultEfficient());

            List<Bus> testBuses2 = GetPart2Buses("67,7,59,61");
            var test2 = new BusContest(testBuses2);
            Console.WriteLine(test2.GetPart2ResultEfficient());

            List<Bus> testBuses3 = GetPart2Buses("67,x,7,59,61");
            var test3 = new BusContest(testBuses3);
            Console.WriteLine(test3.GetPart2ResultEfficient());

            List<Bus> testBuses4 = GetPart2Buses("67,7,x,59,61");
            var test4 = new BusContest(testBuses4);
            Console.WriteLine(test4.GetPart2ResultEfficient());

            List<Bus> testBuses5 = GetPart2Buses("1789,37,47,1889");
            var test5 = new BusContest(testBuses5);
            Console.WriteLine(test5.GetPart2ResultEfficient());


            List<Bus> buses = GetPart2Buses(puzzleInput[1]);
            var busContest = new BusContest(buses);

            Console.WriteLine(busContest.GetPart2ResultEfficient());
        }

        private static List<int> GetPart1BusIds(string input)
        {
            var inServiceBusIds = new List<int>();
            string[] allBusIds = input.Split(',');

            foreach (string busId in allBusIds)
            {
                if (busId != "x")
                {
                    inServiceBusIds.Add(int.Parse(busId));
                }
            }

            return inServiceBusIds;
        }

        private static List<Bus> GetPart2Buses(string input)
        {
            var buses = new List<Bus>();
            double timestampOffset = 0;
            string[] allBusIds = input.Split(',');

            foreach (string busId in allBusIds)
            {
                if (busId != "x")
                {
                    var bus = new Bus(double.Parse(busId), timestampOffset);
                    buses.Add(bus);
                }

                timestampOffset++;
            }

            return buses;
        }
    }
}
