package expr

import (
	"fmt"
	"learning/glox/token"
)

// Expr interface{} implement visit() method, to print ast
type Expr interface {
	Visit() string
}

// Binary binary expr
type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

// Grouping grouping expr
type Grouping struct {
	Expression Expr
}

// Literal literal expr
type Literal struct {
	Value interface{}
}

// Unary unary expr
type Unary struct {
	Operator token.Token
	Right    Expr
}

// Variable var expr
type Variable struct {
	Name token.Token
}

// Visit binary expr implement visit method
func (b *Binary) Visit() string {
	return parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

// Visit grouping expr implement visit method
func (g *Grouping) Visit() string {
	return parenthesize("group", g.Expression)
}

// Visit literal expr implement visit method
func (l *Literal) Visit() string {
	if l.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.Value)
}

// Visit unary expr implement visit method
func (u *Unary) Visit() string {
	return parenthesize(u.Operator.Lexeme, u.Right)
}

// Visit variable expr implement visit method
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
