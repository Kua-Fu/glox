package parser

import (
	"bufio"
	"fmt"
	"learning/glox/astprinter"
	"learning/glox/expr"
	"learning/glox/scanner"
	"learning/glox/token"
	"os"

	"go.uber.org/zap"
)

// Parser parser source code
type Parser struct {
	Tokens  []token.Token // parser token slice
	Current int           // parser current location
}

func StartParse(args []string) {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	// get tokens
	if len(args) > 2 {
		fmt.Println("usage: glox [script]")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		tokens, err := scanner.ScannerLine(line)
		if err != nil {
			sugar.Errorf("scanner err: %v", err)
			return
		}
		parser := Parser{
			Tokens: tokens,
		}
		expr, err := parser.Parse()

		if err != nil {
			sugar.Errorf("parse err: %v", err)
			return
		}

		ast := astprinter.AstPrinter{
			Expr: expr,
		}

		ast.Print()
	}

}

func (p *Parser) Parse() (expr.Expr, error) {
	return p.expression()

}

func (p *Parser) expression() (expr.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (expr.Expr, error) {
	sExpr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.Match(token.BANGEQUAL, token.EQUALEQUAL) {
		operator := p.Previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		sExpr = &expr.Binary{
			Left:     sExpr,
			Operator: operator,
			Right:    right,
		}
	}
	return sExpr, nil
}

func (p *Parser) comparison() (expr.Expr, error) {
	sExpr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.Match(token.GREATER, token.GREATEREQUAL, token.LESS, token.LESSEQUAL) {
		operator := p.Previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		sExpr = &expr.Binary{
			Left:     sExpr,
			Operator: operator,
			Right:    right,
		}
	}

	return sExpr, nil
}

func (p *Parser) term() (expr.Expr, error) {
	sExpr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.Match(token.MINUS, token.PLUS) {
		operator := p.Previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		sExpr = &expr.Binary{
			Left:     sExpr,
			Operator: operator,
			Right:    right,
		}
	}
	return sExpr, nil
}

func (p *Parser) factor() (expr.Expr, error) {
	sExpr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.Match(token.SLASH, token.STAR) {
		operator := p.Previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		sExpr = &expr.Binary{
			Left:     sExpr,
			Operator: operator,
			Right:    right,
		}
	}
	return sExpr, nil
}

func (p *Parser) unary() (expr.Expr, error) {
	if p.Match(token.BANG, token.MINUS) {
		operator := p.Previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &expr.Unary{
			Operator: operator,
			Right:    right,
		}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (expr.Expr, error) {
	if p.Match(token.FALSE) {
		return &expr.Literal{
			Value: false,
		}, nil
	}

	if p.Match(token.TRUE) {
		return &expr.Literal{
			Value: true,
		}, nil
	}

	if p.Match(token.NIL) {
		return &expr.Literal{
			Value: nil,
		}, nil
	}

	if p.Match(token.NUMBER, token.STRING) {
		return &expr.Literal{
			Value: p.Previous().Literal,
		}, nil
	}
	if p.Match(token.LEFTPAREN) {
		sExpr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(token.RIGHTPAREN, "expect ')' after expression")
		if err != nil {
			return nil, err
		}
		return &expr.Grouping{
			Expression: sExpr,
		}, nil
	}

	return nil, nil

}

func (p *Parser) consume(tType token.Type, message string) (token.Token, error) {
	if p.Check(tType) {
		return p.Advance(), nil
	}
	pToken := p.Peek()
	if pToken.Type == token.EOF {
		return token.Token{}, fmt.Errorf(
			"[line %d] Error %s: %s",
			pToken.Line,
			"at end",
			message)
	}
	return token.Token{}, fmt.Errorf(
		"[line %d] Error %s: %s",
		pToken.Line,
		"at '"+pToken.Lexeme+"'",
		message)

}

// Match check next token type match
func (p *Parser) Match(types ...token.Type) bool {
	for _, item := range types {
		if p.Check(item) {
			p.Advance()
			return true
		}
	}
	return false
}

// Check check token type
func (p *Parser) Check(tType token.Type) bool {
	if p.IsAtEnd() {
		return false
	}
	return p.Peek().Type == tType
}

// Advance consume one token
func (p *Parser) Advance() token.Token {
	if !p.IsAtEnd() {
		p.Current++
	}
	return p.Previous()
}

// IsAtEnd check token type is EOF
func (p *Parser) IsAtEnd() bool {
	return p.Peek().Type == token.EOF
}

// Peek not consume tokens
func (p *Parser) Peek() token.Token {
	return p.Tokens[p.Current]
}

// Previous get previous token
func (p *Parser) Previous() token.Token {
	return p.Tokens[p.Current-1]
}
