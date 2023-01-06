package expr

import (
	"fmt"
	"learning/glox/scanner"
)

// Expr 表达式接口，主要实现 visit(), 目前用于打印 ast
type Expr interface {
	Visit() string
}

// Binary 二元表达式
type Binary struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

// Grouping 圆括号表达式
type Grouping struct {
	Expression Expr
}

// Literal 标识符表达式
type Literal struct {
	Value interface{}
}

// Unary 一元表达式
type Unary struct {
	Operator scanner.Token
	Right    Expr
}

// Variable 变量表达式
type Variable struct {
	Name scanner.Token
}

// Visit 二元表达式的打印 ast 接口
func (b *Binary) Visit() string {
	return parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

// Visit 圆括号表达式的打印 ast 接口
func (g *Grouping) Visit() string {
	return parenthesize("group", g.Expression)
}

// Visit 标识符表达式的打印 ast 接口
func (l *Literal) Visit() string {
	if l.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.Value)
}

// Visit 一元表达式的打印 ast 接口
func (u *Unary) Visit() string {
	return parenthesize(u.Operator.Lexeme, u.Right)
}

// Visit 变量表达式的打印 ast 接口
func (v *Variable) Visit() string {
	return v.Name.Lexeme
}

func parenthesize(name string, exprs ...Expr) string {
	res := ""
	res += "("
	res += name
	for _, expr := range exprs {
		res += " "
		res += expr.Visit()
	}
	res += ")"
	return res

}
