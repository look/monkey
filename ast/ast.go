package ast

import (
	"monkey/token"
)

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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifer
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifer struct {
	Token token.Token // the IDENT token
	Value string
}

func (i *Identifer) expressionNode() {}
func (i *Identifer) TokenLiteral() string {
	return i.Token.Literal
}
