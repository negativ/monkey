package parser

import (
	"github.com/negativ/monkey/ast"
	"github.com/negativ/monkey/lexer"
	"github.com/negativ/monkey/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}
