package optimizer

import "fmt"

type ParseTreeExpressionSubtract struct {
	left  ParseTreeExpression
	right ParseTreeExpression
}

func (ptes ParseTreeExpressionSubtract) Children() []ParseTree {
	return []ParseTree{ptes.left, ptes.right}
}

func (ptes ParseTreeExpressionSubtract) Evaluate(env Environment) float64 {
	return ptes.left.Evaluate(env) - ptes.right.Evaluate(env)
}

func (ptes ParseTreeExpressionSubtract) HasKnownValue(env Environment) bool {
	return ptes.left.HasKnownValue(env) && ptes.right.HasKnownValue(env)
}

func (ptes ParseTreeExpressionSubtract) ToString() string {
	return fmt.Sprintf("(%s - %s)", ptes.left.ToString(), ptes.right.ToString())
}
