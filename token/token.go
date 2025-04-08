package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // IDENTIFIERS + LIBERALS
    IDENT = "IDENT"
    INT = "INT"

    // OPERATORS
    PLUS = "+"
    ASSIGN = "="
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    LT = "<"
    GT = ">"
    EQ = "=="
    NOT_EQ = "!="

    // DELIMETERS
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
    TRUE = "TRUE"
    FALSE = "FALSE"
)

var keywordsMap = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
    "true": TRUE,
    "false": FALSE,
}

func LookupIdent(ident string) TokenType {
    if keyword, ok := keywordsMap[ident]; ok {
        return keyword
    }
    return IDENT 
}
