package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// Scanner 扫描器
type Scanner struct {
	Logger   *zap.SugaredLogger
	HadError bool
	source   string  // 原始代码
	runes    []rune  // 原始代码的字符列表
	tokens   []Token // 词法分词列表
	line     int     // 当前所在的行信息
	start    int     // 标记出现的开始位置
	current  int     // 当前扫描的位置
}

func main() {
	args := os.Args
	// 1. 如果传参大于1个，报错
	if len(args) > 2 {
		fmt.Println("usage: glox [script]")
		return
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	scanner := Scanner{
		Logger:   sugar,
		HadError: false,
		line:     1,
		start:    0,
		current:  0,
	}
	// 2. 如果传参为1个，表示文件名称
	if len(args) == 2 {
		scanner.runFile(args[1])
	} else {
		scanner.runPrompt()
	}

}

func (s *Scanner) runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("read file: %s, err: %v \n", path, err)
		return
	}
	s.run(string(data))
	if s.HadError {
		os.Exit(65)
	}
}

func (s *Scanner) runPrompt() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		s.run(line)
		s.HadError = false
	}
}

func (s *Scanner) run(source string) {
	// s.Logger.Infof("--run scanner: %s--\n", source)
	s.source = source
	s.runes = []rune(source)
	s.tokens = []Token{}
	s.start = 0
	s.current = 0
	s.line = 1
	s.HadError = false
	s.scanTokens()
	for _, token := range s.tokens {
		s.Logger.Info(token)
	}

}

func (s *Scanner) error(line int, message string) {
	s.report(line, "", message)

}

func (s *Scanner) report(line int, where, message string) {

	s.Logger.Errorf(
		"[line %d] Error %s: %s", line, where, message,
	)

	s.HadError = true
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.runes)
}

func (s *Scanner) scanTokens() {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()

	}
	endToken := Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	}
	s.tokens = append(s.tokens, endToken)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFTPAREN)
	case ')':
		s.addToken(RIGHTPAREN)
	case '{':
		s.addToken(LEFTBRACE)
	case '}':
		s.addToken(RIGHTBRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANGEQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUALEQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESSEQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATEREQUAL)
		} else {
			s.addToken(GREATER)
		}

	case '/':
		if s.match('/') {
			// 注释
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}

		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		// 忽略空白字符
	case '\n':
		s.line++

	case '"':
		// 字符串常量
		s.addString()
	default:
		if s.isDigit(c) {
			s.addNumber()
		} else if s.isAlpha(c) {
			s.addIdentifier()

		} else {
			s.error(s.line, "unexpected character.")
		}
	}
}

func (s *Scanner) advance() rune {
	s.current++
	return s.runes[s.current-1]
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.runes[s.current-1] != expected {
		return false
	}

	s.current++
	return true
}

// 前瞻，不会消耗字符
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.runes[s.current]
}

func (s *Scanner) peekNext() rune {
	if s.current >= len(s.runes) {
		return '\x00'
	}
	return s.runes[s.current+1]
}

func (s *Scanner) addToken(tType TokenType) {
	s.addTokenWithValue(tType, nil)
}

func (s *Scanner) addTokenWithValue(tType TokenType, value interface{}) {
	text := string(s.runes[s.start:s.current])
	token := Token{
		Type:    tType,
		Lexeme:  text,
		Literal: value,
		Line:    s.line,
	}
	s.tokens = append(s.tokens, token)
}

// 获取字符串常量
func (s *Scanner) addString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.error(s.line, "unterminal string")
		return
	}

	s.advance()
	value := s.runes[s.start+1 : s.current-1]
	fmt.Println("--add string--", value)
	s.addTokenWithValue(STRING, value)
}

// 数字
func (s *Scanner) isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func (s *Scanner) addNumber() {

	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	value := string(s.runes[s.start:s.current])

	fValue, _ := strconv.ParseFloat(value, 64)
	fmt.Println("--add value--", s.runes, s.start, s.current, fValue)
	s.addTokenWithValue(NUMBER, fValue)

}

// 标识符
func (s *Scanner) isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r == '_')
}

func (s *Scanner) isAlphaNumeric(r rune) bool {
	return s.isAlpha(r) || s.isDigit(r)
}

func (s *Scanner) addIdentifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := string(s.runes[s.start:s.current])
	tType, ok := Keywords[text]
	if !ok {
		tType = IDENTIFIER
	}
	s.addToken(tType)
}
