package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  // これから読み込む位置
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// 読み取る位置がinputよりも多い場合は、現在検査中の文字を0とする。
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Lexerを更新する。
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// 次のトークンを出力する。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 改行や空白などを飛ばす関数
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// ==である場合
			ch := l.ch		// 現在検査中の文字
			l.readChar()	// 検査する文字を次の文字へ
			literal := string(ch) + string(l.ch) 
			tok = token.Token{Type: token.EQ, Literal: literal}	// EQ("==")としてトークンを出力
		} else {
			// =である場合
			tok = newToken(token.ASSIGN, l.ch)	// ASSIGN("=")としてトークンを出力
		}
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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			// !=である場合
			ch := l.ch		// 現在検査中の文字
			l.readChar()	// 検査する文字を次の文字へ
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}	// NOT_EQ("!=")としてトークンを出力
		} else {
			// !である場合
			tok = newToken(token.BANG, l.ch)	// BANG("!")としてトークンを出力
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 文字列である場合
			tok.Literal = l.readIdentifier()	// 文字列を抜き出す
			tok.Type = token.LookupIdent(tok.Literal)	// キーワードのタイプを出力
			return tok
		} else if isDigit(l.ch) {
			// 数字である場合
			tok.Type = token.INT
			tok.Literal = l.readNumber()	// 数字を抜き出す
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// トークンを出力する関数
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Lexerから文字列を抜き出し、Lexerを更新する関数
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 文字列かどうか判別する関数
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 空白等を飛ばす関数
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Lexerから数字を抜き出し、Lexerを更新する関数
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 数字かどうか判別する関数
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 次に読み込まれる文字を覗き見る関数
func (l *Lexer) peekChar() byte {
	// 読み取る位置がinputよりも多い場合は、現在検査中の文字を0とする。
	if l.readPosition >= len(l.input) { 
		return 0
	} 
	
	// 次に出力される文字を返す。
	return l.input[l.readPosition]
}
