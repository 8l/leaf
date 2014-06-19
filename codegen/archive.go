package codegen

// Leaf archive for a package
// contains function binaries (with unlinked calls)
// package init code
// and the symbol table
// it also contains the required imports
// the static ones should be absolute
// where the dynamic ones will be filled in on linking
type Archive struct {
}
