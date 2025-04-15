package ast

import "com.lwc.icgo/token"


type ExpressionStatement struct {
    Token token.Token
    Expression Expression
}

func (es *ExpressionStatement) expressionNode() {}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
    return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}
