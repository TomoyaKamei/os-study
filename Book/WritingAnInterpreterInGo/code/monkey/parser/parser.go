package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser 構文解析器
type Parser struct{
	l *lexer.Lexer			// 字句解析器
	curToken token.Token	// 現在のトークン
	peekToken token.Token	// 次のトークン
}

// New 字句解析器から構文解析器の初期化を実施する関数
func New(l *lexer.Lexer) *Parser{
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	
	return p
}

// nextToken curTokenとpeekTokenを次に進めるメソッドである。
func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram プログラムをパースし、ルートノードを作成する関数
func (p *Parser) ParseProgram() *ast.Program{
	return nil
}