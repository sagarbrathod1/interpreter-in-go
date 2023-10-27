package repl

import (
	"bufio"
	"fmt"
	"interpreter-module/lexer"
	"interpreter-module/token"
	"io"
)

const PROMPT = ">> "

// read input source until newline, pass line to lexer, print all tokens until EOF
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
