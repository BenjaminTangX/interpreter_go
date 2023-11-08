package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置(指向当前字符)
	readPosition int  // 所输入字符串中的当前读取位置(指向当前字符串之后的一个字符)
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0: // l.ch == 0 不是 '0'
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

// 此处readChar方法只支持ASCII字符，不能支持所有Unicode字符
// 可以考虑将ch设置为类型rune
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 通过设置tokenType与ch，构建token.Token类型
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
