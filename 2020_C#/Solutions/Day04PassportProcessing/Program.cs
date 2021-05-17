using System;
using System.Collections.Generic;
using System.IO;

using PuzzleInputRetriever;

namespace Day_4_Passport_Processing
{
    class Program
    {
        private static readonly List<TravelDocument> _TravelDocuments = PassportifyPuzzleInput(GetPuzzleInput());

        static void Main(string[] args)
        {
            int part1Solution = NumberOfValidPassports(TravelDocument.ValidationType.Normal);
            Console.WriteLine("Number of valid passports: {0}", part1Solution);

            int part2Solution = NumberOfValidPassports(TravelDocument.ValidationType.Strict);
            Console.WriteLine("Number of valid passports: {0}", part2Solution);
        }

        private static List<string> GetPuzzleInput()
        {
            var puzzleInput = new List<string>();

            using (var reader = new StreamReader(@"..//..//..//Day4PuzzleInput.txt"))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    puzzleInput.Add(line);
                }
            }

            return puzzleInput;
        }

        private static List<TravelDocument> PassportifyPuzzleInput(List<string> puzzleInput)
        {
            var travelDocuments = new List<TravelDocument>();
            var currentDocument = new TravelDocument();

            for (int i = 0; i < puzzleInput.Count; i++)
            {
                if (puzzleInput[i] == "")
                {
                    travelDocuments.Add(currentDocument);
                    currentDocument = new TravelDocument();
                }
                else
                {
                    string[] travelDetails = puzzleInput[i].Split(" ");
                    for (int j = 0; j < travelDetails.Length; j++)
                    {
                        currentDocument.SetProperty(travelDetails[j]); 
                    }
                }
            }
            travelDocuments.Add(currentDocument);

            return travelDocuments;
        }

        private static int NumberOfValidPassports(TravelDocument.ValidationType validationType)
        {
            int count = 0;
            foreach (TravelDocument travelDocument in _TravelDocuments)
            {
                if (travelDocument.Validate(validationType))
                {
                    count++;
                }
            }

            return count;
        }
    }
}
