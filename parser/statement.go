package parser

import (
	"com.lwc.icgo/ast"
	"com.lwc.icgo/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Handle expression/statement.Value part
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    statement := &ast.ReturnStatement{
        Token: p.curToken,
    }

    p.NextToken()
	// TODO: Handle expression/statement.Value part
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return statement
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    statement := &ast.ExpressionStatement{
        Token: p.curToken,
    } 

    statement.Expression = p.parseExpression(LOWEST)

    if p.peekTokenIs(token.SEMICOLON) {
        p.NextToken()
    }
    return statement 
}

