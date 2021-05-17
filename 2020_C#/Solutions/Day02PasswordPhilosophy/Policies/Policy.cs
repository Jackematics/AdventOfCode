using System;
using System.Collections.Generic;
using System.Text;

namespace Day_2_Password_Philosophy.Policies
{
    public abstract class Policy
    {
        public virtual int CountValidPasswords(List<string> passwords)
        {
            int count = 0;

            for (int i = 0; i < passwords.Count; i++)
            {
                if (ValidatePassword(passwords[i]))
                {
                    count++;
                }
            }

            return count;
        }

        public abstract bool ValidatePassword(string password);
    }
}
