package parser_test

import (
	"testing"

	"github.com/negativ/monkey/ast"
	"github.com/negativ/monkey/lexer"
	"github.com/negativ/monkey/parser"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	lex := lexer.New(input)
	par := parser.New(lex)
	prg := par.ParseProgram()

	if prg == nil {
		t.Fatalf("ParseProgram() returned nil!")
	}

	if len(prg.Statements) != 3 {
		t.Fatalf("Parsed statements count doesnt match: %d (%d expected)", len(prg.Statements), 3)
	}

	tests := []struct {
		expectedId string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tst := range tests {
		stmt := prg.Statements[i]

		if !testLetStatement(t, stmt, tst.expectedId) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Let statement have to start with \"let\" keyword. Got \"%s\"", s.TokenLiteral())

		return false
	}

	let, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("Let statement have to be ast.LetStatement type. Got \"%T\"", s)
	}

	if let.Name.Value != name {
		t.Errorf("let statement have to have name \"%s\". Got \"%s\"", name, let.Name.Value)
	}

	if let.Name.TokenLiteral() != name {
		t.Errorf("let statement have to have token literal \"%s\". Got \"%s\"", name, let.Name.TokenLiteral())
	}

	return true
}
