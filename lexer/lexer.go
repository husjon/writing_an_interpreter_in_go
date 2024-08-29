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

func (l *Lexer) peekChar() byte {
	// Reads a single character from the input _without_ incrementing the pointer
	if l.readPosition >= len(l.input) {
		// Reset the current character if we've read past the input (aka end of file)
		return 0 // Since this is a byte, this is a NUL-byte
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.chr) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.chr) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	// I know what it doing, but I don't really understand how it works.
	// This is mainly because of the Go syntax.
	// My assumption is that since `l.chr` already contains a byte,
	//  the first time the for loop goes through, the conditional is met and l.readChar is called.
	// Since it was met, it can loop again until it is no longer met and finally breaks out.

	for l.chr == ' ' || l.chr == '\t' || l.chr == '\n' || l.chr == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
	// I kinda like this syntax, but I wish it was possible to split it over multiple lines,
	// alternatively how it's done in Python (Go + Python pseudocode) with `'a' <= ch <= 'z'`
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	case '-':
		tok = newToken(token.MINUS, l.chr)
	case '*':
		tok = newToken(token.ASTERISK, l.chr)
	case '/':
		tok = newToken(token.SLASH, l.chr)
	case '!':
		tok = newToken(token.BANG, l.chr)
	case '<':
		tok = newToken(token.LT, l.chr)
	case '>':
		tok = newToken(token.GT, l.chr)
	case '{':
		tok = newToken(token.LBRACE, l.chr)
	case '}':
		tok = newToken(token.RBRACE, l.chr)
	case 0: // NUL-byte / End-of-file
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.chr) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.chr) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.chr)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType, Literal: string(ch),
	}
}
