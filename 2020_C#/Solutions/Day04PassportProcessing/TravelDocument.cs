using System;
using System.Collections.Generic;
using System.Text;
using System.Text.RegularExpressions;

namespace Day_4_Passport_Processing
{
    class TravelDocument
    {
        public string BirthYear { get; set; }
        public string IssueYear { get; set; }
        public string ExpirationYear { get; set; }
        public string Height { get; set; }
        public string HairColour { get; set; }
        public string EyeColour { get; set; }
        public string PassportID { get; set; }
        public string CountryID { get; set; }

        public void SetProperty(string property)
        {
            const int dataStartIndex = 4;

            string dataType = property.Substring(0, 3);
            string data = property.Substring(dataStartIndex);

            switch (dataType)
            {
                case "byr":
                    BirthYear = data;
                    break;

                case "iyr":
                    IssueYear = data;
                    break;

                case "eyr":
                    ExpirationYear = data;
                    break;

                case "hgt":
                    Height = data;
                    break;

                case "hcl":
                    HairColour = data;
                    break;

                case "ecl":
                    EyeColour = data;
                    break;

                case "pid":
                    PassportID = data;
                    break;

                case "cid":
                    CountryID = data;
                    break;
            }
        }

        public enum ValidationType
        {
            Normal,
            Strict
        }

        public bool Validate(ValidationType validationType)
        {
            switch (validationType)
            {
                case ValidationType.Normal:
                    return BirthYear != null &&
                           IssueYear != null &&
                           ExpirationYear != null &&
                           Height != null &&
                           HairColour != null &&
                           EyeColour != null &&
                           PassportID != null;

                case ValidationType.Strict:
                    return StrictValidateBirthYear() &&
                           StrictValidateIssueYear() &&
                           StrictValidateExpirationYear() &&
                           StrictValidateHeight() &&
                           StrictValidateHeight() &&
                           StrictValidateHairColour() &&
                           StrictValidateEyeColour() &&
                           StrictValidatePassportID();

                default:
                    throw new ArgumentException("Must be a valid validation type");
            }            
        }

        private bool StrictValidateBirthYear()
        {
            return BirthYear != null &&
                   BirthYear.Length == 4 &&
                   int.Parse(BirthYear) >= 1920 &&
                   int.Parse(BirthYear) <= 2002;
        }

        private bool StrictValidateIssueYear()
        {
            return IssueYear != null &&
                   IssueYear.Length == 4 &&
                   int.Parse(IssueYear) >= 2010 &&
                   int.Parse(IssueYear) <= 2020;
        }

        private bool StrictValidateExpirationYear()
        {
            return ExpirationYear != null &&
                   ExpirationYear.Length == 4 &&
                   int.Parse(ExpirationYear) >= 2020 &&
                   int.Parse(ExpirationYear) <= 2030;
        }

        private bool StrictValidateHeight()
        {
            if (Height == null)
            {
                return false;
            }

            string measurementType = Height.Substring(Height.Length - 2);

            if (measurementType == "cm" ||
                measurementType == "in")
            {
                string measurement = Height.Substring(0, Height.Length - 2);

                switch (measurementType)
                {
                    case "cm":
                        return int.Parse(measurement) >= 150 &&
                               int.Parse(measurement) <= 193;

                    case "in":
                        return int.Parse(measurement) >= 59 &&
                               int.Parse(measurement) <= 76;

                    default:
                        throw new ArgumentException("Measurement type must be cm or in");
                }
            }

            return false;            
        }

        private bool StrictValidateHairColour()
        {
            return  HairColour != null &&
                    Regex.IsMatch(HairColour, @"^#[a-f0-9]{6}$");                   
        }

        private bool StrictValidateEyeColour()
        {
            return EyeColour != null &&
                   (
                       EyeColour == "amb" ||
                       EyeColour == "blu" ||
                       EyeColour == "brn" ||
                       EyeColour == "gry" ||
                       EyeColour == "grn" ||
                       EyeColour == "hzl" ||
                       EyeColour == "oth"
                   );                   
        }

        private bool StrictValidatePassportID()
        {
            return  PassportID != null &&
                    Regex.IsMatch(PassportID, @"^[0-9]{9}$");
        }
    }
}
