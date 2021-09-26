using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace Assembler
{
    class Assembler
    {
        List<string> lines;
        List<string> machineCode;
        Dictionary<string, int> symbols;
        int symbolCount;

        // regex patterns
        Regex ACommand = new Regex(@"@(?<value>.+)");
        Regex CCommand = new Regex(@"(?:(?<dest>.+)=)?(?<comp>[a-zA-Z0-9_+\-!| ]+)(?:;(?<jump>.+))?");
        Regex Symbol = new Regex(@"\((?<name>.+)\)");

        // tables
        Dictionary<string, int> DefaultSymbols = new Dictionary<string, int>()
            {
                { "R0", 0 },
                { "R1", 1 },
                { "R2", 2 },
                { "R3", 3 },
                { "R4", 4 },
                { "R5", 5 },
                { "R6", 6 },
                { "R7", 7 },
                { "R8", 8 },
                { "R9", 9 },
                { "R10", 10 },
                { "R11", 11 },
                { "R12", 12 },
                { "R13", 13 },
                { "R14", 14 },
                { "R15", 15 },
                { "SP", 0 },
                { "LCL", 1 },
                { "ARG", 2 },
                { "THIS", 3 },
                { "THAT", 4 },
                { "SCREEN", 16384 },
                { "KBD", 24576 }
            };

        Dictionary<string, string> JumpTable = new Dictionary<string, string>()
        {
            { "", "000" },
            { "JGT", "001" },
            { "JEQ", "010" },
            { "JGE", "011" },
            { "JLT", "100" },
            { "JNE", "101" },
            { "JLE", "110" },
            { "JMP", "111" }
        };

        Dictionary<string, string> DestTable = new Dictionary<string, string>()
        {
            { "", "000" },
            { "M", "001" },
            { "D", "010" },
            { "MD", "011" },
            { "A", "100" },
            { "AM", "101" },
            { "AD", "110" },
            { "AMD", "111" }
        };

        Dictionary<string, string> CompTable = new Dictionary<string, string>()
        {
            { "0", "0000000" },
            { "1", "0111111" },
            { "-1", "0111010" },

            { "D", "0001100" },
            { "A", "0110000" },
            { "M", "1110000" },

            { "!D", "0001101" },
            { "!A", "0110001" },
            { "!M", "1110001" },

            { "-D", "0001111" },
            { "-A", "0110011" },
            { "-M", "1110011" },

            { "D+1", "0011111" },
            { "A+1", "0110111" },
            { "M+1", "1110111" },

            { "D-1", "0001110" },
            { "A-1", "0110010" },
            { "M-1", "1110010" },

            { "D+A", "0000010" },
            { "D+M", "1000010" },

            { "D-A", "0010011" },
            { "D-M", "1010011" },

            { "A-D", "0000111" },
            { "M-D", "1000111" },

            { "D&A", "0000000" },
            { "D&M", "1000000" },

            { "D|A", "0010101" },
            { "D|M", "1010101" }
        };

        public Assembler(string file)
        {
            lines = File.ReadAllLines(file)
                .Select(l => l.Trim())
                .ToList();
            machineCode = new List<string>();
            symbols = DefaultSymbols;
            symbolCount = 16;
        }

        void RemoveWhiteSpace()
        {
            lines = lines
                .Select(l => l.Replace(" ", String.Empty))
                .Select(l => l.Split("//").First())
                .Where(l => l != "")
                .ToList();
        }

        void ParseSymbols()
        {
            var n = 0;
            var newLines = new List<string>();
            foreach (var line in lines)
            {
                var m = Symbol.Match(line);
                if (m.Success)
                {
                    symbols.Add(m.Groups["name"].Value, n);
                }
                else
                {
                    n++;
                    newLines.Add(line);
                }
            }
            lines = newLines;
        }

        void Translate()
        {
            machineCode = new List<string>();
            for (int i = 0; i < lines.Count; i++)
            {
                var line = lines[i];
                var a = ACommand.Match(line);
                var c = CCommand.Match(line);
                try
                {
                    if (a.Success)
                    {
                        machineCode.Add(TranslateACommand(a.Groups["value"].Value));
                    }
                    else if (c.Success)
                    {
                        machineCode.Add(
                                TranslateCCommand(
                                        c.Groups["comp"].Value,
                                        c.Groups["dest"].Value,
                                        c.Groups["jump"].Value
                                    )
                            );
                    } 
                    else
                    {
                        throw new Exception();
                    }
                }
                catch
                {
                    throw new Exception($"Error parsing line: #{i} : {line}");
                }
            }
        }

        private string TranslateCCommand(string comp, string dest = "", string jump = "")
        {
            var c = CompTable[comp];
            var d = DestTable[dest];
            var j = JumpTable[jump];

            return $"111{c}{d}{j}" ;
        }

        private string TranslateACommand(string value)
        {
            if (int.TryParse(value, out int i))
            {
                return Convert.ToString(i, 2).PadLeft(16, '0');
            }
            else
            {
                if (!symbols.ContainsKey(value))
                {
                    symbols.Add(value, symbolCount);
                    symbolCount++;
                }
                i = symbols[value];
                return Convert.ToString(i, 2).PadLeft(16, '0');
            }
        }

        public void Assemble()
        {
            RemoveWhiteSpace();
            ParseSymbols();
            Translate();
        }

        public void Save(string filePath) 
        {
            var code = string.Join("\n", machineCode);
            File.WriteAllText(filePath, code);
        }
    }
}
