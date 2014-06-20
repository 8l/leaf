Go lanugage has a weird symbol importing structure.

- top level decls are defined in a package level, but uses a file
  level scope.
- so to compile a file, it is a little bit twisted
- that is when we build a node in a file here, we need to enter the
  file scope.

- or, alternatively, I can just change the grammar so that each module
  will be one single file.
- we will then be able to do the compiling with several passes
  - register stage: we first get the symbol types,
    - it could be import, const, type, var, or func
    - we will also support anonymous imports
    - function init will be treated differently
  - layout stage: // this is like building the header file
    - at this stage, we resolve the types of each symbol
    - for types, we will know how to memory layout the symbol
    - for functions, we will know how to memory layout the stack for
      calling this function
    - for consts, we will evaluate its value, and consts are done here
    - for vars, we will know the memory layout of the var
    - after this stage, the package should be ready to be imported by
      other packages, at least conceptually
  - eval stage // 
    - for functions, we will build the function body
    - for vars, we will build the init routine
    - and we will also build all the init functions // or shall there
      be only one init?

const something; // register something as a type
const type; //
var something;
func a;
type t;

func a();
type t struct {
    ...
}

var something = a(); 
