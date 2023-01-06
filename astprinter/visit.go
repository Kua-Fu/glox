package astprinter

import (
	"fmt"
	"learning/glox/expr"
	"learning/glox/scanner"
)

// AstPrinter 用于打印某个表达式的 ast
type AstPrinter struct {
	Expr expr.Expr
}

// Print 真实打印 ast 方法
func (ap AstPrinter) Print() {

	astStr := ap.Expr.Visit()
	fmt.Println("---ast string--", astStr)
}

// StartPrint 入口函数
func StartPrint() {

	token := scanner.Token{
		Type:    scanner.MINUS,
		Lexeme:  "-",
		Literal: nil,
		Line:    1,
	}
	rightExpr := expr.Literal{
		Value: 123,
	}
	unary := expr.Unary{
		Operator: token,
		Right:    &rightExpr,
	}
	newToken := scanner.Token{
		Type:    scanner.STAR,
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
