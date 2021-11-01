// return

// set arg[0] to return val
@SP
A=M-1
D=M
@ARG
A=M
M=D

// set SP back to LCL[0]
@LCL
D=M
@SP
M=D

// save arg[1] addr so it can be set back to SP
@ARG
D=M+1
@R14
M=D

// reset that
@SP
AM=M-1
D=M
@THAT
M=D

// reset this
@SP
AM=M-1
D=M
@THIS
M=D

//reset arg
@SP
AM=M-1
D=M
@ARG
M=D

// reset lcl
@SP
AM=M-1
D=M
@LCL
M=D

// reset addr and restore SP
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