using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using System.Linq;

using PuzzleInputRetriever;

namespace Day6CustomCustoms
{
    class Program
    {
        private static readonly List<string> _Part1Answers = GetPart1InputTransform();
        private static readonly List<string> _Part2Answers = GetPart2InputTransform();

        static void Main(string[] args)
        {
            Console.WriteLine(GetYesCountSum(_Part1Answers));
            Console.WriteLine(GetYesCountSum(_Part2Answers));
        }

        private static List<string> GetPart1InputTransform()
        {
            var answers = new List<string>();
            var answerBuilder = new StringBuilder();

            using (var reader = new StreamReader(@"..//..//..//Day6PuzzleInput.txt"))
            {
                string line;

                while ((line = reader.ReadLine()) != null)
                {
                    if (line != "")
                        answerBuilder.Append(line);
                    else
                    {
                        answers.Add(answerBuilder.ToString());
                        answerBuilder = new StringBuilder();
                    }
                }
                answers.Add(answerBuilder.ToString());
            }

            return answers;
        }

        private static int GetYesCountSum(List<string> answers)
        {
            int sum = 0;
            foreach (string answer in answers)
                sum += GetYesCount(answer);

            return sum;
        }

        private static int GetYesCount(string answer)
        {
            return answer.Distinct().Count();
        }

        private static List<string> GetPart2InputTransform()
        {
            var answers = new List<string>();
            var group = new List<string>();

            using (var reader = new StreamReader(@"..//..//..//Day6PuzzleInput.txt"))
            {
                string line;

                while ((line = reader.ReadLine()) != null)
                {
                    if (line != "")
                        group.Add(line);
                    else
                    {
                        answers.Add(GetGroupYesses(group));
                        group = new List<string>();
                    }                        
                }
                answers.Add(GetGroupYesses(group));
            }

            return answers;
        }

        private static string GetGroupYesses(List<string> group)
        {
            var answers = new StringBuilder();
            string largestAnswerSet = group.Aggregate("", (max, current) => max.Length > current.Length ? max : current);

            foreach (char answer in largestAnswerSet)
            {
                int count = 0;
                foreach (string answerSet in group)
                {
                    if (answerSet.Contains(answer))
                        count++;
                }

                if (count == group.Count)
                    answers.Append(answer);                    
            }

            return answers.ToString();
        }
    }    
}
