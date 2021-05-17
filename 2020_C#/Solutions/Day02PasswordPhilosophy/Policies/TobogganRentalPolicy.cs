using System;
using System.Collections.Generic;
using System.Text;

namespace Day_2_Password_Philosophy.Policies
{
    class TobogganRentalPolicy : Policy
    {
        public override bool ValidatePassword(string password)
        {
            string[] passwordData = password.Split(" ");
            string[] bounds = passwordData[0].Split("-");
            int lowerBound = int.Parse(bounds[0]);
            int upperBound = int.Parse(bounds[1]);

            char key = passwordData[1][0];
            string cipher = passwordData[2];

            int count = 0;
            if (cipher[lowerBound - 1] == key)
            {
                count++;
            }
            if (cipher[upperBound - 1] == key)
            {
                count++;
            }

            return count == 1;
        }
    }
}
