using System;
using System.Collections.Generic;
using System.Text;

namespace Day7HandyHaversacks
{
    public class Bag
    {
        public int Amount { get; set; }
        public string Style { get; set; }
        public string Colour { get; set; }
        public List<Bag> SubBags { get; set; } = new List<Bag>();
    }
}
