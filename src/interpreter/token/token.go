
package token

type TokenType string

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    IDENT = "IDENT" 
    INT = "INT" 
    ASSIGN = "="
    PLUS = "+"
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    FUNCTION = "FUNCTION"
    LET = "LET"
)


type Token struct {
    Type TokenType
    Literal string
}


func New(tokenType TokenType, ch byte) Token {
    return Token{Type: tokenType, Literal: string(ch)}
}


