package ast

import (
	"bytes"

	"com.lwc.icgo/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) expressionNode() {}
func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i IntegerLiteral) expressionNode() {}
func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}
func (i IntegerLiteral) String() string {
	return i.Token.Literal
}

type PrefixExpression struct {
	Token token.Token
    Operator string
    Right Expression
}

func (i PrefixExpression) expressionNode() {}
func (i PrefixExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i PrefixExpression) String() string {
    var bf bytes.Buffer

    bf.WriteString("(")
    bf.WriteString(i.Operator)
    bf.WriteString(i.Right.String())
    bf.WriteString(")")

    return bf.String()
}

type InfixExpression struct {
	Token token.Token
    Operator string
    Right Expression
    Left Expression
}
func (i InfixExpression) expressionNode() {}
func (i InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i InfixExpression) String() string {
    var bf bytes.Buffer

    bf.WriteString("(")
    bf.WriteString(i.Left.String())
    bf.WriteString(" ")
    bf.WriteString(i.Operator)
    bf.WriteString(" ")
    bf.WriteString(i.Right.String())
    bf.WriteString(")")

    return bf.String()
}

