package optimizer

import "fmt"

type ParseTreeExpressionConstant struct {
	value float64
}

func (ptec ParseTreeExpressionConstant) Children() []ParseTree {
	return nil
}

func (ptec ParseTreeExpressionConstant) Evaluate(env Environment) float64 {
	return ptec.value
}

func (ptec ParseTreeExpressionConstant) HasKnownValue(env Environment) bool {
	return true
}

func (ptec ParseTreeExpressionConstant) ToString() string {
	return fmt.Sprintf("%f", ptec.value)
}
