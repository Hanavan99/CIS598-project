package optimizer

type ParseTreeExpressionDivide struct {
	left  ParseTreeExpression
	right ParseTreeExpression
}

func (pted ParseTreeExpressionDivide) Children() []ParseTree {
	return []ParseTree{pted.left, pted.right}
}

func (pted ParseTreeExpressionDivide) Evaluate(env Environment) float64 {
	return pted.left.Evaluate(env) / pted.right.Evaluate(env)
}

func (pted ParseTreeExpressionDivide) HasKnownValue(env Environment) bool {
	return pted.left.HasKnownValue(env) && pted.right.HasKnownValue(env)
}

func (pted ParseTreeExpressionDivide) ToString() string {
	return pted.left.ToString() + " / " + pted.right.ToString()
}
