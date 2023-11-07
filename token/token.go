package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符 + 字面值
	IDENT = "IDENT" // add,foobar
	INT   = "INT"   // 12374981

	// 运算符
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
