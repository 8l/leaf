# asm8

The assembler for E8 virtual machine.


### Install and Test Run

```
$ go install e8vm.net/asm8
$ cat hello.s # $GOPATH/src/e8vm.net/leaf/asm8/tests/hello.s
    add $1, $0, $0      ; init counter
.loop:
    lbu $2, 0x1000($1)  ; load byte
    ; lbu $2, msg($1)   ; load byte
    beq $2, $0, .end    ; +5
.wait:
    lbu $3, 9           ; is output ready?
    bne $3, $0, .wait   ; -2
    sb $2, 9            ; output byte
    addi $1, $1, 1      ; update counter
    j .loop             ; -7
.end:
    sb $0, 0x8($0)
$ asm8 hello.s
Hello, world.
(102 cycles)
```

## New Grammar BNF

```
<Program> = <Decl>*
<Decl> = <ConstDecl> | <VarDecl> | <FuncDecl>
<ConstDecl> = "const" <ConstSpec>
    | "const" "(" <ConstSpec>* ")" ";"
<ConstSpec> = <ident> "=" <int> ";"
<VarDecl> = "var" <VarSpec>
    | "var" "(" <VarSpec>* ")" ";"
<VarSpec> = <ident> <Type> "=" <Value> ";"
<Type> = <BasicType> | <ArrayType>
<ArrayType> = "[" [ <int> ] "]" <BasicType>
<BasicType> = "u8" | "i8" | "u16" | "i16" | "u32" | "i32" 
    | "f64" | "str"
<Value> = <BasicValue> | <ArrayValue> | <string>
<BasicValue> = <int> | <float> | <char> 
<ArrayValue> = "{" [ <BasicValueList> ] "}"
<BasicValueList> = <BasicValue> | <BasicValueList> "," <BasicValue>
<FuncDecl> = "func" "{" <Line>* "}" ";"
<Line> = [ <Label> ] <Op> [ <Args> ] ";"
<Label> = <ident> ":"
<Op> = <ident>
<Args> = <Arg> | <Args> "," <Arg>
<Arg> = "$" <int> 
    | "(" "$" <int> ")"
    | <int> ["(" "$" <int> ")"]
    | <ident> ["(" "$" <int> ")"]
```
