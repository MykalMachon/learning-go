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
	l.readChar() // read in the first character
	return l
}

/*
read the current character into l.ch
and increment l.position and l.readPosition
further into the input.

if we've reached the end of input,
set the l.ch value to 0 (null)
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	identPosStart := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[identPosStart:l.position]
}

func (l *Lexer) readNumber() string {
	numberPosStart := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[numberPosStart:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// skip whitespace between tokens
	l.skipWhitespace()

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
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
