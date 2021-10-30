// push local 2

// load value into D
@LCL
D=M
@2
A=D+A
D=M

// push to stack
@SP
A=M
M=D

// inc sp
@SP
M=M+1