// call f.test 2

// if narg == 0, push 0 onto the stack
// and set narg to 1

// @0
// D=A
// @SP
// A=M
// M=D
// @SP
// M=M+1

// save caller frame
// save return addr
@return-id
D=A
@SP
A=M
M=D
@SP
M=M+1

// save local
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// save arg
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// save this
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// save that
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// set arg
@7 // 2 + 5
D=A
@SP
D=M-D
@ARG
M=D

// jump to function
@f.test
0;JMP


// return addr
(return-id)


(f.test)