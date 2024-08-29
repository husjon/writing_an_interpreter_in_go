package token

type TokenType string

type Token struct {
	// We set the TokenType to be string, allowing us easier debugging at the cost of performance.
	// Another way would be to use ints or bytes.
	Type    TokenType
	Literal string
}

// Token Types
const (
	ILLEGAL = "ILLEGAL" // Represents any token that cannot be parsed.
	EOF     = "EOF"     // Represents the end of the file

	// Identifiers + literals
	IDENT = "IDENT" // Variable names like: `add`, `foobar`, `x`, `y`
	INT   = "INT"   // Integers: 1, 2, 3, 100, 10000 etc

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Comparison
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	// Performs a map lookup of the provided string and if it exists, returns the Token

	if tok, ok := keywords[ident]; ok { // Performs the assignment and conditional check (notice the `;`)
		return tok
	}

	return IDENT
}
