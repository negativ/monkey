package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/negativ/monkey/token"
)

type Lexer struct {
	buffer  string
	pos     int
	readPos int
	char    rune
}

var simpleTokens = map[rune]token.TokenType{
	'*': token.ASTERISK,
	'/': token.SLASH,
	'(': token.LPAREN,
	')': token.RPAREN,
	'{': token.LBRACE,
	'}': token.RBRACE,
	',': token.COMMA,
	';': token.SEMICOLON,
}

var keywords = map[string]token.TokenType{
	"let":   token.LET,
	"fn":    token.FUNCTION,
	"if":    token.IF,
	"else":  token.ELSE,
	"true":  token.TRUE,
	"false": token.FALSE,
}

func New(data string) *Lexer {
	l := &Lexer{buffer: data}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespaces()

	if tok, ok := l.readOperator(); ok {
		l.readChar()

		return tok
	} else if l.isLetter() {
		lit := l.readIdent()
		t := parseKeyword(lit)

		return newToken(t, lit)
	} else if l.isDigit() {
		return newToken(token.NUM, l.readNum())
	} else if l.char == utf8.MaxRune {
		return newToken(token.EOF, "")
	}

	return newToken(token.ILLEGAL, "")
}

func newToken(t token.TokenType, lit string) token.Token {
	return token.Token{Type: t, Literal: lit}
}

func (l *Lexer) readChar() {
	if len(l.buffer) <= l.readPos {
		l.char = utf8.MaxRune

		return
	}

	r, size := utf8.DecodeRuneInString(l.buffer[l.readPos:])

	if r == utf8.RuneError || size == 0 {
		l.char = utf8.MaxRune
		l.pos = len(l.buffer)
		l.readPos = l.pos

		return
	}

	l.char = r
	l.pos = l.readPos
	l.readPos += size
}

func (l *Lexer) peekChar() rune {
	if l.readPos >= len(l.buffer) {
		return utf8.MaxRune
	}

	r, _ := utf8.DecodeRuneInString(l.buffer[l.readPos:])

	return r
}

func (l *Lexer) readOperator() (token.Token, bool) {
	next := l.peekChar()
	var tok token.Token

	if t, ok := simpleTokens[l.char]; ok {
		tok := newToken(t, string(l.char))

		return tok, true
	}

	switch l.char {
	case '+':
		if next == '+' {
			tok = newToken(token.INC, "++")
		} else {
			tok = newToken(token.PLUS, string(l.char))
		}
	case '-':
		if next == '-' {
			tok = newToken(token.DEC, "--")
		} else {
			tok = newToken(token.MINUS, string(l.char))
		}
	case '!':
		if next == '=' {
			tok = newToken(token.NEQ, "!=")
		} else {
			tok = newToken(token.BANG, string(l.char))
		}
	case '=':
		if next == '=' {
			tok = newToken(token.EQ, "==")
		} else {
			tok = newToken(token.ASSIGN, string(l.char))
		}
	case '<':
		if next == '=' {
			tok = newToken(token.LTE, "<=")
		} else if next == '<' {
			tok = newToken(token.BSL, "<<")
		} else {
			tok = newToken(token.LT, string(l.char))
		}
	case '>':
		if next == '=' {
			tok = newToken(token.GTE, ">=")
		} else if next == '>' {
			tok = newToken(token.BSL, ">>")
		} else {
			tok = newToken(token.GT, string(l.char))
		}
	default:
		return newToken(token.ILLEGAL, ""), false
	}

	l.readChar()

	return tok, true
}

func (l *Lexer) readIdent() string {
	pos := l.pos

	for l.isLetter() {
		l.readChar()
	}

	return l.buffer[pos:l.pos]
}

func (l *Lexer) readNum() string {
	pos := l.pos
	scan := func() {
		for l.isDigit() {
			l.readChar()
		}
	}

	scan()

	if l.isDot() {
		l.readChar()
		scan()
	}

	return l.buffer[pos:l.pos]
}

func parseKeyword(kw string) token.TokenType {
	if t, ok := keywords[kw]; ok {
		return t
	}

	return token.IDENT
}

func (l *Lexer) skipWhitespaces() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) isLetter() bool {
	return unicode.IsLetter(l.char)
}

func (l *Lexer) isDigit() bool {
	return unicode.IsDigit(l.char)
}

func (l *Lexer) isDot() bool {
	return l.char == '.'
}

func (l *Lexer) isEOF() bool {
	return l.char == utf8.MaxRune
}

func (l *Lexer) isError() bool {
	return l.char == utf8.RuneError
}
