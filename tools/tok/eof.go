package tok

type eof int

// EOF is a special token for end of file.
// A tokenizer should return this token as the last token
// in a token stream. If a user requests for more tokens
// after EOF is returned, a tokenizer should panic.
const EOF eof = -1

func (t eof) String() string { return "EOF" }
func (t eof) Code() int      { return -1 }
func (t eof) IsSymbol() bool { return true }
