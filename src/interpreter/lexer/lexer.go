

package lexer

import (
	"potato/token"
)


type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}


func NewLexer(input string) *Lexer {
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


func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}


func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}


func (l *Lexer) NextToken() token.Token {
    var t token.Token
    l.skipWhitespace()
    switch l.ch {
    default:
        if isLetter(l.ch) {
            t.Literal = l.readIdentifier()
            t.Type = token.LookupIdent(t.Literal)
            return t
        } else if isDigit(l.ch) {
            t.Type = token.INT
            t.Literal = l.readNumber()
            return t
        } else {
            t = token.NewToken(token.ILLEGAL, l.ch)
        }
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            t = token.NewToken(token.EQ, ch+ch)
        } else {
            t = token.NewToken(token.ASSIGN, l.ch)
        }
    case '+':
        t = token.NewToken(token.PLUS, l.ch)
    case '-':
        t = token.NewToken(token.MINUS, l.ch)
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            t = token.NewToken(token.NEQ, ch + l.ch)
        } else {
            t = token.NewToken(token.BANG, l.ch)
        }
    case '/':
        t = token.NewToken(token.SLASH, l.ch)
    case '*':
        t = token.NewToken(token.ASTERISK, l.ch)
    case '<':
        t = token.NewToken(token.LT, l.ch)
    case '>':
        t = token.NewToken(token.GT, l.ch)
    case ';':
        t = token.NewToken(token.SEMICOLON, l.ch)
    case '(':
        t = token.NewToken(token.LPAREN, l.ch)
    case ')':
        t = token.NewToken(token.RPAREN, l.ch)
    case ',':
        t = token.NewToken(token.COMMA, l.ch)
    case '{':
        t = token.NewToken(token.LBRACE, l.ch)
    case '}':
        t = token.NewToken(token.RBRACE, l.ch)
    case 0:
        t.Type = token.EOF
        t.Literal = ""
    }
    l.readChar()
    return t
}


func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}


func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}


func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
