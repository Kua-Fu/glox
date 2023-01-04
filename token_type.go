package main

// TokenType 定义token类型，使用 go enum 类型
type TokenType int

const (
	// 1. 单字符的token

	// LEFTPAREN ( 左括号
	LEFTPAREN TokenType = iota
	// RIGHTPAREN 右括号 )
	RIGHTPAREN
	// LEFTBRACE { 左花括号
	LEFTBRACE
	// RIGHTBRACE } 右花括号
	RIGHTBRACE
	// COMMA , 逗号
	COMMA
	// DOT . 点号
	DOT
	// MINUS - 减号/负号
	MINUS
	// PLUS + 加号
	PLUS
	// SEMICOLON ; 分号
	SEMICOLON
	// SLASH / 斜线/除法
	SLASH
	// STAR * 星号/乘法
	STAR

	// 2. 单个或者两个字符的token

	// BANG ! 感叹号
	BANG
	// BANGEQUAL != 不等于
	BANGEQUAL
	// EQUAL = 等于
	EQUAL
	// EQUALEQUAL == 等于
	EQUALEQUAL
	// GREATER > 大于
	GREATER
	// GREATEREQUAL >= 大于等于
	GREATEREQUAL
	// LESS < 小于
	LESS
	// LESSEQUAL <= 小于等于
	LESSEQUAL

	// 文字

	// IDENTIFIER 标识符
	IDENTIFIER
	// STRING 字符串
	STRING
	// NUMBER 数字
	NUMBER

	// keyword 关键字

	// AND and 逻辑运算符
	AND
	// CLASS class 类
	CLASS
	// ELSE if else
	ELSE
	// FALSE false 布尔
	FALSE
	// FUN fun 函数
	FUN
	// FOR for 循环
	FOR
	// IF if 判断
	IF
	// NIL nil 空值
	NIL
	// OR or 逻辑运算符
	OR
	// PRINT print 输出
	PRINT
	// RETURN return 返回
	RETURN
	// SUPER super 父类
	SUPER
	// THIS this 类
	THIS
	// TRUE true 布尔
	TRUE
	// VAR var 变量
	VAR
	// WHILE while 循环
	WHILE

	// EOF eof 结束符
	EOF
)

func (tokenType TokenType) String() string {
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
