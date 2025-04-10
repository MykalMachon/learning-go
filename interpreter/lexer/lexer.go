package lexer

import "mykal.codes/monkey/token"

type Lexer struct {
	input        string
	position     int  // the current position in the input
	readPosition int  // the current reading pos in input
	ch           byte // current character being examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar() // read in the first character
	return l
}

/*
*

	read the current character into l.ch
	and increment l.position and l.readPosition
	further into the input.

	if we've reached the end of input,
	set the l.ch value to 0 (null)
*/
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
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
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.ReadChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
