package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special types
	ILLEGAL = "ILLEGAL" // something we don't know about
	EOF     = "EOF"     // end of the file

	// Identifiers and Literals
	IDENT = "IDENT" // add, foobar, x, y, z...
	INT   = "INT"   // 1, 2, 3, 1234

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// if there is a matching token type for the ident string, return it
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// else, this must be an identifier
	return IDENT
}
