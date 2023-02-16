package ast

import "github.com/negativ/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {

}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

type LetStatement struct {
	Name  *LetStatement
	Value *Expression
	Token token.Token
}

func (let *LetStatement) statementNode() {

}

func (let *LetStatement) TokenLiteral() string {
	return let.Token.Literal
}
