using System;
using System.Collections.Generic;
using System.IO;

namespace PuzzleInputRetriever
{
    public static class Retriever
    {
        public static List<string> Retrieve(string path)
        {
            {
                var puzzleInput = new List<string>();

                using (var reader = new StreamReader(path))
                {
                    string line;

                    while ((line = reader.ReadLine()) != null)
                    {
                        puzzleInput.Add(line);
                    }
                }

                return puzzleInput;
            }
        }
    }
}
