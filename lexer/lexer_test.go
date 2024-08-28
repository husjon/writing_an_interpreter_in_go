package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, testToken := range tests {
		token := lexer.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q", i, testToken.expectedType, token.Type)
		}
		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q", i, testToken.expectedType, token.Type)
		}
	}
}
