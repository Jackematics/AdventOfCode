using System;
using System.Collections.Generic;
using System.Text;

namespace Day9EncodingError
{
    class Decryptor
    {
        private List<double> PuzzleInput { get; }
        private List<double> Preamble { get; set; } = new List<double>();
        private int CurrentIndex { get; set; }

        public Decryptor(List<double> puzzleInput, int preambleLength)
        {
            PuzzleInput = puzzleInput;
            Preamble = puzzleInput.GetRange(0, preambleLength);
            CurrentIndex = preambleLength;
        }

        public double GetEncryptionWeakness()
        {
            List<double> contiguousSum = GetContiguousSum();
            contiguousSum.Sort();

            return contiguousSum[0] + contiguousSum[contiguousSum.Count - 1];
        }

        private List<double> GetContiguousSum()
        {
            double invalidNumber = SumNotInPreamble();
            
            for (int i = 0; i < CurrentIndex; i++)
            {
                double setSum = PuzzleInput[i];
                var contiguousSet = new List<double>() { setSum };

                for (int j = i + 1; j < CurrentIndex; j++)
                {
                    double nextInSet = PuzzleInput[j];
                    contiguousSet.Add(nextInSet);
                    setSum += nextInSet;

                    if (setSum == invalidNumber)
                    {
                        return contiguousSet;
                    }                    
                }
            }

            throw new ArgumentException("A contiguous sum must exist");
        }

        private double SumNotInPreamble()
        {
            while(CurrentIndex <= PuzzleInput.Count)
            {
                double currentItem = PuzzleInput[CurrentIndex];
                if (IsValid(currentItem))
                {
                    Preamble.RemoveAt(0);
                    Preamble.Add(currentItem);
                    CurrentIndex++;                    
                }
                else
                {
                    return currentItem;
                }
            }

            throw new ArgumentException("There must exist a valid sum");
        }

        private bool IsValid(double currentItem)
        {
            for (int i = 0; i < Preamble.Count; i++)
            {
                for (int j = 0; j < Preamble.Count; j++)
                {
                    if (i == j)
                    {
                        continue;
                    }

                    if (Preamble[i] + Preamble[j] == currentItem)
                    {
                        return true;
                    }
                }
            }

            return false;
        }
    }
}
