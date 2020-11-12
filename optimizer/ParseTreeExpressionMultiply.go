package optimizer

import "fmt"

type ParseTreeExpressionMultiply struct {
	left  ParseTreeExpression
	right ParseTreeExpression
}

func (ptem ParseTreeExpressionMultiply) Children() []ParseTree {
	return []ParseTree{ptem.left, ptem.right}
}

func (ptem ParseTreeExpressionMultiply) Evaluate(env Environment) float64 {
	return ptem.left.Evaluate(env) * ptem.right.Evaluate(env)
}

func (ptem ParseTreeExpressionMultiply) HasKnownValue(env Environment) bool {
	return ptem.left.HasKnownValue(env) && ptem.right.HasKnownValue(env)
}

func (ptem ParseTreeExpressionMultiply) ToString() string {
	return fmt.Sprintf("(%s * %s)", ptem.left.ToString(), ptem.right.ToString())
}
