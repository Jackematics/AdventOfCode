using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

using PuzzleInputRetriever;

namespace Day7HandyHaversacks
{
    class Program
    {
        private static readonly List<string> _PuzzleInput = GetPuzzleInput();
        private static readonly List<Bag> _Bags = BagifyInput(_PuzzleInput);

        static void Main(string[] args)
        {
            Console.WriteLine(GetShinyGoldBagCount());

            Bag shinyGoldBag = _Bags.Where(x => x.Style == "shiny" && x.Colour == "gold").FirstOrDefault();
            Console.WriteLine(InnerBagCount(shinyGoldBag));
        }

        private static int InnerBagCount(Bag outerBag)
        {
            int count = 0;

            if (outerBag.SubBags.Count == 0)
            {
                return 0;
            }

            foreach (Bag subBag in outerBag.SubBags)
            {
                count += subBag.Amount * (1 + InnerBagCount(_Bags
                    .Where(x => x.Style == subBag.Style && x.Colour == subBag.Colour)
                    .FirstOrDefault()));
            }

            return count;
        }

        private static int GetShinyGoldBagCount()
        {
            int count = 0;
            foreach (Bag bag in _Bags)
            {
                if (CanHoldShinyGoldBag(bag))
                {
                    count++;
                }
            }

            return count;
        }

        private static bool CanHoldShinyGoldBag(Bag outerBag)
        {
            foreach (Bag subBag in outerBag.SubBags)
            {
                if (subBag.Style == "shiny" && subBag.Colour == "gold")
                {
                    return true;
                }

                if
                (
                    CanHoldShinyGoldBag(_Bags
                        .Where(x => x.Style == subBag.Style && x.Colour == subBag.Colour)
                        .FirstOrDefault())
                )
                {
                    return true;
                }
            }

            return false;
        }

        private static List<Bag> BagifyInput(List<string> input)
        {
            var bags = new List<Bag>();

            const int subBagStartIndex = 5;

            foreach (string rule in input)
            {
                string[] data = rule.Split(" ");

                var currentBag = new Bag()
                {
                    Style = data[0],
                    Colour = data[1]
                };

                for (int i = subBagStartIndex; i < data.Length; i++)
                {
                    if (data[i].Contains("bag"))
                    {
                        string amount = data[i - 3];
                        string style = data[i - 2];
                        string colour = data[i - 1];

                        if (style != "no" && colour != "other")
                        {
                            var subBag = new Bag()
                            {
                                Amount = int.Parse(amount),
                                Style = style,
                                Colour = colour,
                            };

                            currentBag.SubBags.Add(subBag);
                        }
                    }
                }

                bags.Add(currentBag);
            }

            return bags;
        }

        private static List<string> GetPuzzleInput()
        {
            var puzzleInput = new List<string>();

            using (var reader = new StreamReader(@"..//..//..//Day7PuzzleInput.txt"))
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
