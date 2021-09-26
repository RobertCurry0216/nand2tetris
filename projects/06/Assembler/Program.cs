using System;
using System.Linq;

namespace Assembler
{
    class Program
    {
        static void Main(string[] args)
        {
            if (args.Length < 2)
            {
                Console.WriteLine("Error: no file provided");
                return;
            }
            var fileIn = args.ElementAt(1);
            var fileOut = args.ElementAtOrDefault(2) ?? "out.hack";

            var assembler = new Assembler(fileIn);
            assembler.Assemble();
            assembler.Save(fileOut);
        }
    }
}
