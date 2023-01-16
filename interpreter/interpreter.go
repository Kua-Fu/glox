package interpreter

import (
	"bufio"
	"fmt"
	"learning/glox/expr"
	"learning/glox/parser"
	"learning/glox/scanner"
	"os"

	"go.uber.org/zap"
)

var (
	l *zap.SugaredLogger
)

// Interpreter interpreter struct
type Interpreter struct {
	Expr expr.Expr
}

// Evaluate evaluate expr
func (i *Interpreter) Evaluate() (interface{}, error) {
	res, err := i.Expr.Evaluate()
	if err != nil {
		return res, err
	}
	return res, nil
}

// StartInterpreter start interpreter, evaluate expr
func StartInterpreter(args []string) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	l = logger.Sugar()
	// get tokens
	if len(args) > 2 {
		l.Info("usage: glox [script]")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		tokens, err := scanner.ScanLine(line)
		if err != nil {
			l.Errorf("scanner err: %v", err)
			continue
		}
		parser := parser.Parser{
			Tokens: tokens,
		}
		expr, err := parser.Parse()

		if err != nil {
			l.Errorf("parse err: %v", err)
			continue
		}

		interpreter := Interpreter{
			Expr: expr,
		}

		res, err := interpreter.Evaluate()
		if err != nil {
			l.Errorf("eval err: %v", err)
			continue
		}
		fmt.Printf("--eval-- %v\n", res)

	}

}
