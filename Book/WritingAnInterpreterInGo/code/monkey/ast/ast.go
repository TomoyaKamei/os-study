package ast

import "monkey/token"

// Node ASTの要素を指す。
type Node interface {
	TokenLiteral() string
}

// Statement Nodeの一種で、文を表す。
type Statement interface {
	Node
	statementNode()
}

// Expression Nodeの一種で、式を表す。
type Expression interface {
	Node
	expressionNode()
}

// Program ASTのルートノードを表す。
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		//Statementの要素がProgramに複数存在する場合
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement let <Identifier> = <Expression>を表し、Node・Statement(文)インターフェースを充足する。
type LetStatement struct {
	Token token.Token // token.LETトークン
	Name  *Identifier // 識別子
	Value Expression  // 式
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier 識別子を表し、Node・Expression(式)インターフェースを充足する。
type Identifier struct {
	Token token.Token // token.IDENTトークン
	Value string      // 識別子名
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
