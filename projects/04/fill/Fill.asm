// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

(PROG)
// i=8191
@8191
D=A
@i
M=D

// while i>= 0 do
(LOOP)
@i
D=M
@PROG
D;JLT
  // check keypress
  @KBD
  D=M
  @WHITE
  D;JEQ

  // set screen[i] to black
  @i
  D=M
  @SCREEN
  A=A+D
  M=-1
  @CONT
  0;JMP

  // set screen[i] to white
  (WHITE)
  @i
  D=M
  @SCREEN
  A=A+D
  M=0

  // i--
  (CONT)
  @i
  M=M-1

  // loop back
  @LOOP
  0;JMP
