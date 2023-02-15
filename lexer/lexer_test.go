package lexer

import (
	"testing"

	"github.com/negativ/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5.0;
		let ten = 10;
		let add = fn(x, y) {
			if y == 0 {
				x + y;
			}
			else if y != 0 {
				x / y + x * y;
			}
		};
		let result = add(five, ten) << 2;
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUM, "5.0"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUM, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.IDENT, "y"},
		{token.EQ, "=="},
		{token.NUM, "0"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.IDENT, "y"},
		{token.NEQ, "!="},
		{token.NUM, "0"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.SLASH, "/"},
		{token.IDENT, "y"},
		{token.PLUS, "+"},
		{token.IDENT, "x"},
		{token.ASTERISK, "*"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.BSL, "<<"},
		{token.NUM, "2"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]: unexpected token type: %s (expected: %s)", i, tok.Type, tt.expectedType)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]: unexpected token literal: %s (expected: %s)", i, tok.Literal, tt.expectedLiteral)
		}
	}

}
