package token

type TokenType uint32

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF

	IDENT
	NUM

	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH
	BANG
	LT
	GT
	LTE
	GTE
	EQ
	NEQ
	INC
	DEC
	BSL
	BSR

	AND
	OR
	XOR

	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET
	IF
	ELSE
	TRUE
	FALSE
	RET
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "<ILLEGAL>"
	case EOF:
		return "<EOF>"
	case IDENT:
		return "<IDENT>"
	case NUM:
		return "<NUM>"
	case ASSIGN:
		return "<ASSIGN>"
	case PLUS:
		return "<PLUS>"
	case COMMA:
		return "<COMMA>"
	case SEMICOLON:
		return "<SEMI>"
	case LPAREN:
		return "<LPAR>"
	case RPAREN:
		return "<RPAR>"
	case LBRACE:
		return "<LBR>"
	case RBRACE:
		return "<RBR>"
	case FUNCTION:
		return "<FUNC>"
	case LET:
		return "<LET>"
	case ASTERISK:
		return "<ASTERISK>"
	case SLASH:
		return "<SLASH>"
	case BANG:
		return "<BANG>"
	case MINUS:
		return "<MINUS>"
	case LT:
		return "<LT>"
	case GT:
		return "<GT>"
	case EQ:
		return "<EQ>"
	case AND:
		return "<AND>"
	case OR:
		return "<OR>"
	case XOR:
		return "<XOR>"
	case IF:
		return "<IF>"
	case ELSE:
		return "<ELSE>"
	case TRUE:
		return "<TRUE>"
	case FALSE:
		return "<FALSE>"
	case INC:
		return "<INC>"
	case DEC:
		return "<DEC>"
	case BSL:
		return "<BSL>"
	case BSR:
		return "<BSR>"
	case NEQ:
		return "<NEQ>"
	case RET:
		return "<RET>"
	}

	panic("Invalid argument")
}
