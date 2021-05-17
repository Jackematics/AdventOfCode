using System;
using System.Collections.Generic;
using System.Text;

namespace Day_2_Password_Philosophy.Policies
{
    public class SledRentalPolicy : Policy
    {
        public override bool ValidatePassword(string password)
        {
            string[] passwordData = password.Split(" ");
            string[] bounds = passwordData[0].Split("-");
            int lowerBound = int.Parse(bounds[0]);
            int upperBound = int.Parse(bounds[1]);

            char key = passwordData[1][0];
            string cipher = passwordData[2];
            int keyOccurrences = CountOccurrences(key, cipher);

            return lowerBound <= keyOccurrences && keyOccurrences <= upperBound;
        }

        private int CountOccurrences(char letter, string cipher)
        {
            int count = 0;
            for (int i = 0; i < cipher.Length; i++)
            {
                if (cipher[i] == letter)
                {
                    count++;
                }
            }
            return count;
        }
    }
}
