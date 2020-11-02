package optimizer

type ParseTreeExpressionVariable struct {
	name string
}

func (ptev ParseTreeExpressionVariable) Children() []ParseTree {
	return nil
}

func (ptev ParseTreeExpressionVariable) Evaluate(env Environment) float64 {
	return env.Get(ptev.name).(float64)
}

func (ptev ParseTreeExpressionVariable) HasKnownValue(env Environment) bool {
	return env.Exists(ptev.name)
}

func (ptev ParseTreeExpressionVariable) ToString() string {
	return ptev.name
}
