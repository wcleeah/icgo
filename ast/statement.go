package ast

import (
	"bytes"

	"com.lwc.icgo/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
    var sb bytes.Buffer

	sb.WriteString(l.TokenLiteral())
	sb.WriteString(" ")
	sb.WriteString(l.Name.String())
	sb.WriteString(" = ")
	if l.Value != nil {
		sb.WriteString(l.Value.String())
	}
	sb.WriteString(";")

    return sb.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (_ *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
    var sb bytes.Buffer

	sb.WriteString(r.TokenLiteral())
	sb.WriteString(" ")
	if r.ReturnValue != nil {
		sb.WriteString(r.ReturnValue.String())
	}
	sb.WriteString(";")

    return sb.String()
}
