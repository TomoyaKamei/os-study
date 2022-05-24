# Writing an Interpreter in Go



## 第1章 字句解析


### 1.1 字句解析
- インタプリタでは、ソースコードをトークン列に変換し、その後抽象構文木に変換する。
- 最初の変換は、字句解析器(トークナイザー、スキャナー)によって実施され、次の変換は構文解析器によって実施される。

### 1.2 トークンを定義する
- トークンを以下のように定義する。
```go
package token

var keywords = map[string]TokenType{
    "fn": Function,
    "let": LET,
}

// 複数文字列の中にキーワードが存在するかどうかを判定する。
func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}

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
    MINUS       = "-"
    BANG        = "!"
    ASTERISK    = "*"
    SLASH       = "/"
    LT          = "<"
    GT          = ">"

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

```go
package lexer

import "monkey/token"

type Lexer struct{
    input           string  
    position        int     // 入力における現在の位置
    readPosition    int     // これから読み込む位置
    ch              byte    // 現在検査中の文字
}

func New(input string) *Lexer{
    l *= &Lexer{input: input}
    l.readChar()
    return l
}

func (l *lexer) readChar(){
    // 読み取る位置がinputよりも多い場合は、現在検査中の文字を0とする。
    if l.readPosition >= len(l.input){
        l.ch = 0
    // 
    }else{
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token{
    var tok token.Token

    l.skipWhiteSpace()

    switch l.ch{
        case '=':
            tok = newToken(token.ASSIGN, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default: 
            if isLetter(l.ch){
                tok.Type = token.LookupIdent(tok.Literal)
                tok.Literal = l.readIdetifier()
                return tok
            } else if isDigit(l.ch){
                tok.Type = token.INT
                tok.LIteral = l.readNumber()
                return tok
            } else {
                tok = newToken(token.ILLEGAL, l.ch)
            }
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string{
    position := l.position
    for isLetter(l.ch){
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string{
    position := l.position
    for isDigit(l.ch){
        l.readChar()
    }
    return l.input(position: l.position)
}

func isLetter(ch byte) bool{
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && <= 'Z' || ch == '_' || ch == '!' || ch == '?'
}

func (l *Lexer) skipWhiteSpace(){
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'{
        l.readChar()
    }
}

func isDigit(ch byte) bool{
    return '0' <= ch && ch <= '9'
}
```

### 1.4 トークン集合の拡充と字句解析器の拡張
### 1.5 REPLの始まり
    - REPLとは、Read・Eval・Print・Loopの略である。
    - これは、コンソールやインタラクティブモードと呼ばれる事もある。
    

## 第2章 構文解析


### 2.1 構文解析器
    - パーサーとは、入力された文字列に対して字句解析器により記号化した入力列から、構文木や抽象構文木を作成する事で、構造化された表現を与える役割を担う。
    - また、抽象構文木等のデータ構造を作成する過程で入力の解析を行う。

### 2.2 パーサージェネレーター
    - パーサジェネレーターは、yacccやbison、ANTLRと呼ばれるツール群である。
    - パーサジェネレーターは、文脈自由文法(context-free grammer; CFG)を入力として用いて、パーサーを作成する。
    - パーサジェネレーターを用いて出力されたパーサーは、入力に対して構文木を出力させる事が出来る。
    - 本書では、パーサージェネレータを使用しない。理由としてはパーサジェネレータはよりよい解法ではあるが、パーサを手動で作成する事によりどのようにパーサが動作しているかがはじめてブラックボックスではなくなるからである。

### 2.3 Monkey言語のための構文解析器を書く
    - Monkey言語のために作成する構文解析器は、"トップダウン演算子優先順位"構文解析器で、別名Pratt構文解析器とも呼ばれている。

### 2.4 構文解析器の第一歩
    - 式(Expression)は値を生成する。
        - e.g. fn(a, b){ return a + b; };, 10, 15
    - 文(Statement)は値を生成しない。
        - e.g. let x = 5;
    - Let文は、let <識別子> = <式>を満たす。
    - 構文解析器の動きおおよそ以下のように実行される。
        1. 構文解析器はEOF_TOKENが出力されるまでループする。
        2. トークンが出力される度、Statement(文)となるキーワードかどうか判別する。
            e.g. TOKENがLetの場合、LetStatementの記述に従っているか以降のTOKENを検査する。
        3. Statement(文)が確定した場合、ルートノードに格納される。

### 2.5 return文の構文解析
    - Return文は、return <式>を満たす。
    
### 2.6 式の構文解析
    - 一般的に式の構文解析は、文の構文解析よりも困難である。
    - 例として、()を用いた構文評価の優先順位である。
        - e.g. - 5 - 10, 5 * (add(10, 20) + 323)など
    - この複雑な構文解析を実施するために使用するのが、"トップダウン演算子優先順位構文解析"/"Pratt構文解析"と呼ばれるものである。
        - Pratt構文解析器の用語
            - 前置演算子
                - オペランドの前に置かれる演算子を指す。
                - e.g. --5
            - 後置演算子
                - オペランドの後に置かれる演算子を指す。
                - e.g. foobar++
            - 中間演算子
                - オペランドの中間に置かれる演算子を指す。
                - e.g. 5 * 8
    - 式文は、一つの式からなる文の事を指す。
        - e.g. x + 10;
    
    

### 2.7 Pratt構文解析の仕組み
### 2.8 構文解析器の拡張
### 2.9 読み込み - 構文解析 - 表示 - 繰り返し

## 第3章 評価
## 第4章 インタプリタの拡張
## 付録