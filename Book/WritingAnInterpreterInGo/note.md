# Writing an Interpreter in Go



## 第1章 字句解析


### 1.1 字句解析
- インタプリタでは、ソースコードをトークン列に変換し、その後抽象構文木に変換する。
- 最初の変換は、字句解析器(トークナイザー、スキャナー)によって実施され、次の変換は構文解析器によって実施される。

### 1.2 トークンを定義する
- トークンを以下のように定義する。
```go
package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string
}

const (
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"

    // 識別子
    IDENT       = "IDENT"   // 変数名
    
    // リテラル
    INT         = "INT"     // 整数

    // 演算子
    ASSIGN      = "="
    PLUS        = "+"

    COMMA       = ","
    SEMICOLON   = ";"

    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"

    // キーワード
    FUNCTION    = "FUNCTION"
    LET         = "LET"
)
```

### 1.3 字句解析器
- テストを定義する。
```go
package lexer

import (
    "testing"
    "monkey/token"
)

func TestNextToken(t *testing.T){
    input := "=+(){},;"

    tests := []struct{
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS,   "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA,  ","},
        {token.SEMICOLON, ";"},
        {token.EOF,    ""}
    }

    l := New(input)

    for i, tt := range tests{
        tok := l.NextToken()

        if tok.Type != tt.expectedType{
            t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral{
            t.Fatal("test[%s] - literal wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }
    }
}
```

### 1.4 トークン集合の拡充と字句解析器の拡張
### 1.5 REPLの始まり

## 第2章 構文解析
## 第3章 評価
## 第4章 インタプリタの拡張
## 付録