

package lexer

import (
	"interpreter/token"
)


type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}


func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}


func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}


func (l *Lexer) NextToken() token.Token {
    var t token.Token
    switch l.ch {
        case '=':
            t = token.New(token.ASSIGN, l.ch)
        case ';':
            t = token.New(token.SEMICOLON, l.ch)
        case '(':
            t = token.New(token.LPAREN, l.ch)
        case ')':
            t = token.New(token.RPAREN, l.ch)
        case ',':
            t = token.New(token.COMMA, l.ch)
        case '+':
            t = token.New(token.PLUS, l.ch)
        case '{':
            t = token.New(token.LBRACE, l.ch)
        case '}':
            t = token.New(token.RBRACE, l.ch)
        case 0:
            t.Type = token.EOF
            t.Literal = ""
    }

    l.readChar()
    return t
}
