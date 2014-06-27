[![BuildStatus](https://travis-ci.org/e8vm/leaf.png?branch=master)](https://travis-ci.org/e8vm/leaf)

Leaf is a programming language designed for bootstraping programming for E8
virtual machine.  It is still under heavy construction, so it does not have a
formal specification yet; most of the language design is still in my mind.
Basically, it will be a C/C++ -like low-level language but with a language
syntax similar to Go language.

Here are some language features.
- Very small run time.
- No macros.
- No garbage collections.
- Programs are organized in packages.
- Type unsafe: you can cast from any type to any type.
- Can bind methods with named type.
- Implicitly supports interfaces.
- Supports anonymous member variable, and hence do not support inheritance.

As my first try here, I will only try to make the compiler work. So the
compiling might be slow, and the assembly generated won't be beautiful.

