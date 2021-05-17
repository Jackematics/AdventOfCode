using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Day14DockingData
{
    public class Initialiser
    {
        private List<ICommand> InitProgram { get; }
        private Dictionary<int, double> Memory { get; set; } = new Dictionary<int, double>();
        private string CurrentBitMask { get; set; }

        public Initialiser(List<ICommand> initProgram)
        {
            InitProgram = initProgram;
        }

        public double GetMemorySum()
        {
            double sum = 0;

            foreach (int key in Memory.Keys)
            {
                sum += Memory[key];
            }

            return sum;
        }

        public void Part1Initialise()
        {
            if (Memory != null)
            {
                Memory.Clear();
            }

            foreach (ICommand command in InitProgram)
            {
                if (command is BitMask)
                {
                    var bitMaskCommand = command as BitMask;
                    CurrentBitMask = bitMaskCommand.BinaryStringValue;
                }

                if (command is MemoryCommand)
                {
                    var memoryCommand = command as MemoryCommand;

                    int address = memoryCommand.Address;

                    string maskifiedBinaryString = MaskifyValue(memoryCommand.BinaryStringValue);
                    double value = IntifyBinaryStringValue(maskifiedBinaryString);

                    if (Memory.ContainsKey(address))
                    {
                        Memory[address] = value;
                    }
                    else
                    {
                        Memory.Add(memoryCommand.Address, value);
                    }
                }
            }
        }

        private string MaskifyValue(string binaryStringValue, int version = 1)
        {
            var result = new StringBuilder();

            for (int i = 0; i < CurrentBitMask.Length; i++)
            {
                if (version == 1)
                {
                    char bit = CurrentBitMask[i] == 'X' ? binaryStringValue[i] : CurrentBitMask[i];
                    result.Append(bit);
                }
                else
                {
                    char bit = CurrentBitMask[i] == '0' ? binaryStringValue[i] : CurrentBitMask[i];
                    result.Append(bit);
                }
            }

            return result.ToString();
        }

        private double IntifyBinaryStringValue(string binaryStringValue)
        {
            double sum = 0;
            for (int i = binaryStringValue.Length - 1; i >= 0; i--)
            {
                if (binaryStringValue[i] == '1')
                {
                    sum += Math.Pow(2, (double)(binaryStringValue.Length - 1 - i));
                }
            }

            return sum;
        }

        public void Part2Initialise()
        {
            if (Memory != null)
            {
                Memory.Clear();
            }

            foreach (ICommand command in InitProgram)
            {
                if (command is BitMask)
                {
                    var bitMaskCommand = command as BitMask;
                    CurrentBitMask = bitMaskCommand.BinaryStringValue;
                }

                if (command is MemoryCommand)
                {
                    var memoryCommand = command as MemoryCommand;

                    int address = memoryCommand.Address;
                    List<int> maskifiedAddresses = GetMaskifiedAddresses(address);

                    string maskifiedBinaryString = MaskifyValue(memoryCommand.BinaryStringValue);
                    double value = IntifyBinaryStringValue(maskifiedBinaryString);

                    if (Memory.ContainsKey(address))
                    {
                        Memory[address] = value;
                    }
                    else
                    {
                        Memory.Add(memoryCommand.Address, value);
                    }
                }
            }
        }

        private List<int> GetMaskifiedAddresses(int address)
        {
            string addressAsBinaryString = Convert.ToString(address, 2).PadLeft(36, '0');
            string maskifiedAddress = MaskifyValue(addressAsBinaryString, 2);
            List<string> defloatedMaskifiedAddresses = DefloatAddress(maskifiedAddress);

        }

        private List<string> DefloatAddress(string maskifiedAddress, int index = 0)
        {
            if ()
        }
    }
}
