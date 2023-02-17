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
	prg := &ast.Program{}
	prg.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			prg.Statements = append(prg.Statements, stmt)
		}

		p.nextToken()
	}

	return prg
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	}
	return nil
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.lookahead(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.lookahead(token.ASSIGN) {
		return nil
	}

	if !p.fastforwardTo(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) lookahead(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()

		return true
	}

	return false
}

func (p *Parser) fastforwardTo(t token.TokenType) bool {
	for {
		if p.currentToken.Type == t {
			return true
		} else if p.currentToken.Type == token.EOF {
			return false
		}

		p.nextToken()
	}
}
