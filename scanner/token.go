package scanner

import "fmt"

// Token token 结构体
type Token struct {
	Type    TokenType   // token类型
	Lexeme  string      // 词素
	Literal interface{} // 实际的值
	Line    int         // token所在的行
}

var (
	// Keywords 关键字
	Keywords = map[string]TokenType{
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
