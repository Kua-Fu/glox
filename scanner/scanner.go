package scanner

import (
	"bufio"
	"fmt"
	"learning/glox/token"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// Scanner 扫描器
type Scanner struct {
	Logger   *zap.SugaredLogger
	HadError bool
	source   string        // source code
	runes    []rune        // source code rune slice
	tokens   []token.Token // tokens
	line     int           // location
	start    int           // token start location
	current  int           // token current location
}

// ScannerLine start scanner
func ScannerLine(line string) ([]token.Token, error) {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	s := Scanner{
		Logger:   sugar,
		HadError: false,
		line:     1,
		start:    0,
		current:  0,
	}
	s.run(line, false)
	return s.tokens, nil

}

// StartScanner start scanner
func StartScanner(args []string) ([]token.Token, error) {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	s := Scanner{
		Logger:   sugar,
		HadError: false,
		line:     1,
		start:    0,
		current:  0,
	}
	// 2. if only one arg, is source file name
	if len(args) == 2 {
		s.runFile(args[1])
	} else {
		s.runPrompt()
	}

	return s.tokens, nil

}

func (s *Scanner) runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("read file: %s, err: %v \n", path, err)
		return
	}
	s.run(string(data), true)
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
		s.run(line, true)
		s.HadError = false
	}
}

func (s *Scanner) run(source string, print bool) {
	// s.Logger.Infof("--run scanner: %s--\n", source)
	s.source = source
	s.runes = []rune(source)
	s.tokens = []token.Token{}
	s.start = 0
	s.current = 0
	s.line = 1
	s.HadError = false
	s.scanTokens()
	if print {
		for _, token := range s.tokens {
			s.Logger.Info(token)
		}
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
	endToken := token.Token{
		Type:    token.EOF,
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
		s.addToken(token.LEFTPAREN)
	case ')':
		s.addToken(token.RIGHTPAREN)
	case '{':
		s.addToken(token.LEFTBRACE)
	case '}':
		s.addToken(token.RIGHTBRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.match('=') {
			s.addToken(token.BANGEQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUALEQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESSEQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATEREQUAL)
		} else {
			s.addToken(token.GREATER)
		}

	case '/':
		if s.match('/') {
			// comment
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}

		} else {
			s.addToken(token.SLASH)
		}
	case ' ', '\r', '\t':
		// ignore null char
	case '\n':
		s.line++

	case '"':
		// string
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

// peek，will not consume char
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

func (s *Scanner) addToken(tType token.Type) {
	s.addTokenWithValue(tType, nil)
}

func (s *Scanner) addTokenWithValue(tType token.Type, value interface{}) {
	text := string(s.runes[s.start:s.current])
	token := token.Token{
		Type:    tType,
		Lexeme:  text,
		Literal: value,
		Line:    s.line,
	}
	s.tokens = append(s.tokens, token)
}

// get string value
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
	s.addTokenWithValue(token.STRING, value)
}

// number
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
	s.addTokenWithValue(token.NUMBER, fValue)

}

// identifier
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
	tType, ok := token.Keywords[text]
	if !ok {
		tType = token.IDENTIFIER
	}
	s.addToken(tType)
}
