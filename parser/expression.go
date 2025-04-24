package parser

import (
	"fmt"
	"strconv"

	"com.lwc.icgo/ast"
	"com.lwc.icgo/token"
)

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
    msg := fmt.Sprintf("no prefix parse function for %s found", t)
    p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefixFn := p.prefixParseFns[p.curToken.Type]

    if prefixFn == nil {
        p.noPrefixParseFnError(p.curToken.Type)
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

func (p *Parser) parseIntegerLiteral() ast.Expression {
    is := p.curToken.Literal
    v, err := strconv.ParseInt(is, 0, 64)
    if err != nil {
        msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
        p.errors = append(p.errors, msg)
        return nil
    }

    return &ast.IntegerLiteral{
        Token: p.curToken,
        Value: v,
    }
}

func (p *Parser) parsePrefixExpression() ast.Expression {
    expression := &ast.PrefixExpression{
        Token: p.curToken,
        Operator: p.curToken.Literal,
    }

    p.NextToken()

    // First we handle the prefix operator
    // Then we call the parseExpression function again to process the next curToken
    // Aka the operant
    expression.Right = p.parseExpression(PREFIX)
    return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    expression := &ast.InfixExpression{
        Token: p.curToken,
        Operator: p.curToken.Literal,
        Left: left,
    }

    pre := p.curPrecedence()
    p.NextToken()
    expression.Right = p.parseExpression(pre)

    return expression
}
