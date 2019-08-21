// 解释器模式（Interpreter Pattern）提供了评估语言的语法或表达式的方式，它属于行为型模式。
// 这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。

package interpreter

import "strings"

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression{
	return &TerminalExpression{data:data}
}

func (t *TerminalExpression) Interpret(context string) bool{
	if strings.Contains(context, t.data){
		return true
	} else {
		return false
	}
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression{
	return &OrExpression{
		expr1:expr1,
		expr2:expr2,
	}
}

func (a *OrExpression)Interpret(context string) bool{
	return a.expr1.Interpret(context) || a.expr2.Interpret(context)
}


type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) *AndExpression{
	return &AndExpression{
		expr1:expr1,
		expr2:expr2,
	}
}

func (a *AndExpression)Interpret(context string) bool{
	return a.expr1.Interpret(context) && a.expr2.Interpret(context)
}

type InterpreterPattern struct {
}

func (i *InterpreterPattern) GetMaleExpression() Expression{
	robert := NewTerminalExpression("Robert")
	john := NewTerminalExpression("John")
	return NewOrExpression(robert, john)
}

func (i *InterpreterPattern) GetMarriedWomanExpression() Expression{
	julie := NewTerminalExpression("Julie")
	married := NewTerminalExpression("Married")
	return NewAndExpression(julie, married)
}