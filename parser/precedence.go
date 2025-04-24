package parser

import "com.lwc.icgo/token"

const (
    _ int = iota
    LOWEST
    EQUALS
    LESSGREATER
    SUM
    PRODUCT
    PREFIX
    CALL
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
    if precedence, ok := precedences[p.peekToken.Type]; ok {
        return precedence
    }
    return LOWEST
}

func (p *Parser) curPrecedence() int {
    if precedence, ok := precedences[p.curToken.Type]; ok {
        return precedence
    }
    return LOWEST
}
