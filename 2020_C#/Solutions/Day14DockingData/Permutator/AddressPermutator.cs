using System;
using System.Collections.Generic;
using System.Text;

namespace Day14DockingData
{
    public class AddressPermutator
    {
        private List<Permutation> Permutations { get; set; }
        private string BinaryAddress { get; set; }
        private int CurrentPermutationNumber { get; set; } = 0;


        public AddressPermutator(string binaryAddress)
        {
            BinaryAddress = binaryAddress;
        }

        public List<string> GetAddressPermutations()
        {
            return null;
        }

        public void PermutateAddresses(string address)
        {
            if (Permutations.Count == 0)
            {
                CreateFirstPermutation(address);
            }

            if (BinaryAddress[currentIndex] == '0' || BinaryAddress[currentIndex] == '1')
            {
                Permutations[CurrentPermutationNumber].BinaryAddress += BinaryAddress[currentIndex];
                currentIndex++;
            }
            else if (BinaryAddress[currentIndex] == 'X')
            {
                var currentBranch = Permutations[CurrentPermutationNumber];
                var newBranch = new Permutation(CurrentPermutationNumber + 1, currentBranch.BinaryAddress += '1');
                currentIndex++;
                PermutateAddresses(newBranch, currentIndex);
                currentBranch.BinaryAddress += '0';

            }
        }

        private void CreateFirstPermutation(string address)
        {
            var permutation = new Permutation(CurrentPermutationNumber, "");
        }
    }
}
