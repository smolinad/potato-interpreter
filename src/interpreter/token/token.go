
package token

type TokenType string

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    IDENT = "IDENT" 
    INT = "INT" 
    
    ASSIGN = "="
    EQ = "=="
    NEQ = "!="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LT = "<"
    GT = ">"

    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    
    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)


type Token struct {
    Type TokenType
    Literal string
}


func NewToken(tokenType TokenType, ch byte) Token {
    return Token{Type: tokenType, Literal: string(ch)}
}


var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}


func LookupIdent(ident string) TokenType {
    if t, ok := keywords[ident]; ok {
        return t
    }
    return IDENT
}
