package expr

import (
	"fmt"
	"learning/glox/token"
	"learning/glox/utils"
)

// Expr interface{} implement visit() method, to print ast
type Expr interface {
	Visit() string
	Evaluate() (interface{}, error)
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

// Evaluate binary expr implement evaluate method
func (b *Binary) Evaluate() (interface{}, error) {
	left, err := b.Left.Evaluate()
	if err != nil {
		return left, err
	}

	right, err := b.Right.Evaluate()
	if err != nil {
		return right, err
	}

	switch b.Operator.Type {

	case token.BANGEQUAL:
		return !isEqual(left, right), nil

	case token.EQUALEQUAL:
		return isEqual(left, right), nil

	case token.GREATER:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft > fRight, nil

	case token.GREATEREQUAL:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft >= fRight, nil

	case token.LESS:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft < fRight, nil

	case token.LESSEQUAL:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft <= fRight, nil

	case token.MINUS:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft - fRight, nil

	case token.SLASH:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft / fRight, nil

	case token.STAR:
		fLeft, err := utils.GetFloatNumber(left)
		if err != nil {
			return fLeft, err
		}
		fRight, err := utils.GetFloatNumber(right)
		if err != nil {
			return fRight, err
		}
		return fLeft * fRight, nil

	case token.PLUS:

		switch left.(type) {
		case string:
			sLeft := left.(string)
			sRight, ok := right.(string)
			if !ok {
				return nil, fmt.Errorf(
					"plus operator should be two string, or two float")
			}

			return sLeft + sRight, nil
		default:
			fLeft, err := utils.GetFloatNumber(left)
			if err != nil {
				return nil, err
			}
			fRight, err := utils.GetFloatNumber(right)
			if err != nil {
				return nil, err
			}
			return fLeft + fRight, nil
		}

	}

	return nil, nil
}

// Visit grouping expr implement visit method
func (g *Grouping) Visit() string {
	return parenthesize("group", g.Expression)
}

// Evaluate grouping expr implement evaluate method
func (g *Grouping) Evaluate() (interface{}, error) {
	return g.Expression.Evaluate()
}

// Visit literal expr implement visit method
func (l *Literal) Visit() string {
	if l.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.Value)
}

// Evaluate literal expr implement evaluate method
func (l *Literal) Evaluate() (interface{}, error) {
	return l.Value, nil
}

// Visit unary expr implement visit method
func (u *Unary) Visit() string {
	return parenthesize(u.Operator.Lexeme, u.Right)
}

// Evaluate unary expr implement evaluate method
func (u *Unary) Evaluate() (interface{}, error) {
	right, err := u.Right.Evaluate()
	if err != nil {
		return right, err
	}
	switch u.Operator.Type {
	case token.BANG:
		return !isTruthy(right), nil
	case token.MINUS:
		fNumber, err := utils.GetFloatNumber(right)
		if err != nil {
			return fNumber, err
		}
		return -1 * fNumber, nil
	}
	return nil, nil
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

func isTruthy(data interface{}) bool {
	if data == nil {
		return false
	}
	bData, ok := data.(bool)
	if ok {
		return bData
	}
	return true
}

func isEqual(left, right interface{}) bool {
	return left == right
}
