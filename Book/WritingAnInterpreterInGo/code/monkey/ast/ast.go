package ast

import (
	"bytes"
	"monkey/token"
)

// Node ASTの要素を指す。
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements{
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement let <Identifier> = <Expression>を表し、Node・Statement(文)インターフェースを充足する。
type LetStatement struct {
	Token token.Token // token.LETトークン
	Name  *Identifier // 識別子
	Value Expression  // 式
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string{
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil{
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier 識別子を表し、Node・Expression(式)インターフェースを充足する。
type Identifier struct {
	Token token.Token // token.IDENTトークン
	Value string      // 識別子名
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

// ReturnStatement return <expression>を表し、Node・Statement(文)インターフェースを充足する。
type ReturnStatement struct {
	Token       token.Token // token.RETURNトークン
	ReturnValue Expression  // 戻り値となる式を保持するためのReturnValueフィールドがある。
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string{
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil{
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement 式文を表し、Node・Statement(文)インターフェースを充足する。
type ExpressionStatement struct{
	Token  token.Token		//式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil{
		return es.Expression.String()
	}
	return ""
}