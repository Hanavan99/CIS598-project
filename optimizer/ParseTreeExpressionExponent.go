package optimizer

import "math"

type ParseTreeExpressionExponent struct {
	base     ParseTreeExpression
	exponent ParseTreeExpression
}

func (ptee ParseTreeExpressionExponent) Children() []ParseTree {
	return []ParseTree{ptee.base, ptee.exponent}
}

func (ptee ParseTreeExpressionExponent) Evaluate(env Environment) float64 {
	return math.Pow(ptee.base.Evaluate(env), ptee.exponent.Evaluate(env))
}

func (ptee ParseTreeExpressionExponent) HasKnownValue(env Environment) bool {
	return ptee.base.HasKnownValue(env) && ptee.exponent.HasKnownValue(env)
}

func (ptee ParseTreeExpressionExponent) ToString() string {
	return ptee.base.ToString() + " ^ " + ptee.exponent.ToString()
}
