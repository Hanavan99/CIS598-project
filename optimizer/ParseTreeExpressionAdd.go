package optimizer

import "fmt"

type ParseTreeExpressionAdd struct {
	left  ParseTreeExpression
	right ParseTreeExpression
}

func (ptea ParseTreeExpressionAdd) Children() []ParseTree {
	return []ParseTree{ptea.left, ptea.right}
}

func (ptea ParseTreeExpressionAdd) Evaluate(env Environment) float64 {
	return ptea.left.Evaluate(env) + ptea.right.Evaluate(env)
}

func (ptea ParseTreeExpressionAdd) HasKnownValue(env Environment) bool {
	return ptea.left.HasKnownValue(env) && ptea.right.HasKnownValue(env)
}

func (ptea ParseTreeExpressionAdd) ToString() string {
	return fmt.Sprintf("(%s + %s)", ptea.left.ToString(), ptea.right.ToString())
}
