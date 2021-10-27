// push constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
AM=M-1
D=M
A=A-1
M=D+M
// pop static 0
@16
D=A
@0
D=D+A
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D
