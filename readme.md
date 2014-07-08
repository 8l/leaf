[!BuildStatus](https://travis-ci.org/e8vm/leaf.png?branch=master)](https://travis-ci.org/e8vm/leaf)

Leaf is a programming language designed for bootstraping programming for E8
virtual machine.  It is still under heavy construction, so it does not have a
formal specification yet; most of the language design is still in my mind.
Basically, it will be a C/C++ -like low-level language but with a language
syntax similar to Go language.

Here are some language features.
- Very small run time.
- Type unsafe: you can cast from any type to any type.
- Programs are organized in packages.
- No macros.
- No garbage collection.
- Can bind methods with named type (like Go)
- Has interfaces (like Go).
- Has anonymous member variable (like Go), and hence no inheritance.

As my first try here, I will only try to make the compiler work. So the
compiling might be slow, and the assembly generated won't be beautiful.

### Install

`go get e8vm.net/leaf`

### News

- `2014.6.28` We can compile and run basic hello world now!! Finally!

### TODO

- Write a testing framework
- Variable Declaring
- Operator Expressions
- Boolean Type
- If statement
- Simple for statement // similar to while
