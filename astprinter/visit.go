package astprinter

import (
	"fmt"
	"learning/glox/expr"
	"learning/glox/token"
)

// AstPrinter print expr ast
type AstPrinter struct {
	Expr expr.Expr
}

// Print print ast
func (ap AstPrinter) Print() {
	astStr := ap.Expr.Visit()
	fmt.Println("--ast--", astStr)
}

// StartPrint main method
func StartPrint() {

	iToken := token.Token{
		Type:    token.MINUS,
		Lexeme:  "-",
		Literal: nil,
		Line:    1,
	}
	rightExpr := expr.Literal{
		Value: 123,
	}
	unary := expr.Unary{
		Operator: iToken,
		Right:    &rightExpr,
	}
	newToken := token.Token{
		Type:    token.STAR,
		Lexeme:  "*",
		Literal: nil,
		Line:    1,
	}

	groupExpr := expr.Literal{
		Value: 45.67,
	}

	group := expr.Grouping{
		Expression: &groupExpr,
	}
	expression := expr.Binary{
		Left:     &unary,
		Operator: newToken,
		Right:    &group,
	}

	astPrinter := AstPrinter{
		Expr: &expression,
	}

	astPrinter.Print()

}
