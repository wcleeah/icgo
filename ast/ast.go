package ast

import "bytes"

type Node interface {
	// for debug use
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	// for separating the interface with Expression
	statementNode()
}

type Expression interface {
	Node
	// for separating the interface with Statement
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p Program) String() string {
	var sb bytes.Buffer
	for _, s := range p.Statements {
		sb.WriteString(s.String())
	}
    return sb.String()
}
