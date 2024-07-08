

package parser

import (
    "potato/ast"
    "potato/lexer"
    "potato/token"
)


type Parser struct {
    l *lexer.Lexer

    currToken token.Token
    peekToken token.Token
}


func NewParser(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    p.NextToken()
    p.NextToken()
    
    return p
}


func (p *Parser) NextToken() {
    p.currToken = p.peekToken
    p.peekToken = p.l.NextToken()
}


func (p *Parser) ParseProgram() *ast.Program {
    return nil
}


