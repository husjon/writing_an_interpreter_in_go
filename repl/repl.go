package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// Initializes a scanner on the input buffer
	scanner := bufio.NewScanner(in)

	for { // A for loop in Go without a condition is an infinite loop
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan() // Performs a scan either to the end of input (aka end of line in this case) or and error.

		if !scanned {
			return
		}

		line := scanner.Text() // Gets the text as a string
		l := lexer.New(line)

		// Iterates over the input and prints out the provided tokens that was parsed until EOF.
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
