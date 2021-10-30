// if-goto NAME

// pop top of stack
@SP
AM=M-1
D=M

// check eq 0
@NAME-FALSE
D;JEQ

@NAME
0;JMP

(NAME-FALSE)
