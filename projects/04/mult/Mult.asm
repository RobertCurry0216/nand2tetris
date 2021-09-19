// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

// a = R0
@R0
D=M
@a
M=D
// n = R1
@R1
D=M
@n
M=D
// sum = 0
@sum
M=0
// while n > 0 do
(LOOP)
  @n
  D=M
  @OUTPUT
  D;JEQ
  // decrease n
  @n
  M=D-1
  // sum = sum + a
  @a
  D=M
  @sum
  M=D+M

  @LOOP
  0;JMP

(OUTPUT)
//output result
@sum
D=M
@R2
M=D
//END
(END)
  @END
  0;JMP