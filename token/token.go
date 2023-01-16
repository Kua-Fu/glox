package token

import "fmt"

// Token token struct
type Token struct {
	Type    Type        // token type
	Lexeme  string      // token string value
	Literal interface{} // token real value
	Line    int         // token location
}

var (
	// Keywords keyword
	Keywords = map[string]Type{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
)

func (t Token) String() string {
	return fmt.Sprintf("type: %s, lexme: %s, literal: %v", t.Type, t.Lexeme, t.Literal)
}
