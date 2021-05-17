using System;
using System.Collections.Generic;
using System.IO;
using System.Net;

using PuzzleInputRetriever;

namespace Day1
{
    class Program
    {
        static void Main(string[] args)
        {
            List<int> expenseReport = GetExpenseReport();
            int part1Result = CalculatePart1Result(expenseReport);

            Console.WriteLine("Part 1 result:" + part1Result);

            int part2Result = CalculatePart2Result(expenseReport);

            Console.WriteLine("Part 2 result:" + part2Result);
        }

        private static List<int> GetExpenseReport()
        {
            var expenses = new List<int>();

            using (StreamReader reader = new StreamReader(@"..\\..\\..\\Day1PuzzleInput.txt"))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    expenses.Add(int.Parse(line));
                }
            }

            return expenses;
        }

        private static int CalculatePart1Result(List<int> expenseReport)
        {
            for (int i = 0; i < expenseReport.Count; i++)
            {
                for (int j = 0; j < expenseReport.Count; j++)
                {
                    if ((expenseReport[i] + expenseReport[j]) == 2020)
                    {
                        return expenseReport[i] * expenseReport[j];
                    }
                }
            }

            throw new ArgumentException("Expense report must contain two numbers that sum to 2020");
        }

        private static int CalculatePart2Result(List<int> expenseReport)
        {
            for (int i = 0; i < expenseReport.Count; i++)
            {
                for (int j = 0; j < expenseReport.Count; j++)
                {
                    for (int k = 0; k < expenseReport.Count; k++)
                    {
                        if ((expenseReport[i] + expenseReport[j] + expenseReport[k]) == 2020)
                        {
                            return expenseReport[i] * expenseReport[j] * expenseReport[k];
                        }
                    }                    
                }
            }

            throw new ArgumentException("Expense report must contain two numbers that sum to 2020");
        }
    }
}
