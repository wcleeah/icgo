package parser

import (
	"com.lwc.icgo/ast"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefixFn := p.prefixParseFns[p.curToken.Type]

    if prefixFn == nil {
        return nil
    }
    leftExp := prefixFn() 

    return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{
        Token: p.curToken,
        Value: p.curToken.Literal,
    }
}
