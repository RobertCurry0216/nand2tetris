// function SimpleFunction.test 2
(SimpleFunction.test)
@SP
D=M
@LCL
AM=D
M=0
A=A+1
M=0
A=A+1
D=A
@SP
M=D
// push local 0
@LCL
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
// push local 1
@LCL
D=M
@1
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
// not
@SP
A=M-1
M=!M
// push argument 0
@ARG
A=M
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
// push argument 1
@ARG
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
M=M-D
// return
@SP
A=M-1
D=M
@ARG
A=M
M=D
@LCL
D=M
@SP
M=D
@ARG
D=M+1
@R14
M=D
@SP
AM=M-1
D=M
@THAT
M=D
@SP
AM=M-1
D=M
@THIS
M=D
@SP
AM=M-1
D=M
@ARG
M=D
@SP
AM=M-1
D=M
@LCL
M=D
@SP
AM=M-1
D=M
@R15
M=D
@R14
D=M
@SP
M=D
@R15
A=M;JMP
