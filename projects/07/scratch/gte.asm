// gte

// load top of stack into D
@SP
AM=M-1
D=M

// check
A=A-1
D=M-D

// in not greater
@is-gt
D;JGT
@end
D=0;JMP

// is greater
(is-gt)
D=-1

(end)
@SP
A=M-1
M=D