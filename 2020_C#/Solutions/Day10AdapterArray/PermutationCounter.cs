using System;
using System.Collections.Generic;
using System.Text;

namespace Day10AdapterArray
{
    public class PermutationCounter
    {
        private List<int> OrderedAdaptors { get; }
        private List<List<int>> ChoiceSet { get; set; } = new List<List<int>>();
        private List<int> ArrangementSetCounts { get; set; } = new List<int>();
        private int ArrangementCount { get; set; } = 1;

        public PermutationCounter(List<int> orderedAdaptors)
        {
            OrderedAdaptors = orderedAdaptors;
            OrderedAdaptors.Insert(0, 0);
            OrderedAdaptors.Add(OrderedAdaptors[OrderedAdaptors.Count - 1] + 3);

            SetChoiceSets();

            foreach(List<int> set in ChoiceSet)
            {
                SetAllBranches(set, set[0]);
                ArrangementSetCounts.Add(ArrangementCount);
                ArrangementCount = 1;
            }
        }

        public double GetArrangementCount()
        {
            double count = 1;
            foreach (int arrangementCount in ArrangementSetCounts)
            {
                count *= Convert.ToDouble(arrangementCount);
            }

            return count;
        }
        
        private void SetChoiceSets()
        {
            List<int> choices = new List<int>();

            for (int i = 0; i < OrderedAdaptors.Count - 1; i++)
            {
                choices.Add(OrderedAdaptors[i]);

                if (OrderedAdaptors[i + 1] == OrderedAdaptors[i] + 3)
                {
                    ChoiceSet.Add(choices);
                    choices = new List<int>();
                }                
            }
        }


        private void SetAllBranches(List<int> currentSet, int adaptor = 0)
        {
            List<int> validJumps = GetValidJumps(adaptor, currentSet);

            if (validJumps.Count == 0)
            {
                return;
            }
            
            ArrangementCount += validJumps.Count - 1;                
            foreach(int jump in validJumps)
            {
                SetAllBranches(currentSet, jump);
            }
        }

        private List<int> GetValidJumps(int adaptor, List<int> set)
        {
            var validJumps = new List<int>();

            for (int i = 1; i <= 3; i++)
            {
                if (set.Contains(adaptor + i))
                {
                    validJumps.Add(adaptor + i);
                }
            }

            return validJumps;
        }
    }
}
