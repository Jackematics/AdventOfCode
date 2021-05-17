using System;
using System.Collections.Generic;
using System.IO;

using PuzzleInputRetriever;

using Day_2_Password_Philosophy.Policies;

namespace Day_2_Password_Philosophy
{
    class Program
    {
        static private List<string> _Passwords = GetPasswords();
        
        static void Main(string[] args)
        {
            var part1Policy = new SledRentalPolicy();
            var part2Policy = new TobogganRentalPolicy();

            PrintResult(part1Policy, 1);
            PrintResult(part2Policy, 2);
        }

        private static List<string> GetPasswords()
        {
            var passwords = new List<string>();

            using (var reader = new StreamReader(@"..\\..\\..\\Day2PuzzleInput.txt"))
            {
                string line;

                while ((line = reader.ReadLine()) != null)
                {
                    passwords.Add(line);
                }
            }

            return passwords;
        }

        private static void PrintResult(Policy policy, int puzzlePart)
        {
            int validPasswordCount = policy.CountValidPasswords(_Passwords);
            Console.WriteLine("Part {0} Number of valid passwords: {1}", puzzlePart, validPasswordCount);
        }
    }
}
