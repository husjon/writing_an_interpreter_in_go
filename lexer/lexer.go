package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // The current position in input (points to the current character)
	readPosition int  // The current reading position in input (after the current character)
	chr          byte // The current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// Reads a single character from the input and increments the pointers
	if l.readPosition >= len(l.input) {
		// Reset the current character if we've read past the input (aka end of file)
		l.chr = 0 // Since this is a byte, this is a NUL-byte
	} else {
		l.chr = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.chr {
	case '=':
		tok = newToken(token.ASSIGN, l.chr)
	case ';':
		tok = newToken(token.SEMICOLON, l.chr)
	case '(':
		tok = newToken(token.LPAREN, l.chr)
	case ')':
		tok = newToken(token.RPAREN, l.chr)
	case ',':
		tok = newToken(token.COMMA, l.chr)
	case '+':
		tok = newToken(token.PLUS, l.chr)
	case '{':
		tok = newToken(token.LBRACE, l.chr)
	case '}':
		tok = newToken(token.RBRACE, l.chr)
	case 0: // NUL-byte / End-of-file
		tok.Literal = ""
		tok = newToken(token.EOF, l.chr)
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType, Literal: string(ch),
	}
}
