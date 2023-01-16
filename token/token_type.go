package token

// Type define token type, use go enum
type Type int

const (
	// 1. one char token

	// LEFTPAREN (
	LEFTPAREN Type = iota
	// RIGHTPAREN )
	RIGHTPAREN
	// LEFTBRACE {
	LEFTBRACE
	// RIGHTBRACE }
	RIGHTBRACE
	// COMMA ,
	COMMA
	// DOT . dot
	DOT
	// MINUS -
	MINUS
	// PLUS +
	PLUS
	// SEMICOLON ;
	SEMICOLON
	// SLASH /
	SLASH
	// STAR *
	STAR

	// 2. one or two char token

	// BANG !
	BANG
	// BANGEQUAL !=
	BANGEQUAL
	// EQUAL =
	EQUAL
	// EQUALEQUAL ==
	EQUALEQUAL
	// GREATER >
	GREATER
	// GREATEREQUAL >=
	GREATEREQUAL
	// LESS <
	LESS
	// LESSEQUAL <=
	LESSEQUAL

	// text

	// IDENTIFIER identifier
	IDENTIFIER
	// STRING string
	STRING
	// NUMBER number
	NUMBER

	// keyword keyword

	// AND and
	AND
	// CLASS class
	CLASS
	// ELSE if else
	ELSE
	// FALSE false
	FALSE
	// FUN fun
	FUN
	// FOR for
	FOR
	// IF if
	IF
	// NIL nil
	NIL
	// OR or
	OR
	// PRINT print
	PRINT
	// RETURN return
	RETURN
	// SUPER super
	SUPER
	// THIS this
	THIS
	// TRUE true
	TRUE
	// VAR var
	VAR
	// WHILE while
	WHILE

	// EOF eof
	EOF
)

func (tokenType Type) String() string {
	var (
		res string
	)
	switch tokenType {
	case LEFTPAREN:
		res = "left_paren"
	case RIGHTPAREN:
		res = "right_peren"
	case LEFTBRACE:
		res = "left_brace"
	case RIGHTBRACE:
		res = "right_brace"
	case COMMA:
		res = "comma"
	case DOT:
		res = "dot"
	case MINUS:
		res = "minus"
	case PLUS:
		res = "plus"
	case SEMICOLON:
		res = "semicolon"
	case SLASH:
		res = "slash"
	case STAR:
		res = "star"
	case BANG:
		res = "bang"
	case BANGEQUAL:
		res = "bang_equal"
	case EQUAL:
		res = "equal"
	case EQUALEQUAL:
		res = "equal_equal"
	case GREATER:
		res = "greater"
	case GREATEREQUAL:
		res = "greater_equal"
	case LESS:
		res = "less"
	case LESSEQUAL:
		res = "less_equal"
	case IDENTIFIER:
		res = "identifier"
	case STRING:
		res = "string"
	case NUMBER:
		res = "number"
	case AND:
		res = "and"
	case CLASS:
		res = "class"
	case ELSE:
		res = "else"
	case FALSE:
		res = "false"
	case FUN:
		res = "fun"
	case FOR:
		res = "for"
	case IF:
		res = "if"
	case NIL:
		res = "nil"
	case OR:
		res = "or"
	case PRINT:
		res = "print"
	case SUPER:
		res = "super"
	case THIS:
		res = "this"
	case TRUE:
		res = "true"
	case VAR:
		res = "var"
	case WHILE:
		res = "while"
	case EOF:
		res = "eof"

	default:
		res = "unknown"

	}
	return res
}
