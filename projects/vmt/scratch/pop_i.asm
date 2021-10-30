// pop local 4

// set A to LCL+4
@LCL
D=M
@4
D=D+A

// store in R15
@R15
M=D

// get top of stack and dec SP
@SP
AM=M-1
D=M

// put value into place
@R15
A=M
M=D