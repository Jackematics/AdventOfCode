using System;
using System.Collections.Generic;
using System.IO;

using PuzzleInputRetriever;

namespace Day9EncodingError
{
    class Program
    {
        static void Main(string[] args)
        {
            var exampleSet = new List<double>()
            {
                35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576
            };

            var testDecryptor = new Decryptor(exampleSet, 5);
            Console.WriteLine(testDecryptor.GetEncryptionWeakness());

            var decryptor = new Decryptor(GetPuzzleInput(), 25);

            Console.WriteLine(decryptor.GetEncryptionWeakness());
        }

        private static List<double> GetPuzzleInput()
        {
            var puzzleInput = new List<double>();

            using (var reader = new StreamReader(@"..//..//..//Day9PuzzleInput.txt"))
            {
                string line;

                while ((line = reader.ReadLine()) != null)
                {
                    puzzleInput.Add(double.Parse(line));
                }
            }

            return puzzleInput;
        }
    }
}
