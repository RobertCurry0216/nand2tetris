// push constant 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
// push constant 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop static 8
@static
D=M
@8
D=D+A
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D
// pop static 3
@static
D=M
@3
D=D+A
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D
// pop static 1
@static
D=M
@1
D=D+A
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D
// push static 3
@static
D=M
@3
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// push static 1
@static
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// sub
@SP
AM=M-1
D=M
A=A-1
M=D-M
// push static 8
@static
D=M
@8
A=D+A
D=M
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