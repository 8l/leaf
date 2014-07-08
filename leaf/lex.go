package leaf

import (
	"flag"
	"fmt"
	"os"

	"e8vm.net/leaf/leaf/lexer"
)

func mainLex(args []string) {
	fset := flag.NewFlagSet("leaf-lex", flag.ExitOnError)
	fset.Parse(args)

	files := fset.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no input file.")
		return
	}

	onError := func(e error) {
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
		}
	}

	for _, f := range files {
		fmt.Printf("[%s]\n", f)

		fin, e := os.Open(f)
		if e != nil {
			onError(e)
			continue
		}

		lex := lexer.New(fin, f)
		lex.OnError(onError)

		for lex.Scan() {
			tok := lex.Token()
			fmt.Println(tok.String())
		}

		onError(fin.Close())
	}
}
