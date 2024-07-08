package ast

import "potato/token"

type Node interface {
    TokenLiteral() string
}


type Statement interface {
    Node
    statementNode()
}


type Expression interface {
    Node
    expressionNode()
}


type Program struct {
    Statements []Statement
}


type LetStatement struct {
    Token token.Token
    Name *Identifier
    Value Expression
}

func (let_stmt *LetStatement) statementNode() {}
func (let_stmt *LetStatement) TokenLiteral() string { return let_stmt.Token.Literal }


type Identifier struct {
    Token token.Token
    Value string
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }


func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}
