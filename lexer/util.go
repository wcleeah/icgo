package lexer

import "com.lwc.icgo/token"

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
} 

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
